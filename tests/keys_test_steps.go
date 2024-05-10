package tests

import (
	"errors"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func getImportStateIDFunc(name string, key *string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		resource, ok := s.RootModule().Resources[name]
		if !ok || resource == nil {
			return "", errors.New("Resource not found")
		}
		*key = resource.Primary.Attributes["key"]
		return *key, nil
	}
}

func TestAccServerKey(t *testing.T) []resource.TestStep {
	serverKey, _ := os.ReadFile("test_resources/key_server.tf")
	serverKeyPatch, _ := os.ReadFile("test_resources/key_server_patch.tf")

	name := "statsig_keys.server_key"
	var key string

	return []resource.TestStep{
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
			),
		},
		RefreshNoopPlanCheck(),
		{
			ImportStateIdFunc: getImportStateIDFunc(name, &key),
			ImportState:       true,
			ResourceName:      name,
		},
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
	}
}

func TestAccClientKey(t *testing.T) []resource.TestStep {
	clientKey, _ := os.ReadFile("test_resources/key_client.tf")
	clientKeyPatch, _ := os.ReadFile("test_resources/key_client_patch.tf")

	name := "statsig_keys.client_key"
	var key string

	return []resource.TestStep{
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
			),
		},
		RefreshNoopPlanCheck(),
		{
			ImportStateIdFunc: getImportStateIDFunc(name, &key),
			ImportState:       true,
			ResourceName:      name,
		},
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
	}
}

func TestAccConsoleKey(t *testing.T) []resource.TestStep {
	consoleKey, _ := os.ReadFile("test_resources/key_console.tf")
	consoleKeyPatch, _ := os.ReadFile("test_resources/key_console_patch.tf")

	name := "statsig_keys.console_key"
	var key string

	return []resource.TestStep{
		{
			Config: string(consoleKey),
			Check: resource.ComposeTestCheckFunc(
				resource.TestMatchResourceAttr(name, "key", regexp.MustCompile("^console-.*")),
				resource.TestCheckResourceAttr(name, "type", "CONSOLE"),
				resource.TestCheckResourceAttrSet(name, "description"),
				resource.TestCheckResourceAttr(name, "scopes.0", "omni_read_only"),
				resource.TestCheckNoResourceAttr(name, "target_app_id"),
				resource.TestCheckResourceAttr(name, "secondary_target_app_ids.#", "0"),
			),
		},
		RefreshNoopPlanCheck(),
		{
			ImportStateIdFunc: getImportStateIDFunc(name, &key),
			ImportState:       true,
			ResourceName:      name,
		},
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
	}
}
