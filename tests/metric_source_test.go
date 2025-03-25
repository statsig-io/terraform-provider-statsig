package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccMetricSourceQuery(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{isWHN: true}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/metric_source_query.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyMetricSourceBaseSetup(t, "statsig_metric_source.query"),
					verifyMetricSourceQuerySetup(t, "statsig_metric_source.query"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/metric_source_query.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccMetricSourceTable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{isWHN: true}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/metric_source_table.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyMetricSourceBaseSetup(t, "statsig_metric_source.table"),
					verifyMetricSourceTableSetup(t, "statsig_metric_source.table"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/metric_source_table.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyMetricSourceBaseSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "A short description of this metric source.", local["description"])

		assert.Equal(t, "false", local["is_read_only"])

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

func verifyMetricSourceQuerySetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "query_metric_source", local["name"])

		assert.Equal(t, "query", local["source_type"])

		assert.Equal(t, "SELECT * FROM `shoppy-sales.kenny_dev.shoppy-events`", local["sql"])

		return nil
	}
}

func verifyMetricSourceTableSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "table_metric_source", local["name"])

		assert.Equal(t, "table", local["source_type"])

		assert.Equal(t, "`shoppy-sales.kenny_dev.shoppy-events`", local["table_name"])

		assert.Equal(t, "1", local["custom_field_mapping.#"])
		assert.Equal(t, "price_usd_cents", local["custom_field_mapping.0.key"])
		assert.Equal(t, "price_usd/100", local["custom_field_mapping.0.formula"])

		return nil
	}
}
