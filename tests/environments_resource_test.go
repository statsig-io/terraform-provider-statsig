package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccEnvironments(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/environments.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyEnvironmentsSetup(t, "statsig_environments.example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/environments.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyEnvironmentsSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "development",			local["environments.0.name"])
		assert.Equal(t, "0.35629335902367476",	local["environments.0.id"])
		assert.Equal(t, "false",				local["environments.0.is_production"])
		assert.Equal(t, "false",				local["environments.0.requires_review"])
		assert.Equal(t, "true",					local["environments.0.requires_release_pipeline"])
		
		assert.Equal(t, "staging",				local["environments.1.name"])
		assert.Equal(t, "0.015089163460661137", local["environments.1.id"])
		assert.Equal(t, "false",				local["environments.1.is_production"])
		assert.Equal(t, "false",				local["environments.1.requires_review"])
		assert.Equal(t, "true",					local["environments.1.requires_release_pipeline"])
		
		assert.Equal(t, "production",			local["environments.2.name"])
		assert.Equal(t, "0.4067426155658289",	local["environments.2.id"])
		assert.Equal(t, "true",					local["environments.2.is_production"])
		assert.Equal(t, "true",					local["environments.2.requires_review"])
		assert.Equal(t, "true",					local["environments.2.requires_release_pipeline"])
		
		return nil
	}
}
