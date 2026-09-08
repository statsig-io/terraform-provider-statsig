package tests

import (
	"fmt"
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
			name:   "segment",
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
