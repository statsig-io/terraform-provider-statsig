package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccServerKey(t *testing.T) {
	serverKey, _ := os.ReadFile("test_resources/key_server.tf")
	serverKeyPatch, _ := os.ReadFile("test_resources/key_server_patch.tf")

	name := "statsig_keys.server_key"
	var key string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				Config: string(serverKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^secret-.*")),
					resource.TestCheckResourceAttr(name, "type", "SERVER"),
					resource.TestCheckResourceAttrSet(name, "description"),
					resource.TestCheckResourceAttr(name, "environments.#", "1"),
					resource.TestCheckResourceAttr(name, "environments.0", "production"),
					resource.TestCheckResourceAttr(name, "scopes.#", "0"),
					resource.TestCheckNoResourceAttr(name, "target_app_id"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
					testAccExtractResourceAttr(name, "key", &key),
				),
			},
			RefreshNoopPlanCheck(),
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_key", key)
				},
				Config: string(serverKeyPatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^secret-.*")),
					resource.TestCheckResourceAttr(name, "type", "SERVER"),
					resource.TestCheckResourceAttrSet(name, "description"),
					resource.TestCheckResourceAttr(name, "environments.#", "2"),
					resource.TestCheckResourceAttr(name, "environments.0", "production"),
					resource.TestCheckResourceAttr(name, "environments.1", "staging"),
					resource.TestCheckResourceAttr(name, "scopes.#", "0"),
					resource.TestCheckNoResourceAttr(name, "target_app_id"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccClientKey(t *testing.T) {
	clientKey, _ := os.ReadFile("test_resources/key_client.tf")
	clientKeyPatch, _ := os.ReadFile("test_resources/key_client_patch.tf")

	name := "statsig_keys.client_key"
	var key string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				Config: string(clientKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^client-.*")),
					resource.TestCheckResourceAttr(name, "type", "CLIENT"),
					resource.TestCheckResourceAttrSet(name, "description"),
					resource.TestCheckResourceAttr(name, "environments.#", "1"),
					resource.TestCheckResourceAttr(name, "environments.0", "production"),
					resource.TestCheckResourceAttr(name, "scopes.#", "1"),
					resource.TestCheckResourceAttr(name, "scopes.0", "client_download_config_specs"),
					resource.TestCheckNoResourceAttr(name, "target_app_id"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
					testAccExtractResourceAttr(name, "key", &key),
				),
			},
			RefreshNoopPlanCheck(),
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_key", key)
				},
				Config: string(clientKeyPatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^client-.*")),
					resource.TestCheckResourceAttr(name, "type", "CLIENT"),
					resource.TestCheckResourceAttrSet(name, "description"),
					resource.TestCheckResourceAttr(name, "environments.#", "2"),
					resource.TestCheckResourceAttr(name, "environments.0", "production"),
					resource.TestCheckResourceAttr(name, "environments.1", "staging"),
					resource.TestCheckResourceAttr(name, "scopes.#", "1"),
					resource.TestCheckResourceAttr(name, "scopes.0", "client_download_config_specs"),
					resource.TestCheckNoResourceAttr(name, "target_app_id"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccConsoleKey(t *testing.T) {
	consoleKey, _ := os.ReadFile("test_resources/key_console.tf")
	consoleKeyPatch, _ := os.ReadFile("test_resources/key_console_patch.tf")

	name := "statsig_keys.console_key"
	var key string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				Config: string(consoleKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^console-.*")),
					resource.TestCheckResourceAttr(name, "type", "CONSOLE"),
					resource.TestCheckResourceAttrSet(name, "description"),
					resource.TestCheckResourceAttr(name, "scopes.0", "omni_read_only"),
					resource.TestCheckNoResourceAttr(name, "target_app_id"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
					testAccExtractResourceAttr(name, "key", &key),
				),
			},
			RefreshNoopPlanCheck(),
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_key", key)
				},
				Config: string(consoleKeyPatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^console-.*")),
					resource.TestCheckResourceAttr(name, "type", "CONSOLE"),
					resource.TestCheckResourceAttrSet(name, "description"),
					resource.TestCheckResourceAttr(name, "scopes.0", "omni_read_write"),
					resource.TestCheckNoResourceAttr(name, "target_app_id"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
				),
			},
		},
	})
}

