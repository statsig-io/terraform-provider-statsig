package tests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// An Optional+Computed attribute the configuration never mentions is unknown in
// the plan, and the request used to carry its Go zero value: "", false, []. The
// Console API reads a present field as an instruction to set it, so editing one
// attribute could clear another that was set in the Statsig Console. Repairing
// the update URL is what made those requests land, so these run against the fake
// Console API, which keeps a field it is not sent.
//
// They need no Statsig credentials.

func gateWithoutIsEnabledConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_gate" "regression" {
  name        = "regression_gate"
  description = %q
  id_type     = "userID"
}
`, description)
}

func keysWithoutScopesConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_keys" "regression" {
  description  = %q
  type         = "SERVER"
  environments = ["production"]
}
`, description)
}

// metricWarehouseNativeConfig omits warehouse_native.metric_dimension_columns
// and warehouse_native.denominator_criteria, the same as the example config in
// test_resources/metric_warehouse_native.tf.
func metricWarehouseNativeConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_metric" "regression" {
  id          = "regression_wn_metric"
  name        = "Regression Warehouse Native Metric"
  description = %q
  type        = "user_warehouse"
  unit_types  = ["userID"]
  warehouse_native = {
    metric_source_name = "shoppy_events"
    aggregation        = "count"
  }
}
`, description)
}

func gateWithRuleConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_gate" "regression" {
  name        = "regression_gate"
  description = %q
  id_type     = "userID"
  rules = [{
    name            = "everyone"
    pass_percentage = 100
    conditions      = [{ type = "public" }]
  }]
}
`, description)
}

