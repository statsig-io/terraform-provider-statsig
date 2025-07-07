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