// POST /console/v1/keys takes targetAppID and answers with primaryTargetApp, a
// display name, so no response can supply the value target_app_id holds.
// Overwriting the attribute from the response made the applied state differ
// from the plan, which aborts the apply. This case runs against the fake
// Console API and needs no Statsig credentials.
func TestAccKeysKeepsConfiguredTargetApps(t *testing.T) {
	const targetAppId = "4SRgGcr8uWNVW3c2OGWFZC"
	const secondaryTargetAppId = "2Kd9hLpQzXcVbNmR4TsYuI"

	startFakeConsoleAPI(t)

	config := fmt.Sprintf(`
resource "statsig_keys" "regression" {
  description              = "edge server key"
  type                     = "SERVER"
  target_app_id            = %q
  secondary_target_app_ids = [%q]
  environments             = ["production"]
  scopes                   = []
}
`, targetAppId, secondaryTargetAppId)

	name := "statsig_keys.regression"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccLocalProviders(),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "target_app_id", targetAppId),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "1"),
					resource.TestCheckResourceAttr(name, "secondary_target_app_ids.0", secondaryTargetAppId),
				),
			},
		},
	})
}

const keysTargetAppId = "4SRgGcr8uWNVW3c2OGWFZC"
const keysSecondaryTargetAppId = "2Kd9hLpQzXcVbNmR4TsYuI"

func keysWithTargetAppsConfig(description string) string {
	return fmt.Sprintf(`
resource "statsig_keys" "regression" {
  description              = %q
  type                     = "SERVER"
  target_app_id            = %q
  secondary_target_app_ids = [%q]
  environments             = ["production"]
  scopes                   = []
}
`, description, keysTargetAppId, keysSecondaryTargetAppId)
}

// patchedKeyFields decodes the body of the one PATCH the fake Console API saw,
// so a test can tell an absent field from a null one.
func patchedKeyFields(t *testing.T, api *fakeConsoleAPI) map[string]json.RawMessage {
	t.Helper()

	bodies := api.requestBodiesFor("PATCH /console/v1/keys/" + fakeGeneratedKey)
	require.Len(t, bodies, 1, "the key must have been updated exactly once")

	var body map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(bodies[0], &body))
	return body
}

// The Console API leaves a target app assigned when the field is absent from the
// request, so dropping the attributes from a config must send a PATCH that omits
// them. Without this an unrelated description edit unassigns a live target app.
func TestAccKeysOmitsTargetAppsWhenConfigDropsThem(t *testing.T) {
	api := startFakeConsoleAPI(t)

	withoutTargetApps := `
resource "statsig_keys" "regression" {
  description  = "edge server key, renamed"
  type         = "SERVER"
  environments = ["production"]
  scopes       = []
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccLocalProviders(),
		Steps: []resource.TestStep{
			{Config: keysWithTargetAppsConfig("edge server key")},
			{Config: withoutTargetApps},
		},
	})

	body := patchedKeyFields(t, api)
	assert.NotContains(t, body, "targetAppID", "an absent field is what leaves the assignment alone")
	assert.NotContains(t, body, "secondaryTargetAppIDs", "an absent field is what leaves the assignment alone")
}

// An explicitly empty value is the one way to unassign a target app from
// Terraform. The API reads null and an empty array as a clear, and never reads
// an empty string that way.
func TestAccKeysClearsTargetAppsWhenConfigIsExplicitlyEmpty(t *testing.T) {
	api := startFakeConsoleAPI(t)

	emptyTargetApps := `
resource "statsig_keys" "regression" {
  description              = "edge server key"
  type                     = "SERVER"
  target_app_id            = ""
  secondary_target_app_ids = []
  environments             = ["production"]
  scopes                   = []
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccLocalProviders(),
		Steps: []resource.TestStep{
			{Config: keysWithTargetAppsConfig("edge server key")},
			{Config: emptyTargetApps},
		},
	})

	body := patchedKeyFields(t, api)
	assert.JSONEq(t, "null", string(body["targetAppID"]))
	assert.JSONEq(t, "[]", string(body["secondaryTargetAppIDs"]))
	assert.NotEqual(t, `""`, string(body["targetAppID"]),
		"the API stores an empty string as an identifier instead of clearing the assignment")
}

// A create has no prior state for the plan to reuse, so an omitted
// secondary_target_app_ids still resolves to an empty list. A null list there
// breaks length() and for_each in a working configuration.
func TestAccKeysCreatesEmptySecondaryTargetAppIds(t *testing.T) {
	startFakeConsoleAPI(t)

	config := `
resource "statsig_keys" "regression" {
  description  = "server key with no target app"
  type         = "SERVER"
  environments = ["production"]
  scopes       = []
}
`

	name := "statsig_keys.regression"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccLocalProviders(),
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(name, tfjsonpath.New("secondary_target_app_ids"),
						knownvalue.ListSizeExact(0)),
					statecheck.ExpectKnownValue(name, tfjsonpath.New("target_app_id"),
						knownvalue.Null()),
				},
			},
		},
	})
}