// gateWithEmptyRuleListsConfig states the empty lists that gateWithRuleConfig
// omits, which is the difference between "leave this alone" and "clear this".
func gateWithEmptyRuleListsConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_gate" "regression" {
  name        = "regression_gate"
  description = %q
  id_type     = "userID"
  rules = [{
    name            = "everyone"
    pass_percentage = 100
    environments    = []
    conditions      = [{ type = "public", target_value = [] }]
  }]
}
`, description)
}

func TestAccUnrelatedEditKeepsConsoleSetValues(t *testing.T) {
	t.Run("an enabled gate stays enabled", func(t *testing.T) {
		api := startFakeConsoleAPI(t)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccLocalProviders(),
			Steps: []resource.TestStep{
				{Config: gateWithoutIsEnabledConfig("first")},
				{
					PreConfig: func() {
						api.setRecordField("gates", "regression_gate", "isEnabled", true)
					},
					Config: gateWithoutIsEnabledConfig("second"),
					Check: checkRecordField(api, "gates", "regression_gate", "isEnabled", true,
						"editing the description disabled a gate the config never mentions"),
				},
			},
		})
	})

	t.Run("a key keeps the scopes it was granted", func(t *testing.T) {
		api := startFakeConsoleAPI(t)
		granted := []interface{}{"gates:read"}

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccLocalProviders(),
			Steps: []resource.TestStep{
				{Config: keysWithoutScopesConfig("first")},
				{
					PreConfig: func() {
						api.setRecordField("keys", fakeGeneratedKey, "scopes", granted)
					},
					Config: keysWithoutScopesConfig("second"),
					Check: checkRecordField(api, "keys", fakeGeneratedKey, "scopes", granted,
						"editing the description revoked scopes the config never mentions"),
				},
			},
		})
	})
}

// The rule reaches nested attributes too, so the update request has to leave
// out a nested attribute the configuration never mentions. These read the
// request the provider built rather than the fake's records, because a nested
// field is only visible in the body: the fake replaces a whole nested object
// the way the Console API does.
func TestAccUnrelatedEditOmitsUnspecifiedNestedAttributes(t *testing.T) {
	t.Run("a warehouse native metric leaves out the columns the config omits", func(t *testing.T) {
		api := startFakeConsoleAPI(t)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccLocalProviders(),
			Steps: []resource.TestStep{
				{Config: metricWarehouseNativeConfig("first")},
				{Config: metricWarehouseNativeConfig("second")},
			},
		})

		body := lastRequestBody(t, api, "POST /console/v1/metrics/regression_wn_metric")

		var warehouseNative map[string]json.RawMessage
		require.NoError(t, json.Unmarshal(body["warehouseNative"], &warehouseNative))

		assert.JSONEq(t, `"shoppy_events"`, string(warehouseNative["metricSourceName"]))
		for _, field := range []string{"metricDimensionColumns", "denominatorCriteria"} {
			assert.NotContains(t, warehouseNative, field,
				"editing the description sent %s, which clears it in the Statsig Console", field)
		}
	})

	t.Run("a gate rule leaves out the id and the lists the config omits", func(t *testing.T) {
		api := startFakeConsoleAPI(t)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccLocalProviders(),
			Steps: []resource.TestStep{
				{Config: gateWithRuleConfig("first")},
				{Config: gateWithRuleConfig("second")},
			},
		})

		rule, condition := gateRuleFromLastUpdate(t, api)

		assert.NotContains(t, rule, "id",
			"an empty rule id renames the rule the Console API matches on")
		assert.NotContains(t, rule, "environments",
			"editing the description scoped the rule to no environments")
		assert.NotContains(t, condition, "targetValue",
			"editing the description cleared the condition's targeting values")
	})
}

// The other half of the rule: an empty value the configuration does state is an
// instruction to clear, so it still has to be sent. Without this the fix above
// could be "never send a nested list", which would break a deliberate clear.
func TestAccExplicitlyEmptyNestedAttributesAreStillSent(t *testing.T) {
	api := startFakeConsoleAPI(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccLocalProviders(),
		Steps: []resource.TestStep{
			{Config: gateWithEmptyRuleListsConfig("first")},
			{Config: gateWithEmptyRuleListsConfig("second")},
		},
	})

	rule, condition := gateRuleFromLastUpdate(t, api)

	assert.JSONEq(t, `[]`, string(rule["environments"]),
		"an empty environments list the config states was not sent, so it cannot clear")
	assert.JSONEq(t, `[]`, string(condition["targetValue"]),
		"an empty target_value list the config states was not sent, so it cannot clear")
}

// gateRuleFromLastUpdate returns the single rule and its single condition from
// the last gate update, which is where a nested attribute is visible.
func gateRuleFromLastUpdate(t *testing.T, api *fakeConsoleAPI) (rule, condition map[string]json.RawMessage) {
	t.Helper()

	body := lastRequestBody(t, api, "PATCH /console/v1/gates/regression_gate")

	var rules []map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(body["rules"], &rules))
	require.Len(t, rules, 1)

	var conditions []map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(rules[0]["conditions"], &conditions))
	require.Len(t, conditions, 1)

	return rules[0], conditions[0]
}

// lastRequestBody decodes the body of the last request matching
// "METHOD /path", which in these tests is the update.
func lastRequestBody(t *testing.T, api *fakeConsoleAPI, entry string) map[string]json.RawMessage {
	t.Helper()

	bodies := api.requestBodiesFor(entry)
	require.NotEmpty(t, bodies, "%s was never requested", entry)

	var body map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(bodies[len(bodies)-1], &body))
	return body
}

// checkRecordField reads the fake Console API rather than Terraform state, so
// the assertion is about what the remote still holds. It runs as a step check
// because the harness destroys the resource once the last step finishes.
func checkRecordField(
	api *fakeConsoleAPI,
	collection string,
	id string,
	field string,
	want interface{},
	message string,
) resource.TestCheckFunc {
	return func(*terraform.State) error {
		got := api.recordField(collection, id, field)
		if !reflect.DeepEqual(want, got) {
			return fmt.Errorf("%s: %s is %v, want %v", message, field, got, want)
		}
		return nil
	}
}
