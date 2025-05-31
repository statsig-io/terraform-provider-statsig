package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccEntityPropertyFull(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{isWHN: true}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/entity_property_full.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyEntityPropertySetup(t, "statsig_entity_property.example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/entity_property_full.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyEntityPropertySetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "my_entity_property", local["name"])

		assert.Equal(t, "A short description of this entity property.", local["description"])

		assert.Equal(t, "false", local["is_read_only"])

		assert.Equal(t, "SELECT * FROM `shoppy-sales.kenny_dev.users`", local["sql"])

		assert.Equal(t, "1", local["id_type_mapping.#"])
		assert.Equal(t, "userID", local["id_type_mapping.0.statsig_unit_id"])
		assert.Equal(t, "user_id", local["id_type_mapping.0.column"])

		assert.Equal(t, "ts", local["timestamp_column"])

		assert.Equal(t, "1", local["tags.#"])
		assert.Equal(t, "test-tag", local["tags.0"])

		assert.Equal(t, "SDK_KEY", local["owner.owner_type"])

		return nil
	}
}
