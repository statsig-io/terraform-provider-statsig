package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccEventCountMetric(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/metric_event_count.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyEventCountMetricSetup(t, "statsig_metric.custom_event_count_metric"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/metric_event_count.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyEventCountMetricSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "A short description of this metric.", local["description"])
		assert.Equal(t, "14", local["custom_roll_up_end"])
		assert.Equal(t, "0", local["custom_roll_up_start"])
		assert.Equal(t, "decrease", local["directionality"])
		assert.Equal(t, "false", local["is_permanent"])
		assert.Equal(t, "false", local["is_read_only"])
		assert.Equal(t, "false", local["is_verified"])
		assert.Equal(t, "1", local["metric_events.#"])
		assert.Equal(t, "test_event_1", local["metric_events.0.name"])
		assert.Equal(t, "0", local["metric_events.0.criteria.#"])
		assert.Equal(t, "Custom Event Count Metric", local["name"])
		assert.Equal(t, "custom", local["rollup_time_window"])
		assert.Equal(t, "1", local["tags.#"])
		assert.Equal(t, "test-tag", local["tags.0"])
		assert.Equal(t, "event_count_custom", local["type"])
		assert.Equal(t, "1", local["unit_types.#"])
		assert.Equal(t, "userID", local["unit_types.0"])

		return nil
	}
}

func TestAccWarehouseNativeMetric(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{isWHN: true}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/metric_warehouse_native.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyWarehouseNativeMetricSetup(t, "statsig_metric.warehouse_native_metric"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/metric_warehouse_native.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyWarehouseNativeMetricSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "A short description of this metric.", local["description"])
		assert.Equal(t, "increase", local["directionality"])
		assert.Equal(t, "false", local["is_permanent"])
		assert.Equal(t, "false", local["is_read_only"])
		assert.Equal(t, "false", local["is_verified"])
		assert.Equal(t, "shoppy_events", local["warehouse_native.metric_source_name"])
		assert.Equal(t, "count", local["warehouse_native.aggregation"])
		assert.Equal(t, "1", local["warehouse_native.criteria.#"])
		assert.Equal(t, "metadata", local["warehouse_native.criteria.0.type"])
		assert.Equal(t, "event", local["warehouse_native.criteria.0.column"])
		assert.Equal(t, "=", local["warehouse_native.criteria.0.condition"])
		assert.Equal(t, "1", local["warehouse_native.criteria.0.values.#"])
		assert.Equal(t, "add_to_cart", local["warehouse_native.criteria.0.values.0"])
		assert.Equal(t, "7", local["warehouse_native.cuped_attribution_window"])
		assert.Equal(t, "150", local["warehouse_native.cap"])
		assert.Equal(t, "Warehouse Native Metric", local["name"])
		assert.Equal(t, "1", local["tags.#"])
		assert.Equal(t, "test-tag", local["tags.0"])
		assert.Equal(t, "user_warehouse", local["type"])
		assert.Equal(t, "1", local["unit_types.#"])
		assert.Equal(t, "userID", local["unit_types.0"])

		return nil
	}
}
