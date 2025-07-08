package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSettingsProject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/settings_project.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifySettingsProjectSetup(t, "statsig_settings_project.example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/settings_project.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifySettingsProjectSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "Console API Test", local["name"])
		assert.Equal(t, "CLOSED", local["visibility"])
		assert.Equal(t, "user_id", local["default_unit_type"])

		return nil
	}
}

func TestAccSettingsProjectSparse(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/settings_project_sparse.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifySettingsProjectSparseSetup(t, "statsig_settings_project.sparse_example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/settings_project_sparse.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifySettingsProjectSparseSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "Console API Test", local["name"])
		assert.Equal(t, "CLOSED", local["visibility"])

		return nil
	}
}
