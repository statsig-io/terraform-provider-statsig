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

func TestAccSegmentRuleBased(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/segment_rule_based.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyRuleBasedSegmentSetup(t, "statsig_segment.rule_based_segment"),
					verifyRuleBasedSegmentOutput(t, "rule_based_segment"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/segment_rule_based.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyRuleBasedSegmentOutput(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		o := s.RootModule().Outputs[name]

		rules := o.Value.(map[string]interface{})["rules"].([]interface{})
		mainRule := rules[0].(map[string]interface{})

		assert.Equal(t, "All Conditions", mainRule["name"])
		assert.Equal(t, "10", fmt.Sprint(mainRule["pass_percentage"]))

		conditions := mainRule["conditions"].([]interface{})

		assert.Equal(t, "user_id", conditions[0].(map[string]interface{})["type"])
		assert.Equal(t, "email", conditions[1].(map[string]interface{})["type"])
		assert.Equal(t, "custom_field", conditions[2].(map[string]interface{})["type"])
		assert.Equal(t, "app_version", conditions[3].(map[string]interface{})["type"])
		assert.Equal(t, "browser_name", conditions[4].(map[string]interface{})["type"])
		assert.Equal(t, "browser_version", conditions[5].(map[string]interface{})["type"])
		assert.Equal(t, "os_name", conditions[6].(map[string]interface{})["type"])
		assert.Equal(t, "os_version", conditions[7].(map[string]interface{})["type"])
		assert.Equal(t, "country", conditions[8].(map[string]interface{})["type"])
		assert.Equal(t, "environment_tier", conditions[9].(map[string]interface{})["type"])
		assert.Equal(t, "passes_segment", conditions[10].(map[string]interface{})["type"])
		assert.Equal(t, "fails_segment", conditions[11].(map[string]interface{})["type"])
		assert.Equal(t, "ip_address", conditions[12].(map[string]interface{})["type"])

		return nil
	}
}

func verifyRuleBasedSegmentSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "rule_based_segment", local["id"])

		assert.Equal(t, "Rule Based Segment", local["name"])

		assert.Equal(t, "A short description of this rule based segment", local["description"])

		assert.Equal(t, "rule_based", local["type"])

		assert.Equal(t, "userID", local["id_type"])

		assert.Equal(t, "user_id", local["rules.0.conditions.0.type"])
		assert.Equal(t, "1", local["rules.0.conditions.0.target_value.0"])
		assert.Equal(t, "2", local["rules.0.conditions.0.target_value.1"])
		assert.Equal(t, "any", local["rules.0.conditions.0.operator"])
		assert.Equal(t, "", local["rules.0.conditions.0.field"])

		assert.Equal(t, "test-tag", local["tags.0"])
		return nil
	}
}
