package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

// These tests cover the update path for the resources whose URL id is optional
// and computed. When the config left that id unset, the plugin framework marked
// it unknown in the plan, the client turned the unknown into an empty string,
// and every update and delete went to the collection URL instead of the
// resource. The API answered 404, the provider ignored the status, and the
// zero-value response nulled every attribute in state.
//
// They run against the fake Console API, so they need no Statsig credentials.

func gateConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_gate" "regression" {
  name        = "regression_gate"
  description = %q
  is_enabled  = true
  id_type     = "userID"
}
`, description)
}

func gateWithExplicitIdConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_gate" "regression" {
  id          = "regression_gate_explicit"
  name        = "regression_gate_explicit"
  description = %q
  is_enabled  = true
  id_type     = "userID"
}
`, description)
}

// segmentConfig varies only the rule name because rules are the only thing a
// segment update can change. The Console API writes them through
// segments/<id>/conditional and the provider sends nothing else with them, so
// editing any other attribute fails the apply. This config proves the rules
// request reaches the resource URL, not that segment updates work in general.
func segmentConfig(ruleName string) string {
	return fmt.Sprintf(`
resource "statsig_segment" "regression" {
  name    = "regression_segment"
  type    = "rule_based"
  id_type = "userID"
  rules = [{
    name            = %q
    pass_percentage = 100
    conditions      = [{ type = "email", operator = "any", target_value = ["a@b.com"] }]
  }]
}
`, ruleName)
}

func keysConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_keys" "regression" {
  description  = %q
  type         = "SERVER"
  environments = ["production"]
  scopes       = []
}
`, description)
}

func metricConfig(id string, description string) string {
	return fmt.Sprintf(`
resource "statsig_metric" "regression" {
  id                   = %q
  name                 = "Regression Metric"
  description          = %q
  type                 = "event_count_custom"
  custom_roll_up_end   = 14
  custom_roll_up_start = 0
  rollup_time_window   = "custom"
  unit_types           = ["userID"]
  metric_events = [
    {
      criteria = []
      name     = "regression_event"
    }
  ]
}
`, id, description)
}

func tagConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_tag" "regression" {
  name        = "regression_tag"
  description = %q
  is_core     = false
}
`, description)
}

func TestAccUpdatePathAddressesTheResource(t *testing.T) {
	cases := []struct {
		name   string
		create string
		update string
		wants  []string
	}{
		{
			name:   "gate",
			create: gateConfig("first"),
			update: gateConfig("second"),
			wants: []string{
				"PATCH /console/v1/gates/regression_gate",
				"DELETE /console/v1/gates/regression_gate",
			},
		},
		{
			// Only the rules request is covered here. See segmentConfig.
			name:   "segment rules",
			create: segmentConfig("rule one"),
			update: segmentConfig("rule one renamed"),
			wants: []string{
				"POST /console/v1/segments/regression_segment/conditional",
				"DELETE /console/v1/segments/regression_segment",
			},
		},
		{
			// The key is server-generated, so writing it into the config is not
			// an option here the way an explicit gate id is.
			name:   "keys",
			create: keysConfig("first"),
			update: keysConfig("second"),
			wants: []string{
				"PATCH /console/v1/keys/" + fakeGeneratedKey,
				"DELETE /console/v1/keys/" + fakeGeneratedKey,
			},
		},
		{
			// Control: setting id in the config kept the plan value known, so
			// this path worked before the fix and must keep working.
			name:   "gate with explicit id",
			create: gateWithExplicitIdConfig("first"),
			update: gateWithExplicitIdConfig("second"),
			wants: []string{
				"PATCH /console/v1/gates/regression_gate_explicit",
				"DELETE /console/v1/gates/regression_gate_explicit",
			},
		},
		{
			// Control: statsig_tag puts a required name in the URL, so the
			// framework never marked it unknown and it was never affected.
			name:   "tag",
			create: tagConfig("first"),
			update: tagConfig("second"),
			wants: []string{
				"PATCH /console/v1/tags/regression_tag",
				"DELETE /console/v1/tags/regression_tag",
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			api := startFakeConsoleAPI(t)

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccLocalProviders(),
				Steps: []resource.TestStep{
					{Config: testCase.create},
					{Config: testCase.update},
				},
			})

			log := api.requestLog()
			for _, want := range testCase.wants {
				assert.Contains(t, log, want)
			}
			assertAddressedResources(t, log)
		})
	}
}

// statsig_metric.id carries no length validator, so a config can set it to the
// empty string and the plan value stays known. The update then composes
// "metrics/" and posts the whole metric to the collection, which the Console API
// reads as a create.
func TestAccMetricUpdateRefusesAnEmptyId(t *testing.T) {
	api := startFakeConsoleAPI(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccLocalProviders(),
		Steps: []resource.TestStep{
			{Config: metricConfig("regression_metric", "first")},
			{
				Config:      metricConfig("", "second"),
				ExpectError: regexp.MustCompile("resource id is empty"),
			},
		},
	})

	log := api.requestLog()
	assert.Contains(t, log, "POST /console/v1/metrics", "the metric was never created")
	assertAddressedResources(t, log)
}

// assertAddressedResources fails on any request aimed at a collection where a
// single resource was meant. An empty id left either a trailing slash or an
// empty path segment before a sub-resource.
func assertAddressedResources(t *testing.T, log []string) {
	t.Helper()

	for _, request := range log {
		assert.NotContains(t, request, "//", "%q addressed the collection, not the resource", request)
		assert.False(t, strings.HasSuffix(request, "/"), "%q addressed the collection, not the resource", request)
	}
}
