package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
