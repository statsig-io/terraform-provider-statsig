package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccDynamicConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/dynamic_config_basic.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyDynamicConfigSetup(t, "statsig_dynamic_config.example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/dynamic_config_basic.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyDynamicConfigSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "test_dynamic_config", local["id"])
		assert.Equal(t, "test_dynamic_config", local["name"])
		assert.Equal(t, "userID", local["id_type"])
		assert.Equal(t, "A test dynamic config", local["description"])
		assert.Equal(t, "true", local["is_enabled"])
		assert.Equal(t, "{}", local["default_value_json5"])

		return nil
	}
}

func TestAccDynamicConfigTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/dynamic_config_template.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyDynamicConfigSetup(t, "statsig_dynamic_config.example_template"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/dynamic_config_template.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccDynamicConfigFull(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/dynamic_config_full.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyDynamicConfigFullSetup(t, "statsig_dynamic_config.my_dynamic_config"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/dynamic_config_full.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyDynamicConfigFullSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "my_dynamic_config", local["id"])

		assert.Equal(t, "my_dynamic_config", local["name"])

		assert.Equal(t, "A short description of what this Dynamic Config is used for.", local["description"])

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

		return nil
	}
}
