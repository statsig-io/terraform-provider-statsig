package statsig

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExperimentBasic(t *testing.T) {
	basicExperiment, _ := os.ReadFile("basic_experiment.tf")

	key := "statsig_experiment.my_experiment"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: string(basicExperiment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExperimentExists(key),

					// Test Group
					resource.TestCheckResourceAttr(key, "groups.0.name", "Test Group"),
					resource.TestCheckResourceAttr(key, "groups.0.size", "50"),
					resource.TestCheckResourceAttr(key, "groups.0.parameter_values_json", "{\"a_bool\":true,\"a_string\":\"test_string\"}"),

					// Control Group
					resource.TestCheckResourceAttr(key, "groups.1.name", "Control Group"),
					resource.TestCheckResourceAttr(key, "groups.1.size", "50"),
					resource.TestCheckResourceAttr(key, "groups.1.parameter_values_json", "{\"a_bool\":false,\"a_string\":\"control_string\"}"),
				),
			},
		},
	})
}

func testAccCheckExperimentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "statsig_experiment" {
			continue
		}

		experimentID := rs.Primary.ID
		k := testAccProvider.Meta().(string)

		e := fmt.Sprintf("/experiments/%s", experimentID)
		r, err := makeAPICall(k, e, "DELETE", nil)

		if err != nil {
			return err
		}

		// Should 404 as the experiment was already deleted
		if r.StatusCode != 404 {
			return errors.New(fmt.Sprintf("Status %v, Message: %s, Errors: %v", r.StatusCode, r.Message, r.Errors))
		}
	}

	return nil
}

func testAccCheckExperimentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ExperimentID set")
		}

		sdkKey, ok := os.LookupEnv("statsig_console_key")

		e := fmt.Sprintf("/experiments/%s", rs.Primary.ID)
		r, err := makeAPICall(sdkKey, e, "GET", nil)

		if err != nil {
			return err
		}

		data := r.Data.(map[string]interface{})
		groups := data["groups"].([]interface{})
		if len(groups) != 2 {
			return fmt.Errorf("invalid group size")
		}

		aGroup := groups[0].(map[string]interface{})
		paramValues := aGroup["parameterValues"].(map[string]interface{})
		if paramValues["a_bool"] == nil || paramValues["a_string"] == nil {
			return fmt.Errorf("invalid experiment group")
		}

		return nil
	}
}
