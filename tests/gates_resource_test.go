package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccGateFull(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/gate_full.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyFullGateSetup(t, "statsig_gate.my_gate"),
					verifyFullGateOutput(t),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/gate_full.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyFullGateOutput(t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		o, _ := s.RootModule().Outputs["my_gate"]

		rules := o.Value.(map[string]interface{})["rules"].([]interface{})
		mainRule := rules[0].(map[string]interface{})
		devRule := rules[1].(map[string]interface{})

		assert.Equal(t, "All Conditions", mainRule["name"])
		assert.Equal(t, "10", fmt.Sprint(mainRule["pass_percentage"]))

		assert.Equal(t, "Development Conditions", devRule["name"])
		assert.Equal(t, "10", fmt.Sprint(devRule["pass_percentage"]))
		assert.Equal(t, "development", devRule["environments"].([]interface{})[0])

		conditions := mainRule["conditions"].([]interface{})
		devConditions := devRule["conditions"].([]interface{})

		assert.Equal(t, "public", conditions[0].(map[string]interface{})["type"])
		assert.Equal(t, "user_id", conditions[1].(map[string]interface{})["type"])
		assert.Equal(t, "email", conditions[2].(map[string]interface{})["type"])
		assert.Equal(t, "custom_field", conditions[3].(map[string]interface{})["type"])
		assert.Equal(t, "app_version", conditions[4].(map[string]interface{})["type"])
		assert.Equal(t, "browser_name", conditions[5].(map[string]interface{})["type"])
		assert.Equal(t, "browser_version", conditions[6].(map[string]interface{})["type"])
		assert.Equal(t, "os_name", conditions[7].(map[string]interface{})["type"])
		assert.Equal(t, "os_version", conditions[8].(map[string]interface{})["type"])
		assert.Equal(t, "country", conditions[9].(map[string]interface{})["type"])
		assert.Equal(t, "passes_gate", conditions[10].(map[string]interface{})["type"])
		assert.Equal(t, "fails_gate", conditions[11].(map[string]interface{})["type"])
		assert.Equal(t, "time", conditions[12].(map[string]interface{})["type"])
		assert.Equal(t, "environment_tier", conditions[13].(map[string]interface{})["type"])
		assert.Equal(t, "passes_segment", conditions[14].(map[string]interface{})["type"])
		assert.Equal(t, "fails_segment", conditions[15].(map[string]interface{})["type"])
		assert.Equal(t, "ip_address", conditions[16].(map[string]interface{})["type"])
		assert.Equal(t, "public", devConditions[0].(map[string]interface{})["type"])

		return nil
	}
}

func verifyFullGateSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "my_gate", local["id"])

		assert.Equal(t, "my_gate", local["name"])

		assert.Equal(t, "A short description of what this Gate is used for.", local["description"])

		assert.Equal(t, "userID", local["id_type"])

		assert.Equal(t, "public", local["rules.0.conditions.0.type"])
		assert.Equal(t, "", local["rules.0.conditions.0.target_value"])
		assert.Equal(t, "", local["rules.0.conditions.0.operator"])
		assert.Equal(t, "", local["rules.0.conditions.0.field"])

		assert.Equal(t, "user_id", local["rules.0.conditions.1.type"])
		assert.Equal(t, "1", local["rules.0.conditions.1.target_value.0"])
		assert.Equal(t, "2", local["rules.0.conditions.1.target_value.1"])
		assert.Equal(t, "any", local["rules.0.conditions.1.operator"])
		assert.Equal(t, "", local["rules.0.conditions.1.field"])

		assert.Equal(t, "true", local["measure_metric_lifts"])

		return nil
	}
}
