package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/role.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyRoleSetup(t, "statsig_role.example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/role.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifyRoleSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "Custom Role", local["name"])
		assert.Equal(t, "false", local["permissions.invitation_access"])
		assert.Equal(t, "false", local["permissions.create_configs"])
		assert.Equal(t, "false", local["permissions.edit_or_delete_configs"])
		assert.Equal(t, "false", local["permissions.launch_to_production"])
		assert.Equal(t, "false", local["permissions.launch_or_disable_configs"])
		assert.Equal(t, "false", local["permissions.start_experiments"])
		assert.Equal(t, "false", local["permissions.create_or_edit_templates"])
		assert.Equal(t, "false", local["permissions.create_or_edit_dashboards"])
		assert.Equal(t, "false", local["permissions.create_teams"])
		assert.Equal(t, "false", local["permissions.edit_dynamic_config_schemas"])
		assert.Equal(t, "false", local["permissions.create_release_pipelines"])
		assert.Equal(t, "false", local["permissions.approve_required_review_release_pipeline_phase"])
		assert.Equal(t, "false", local["permissions.self_approve_review"])
		assert.Equal(t, "false", local["permissions.approve_reviews"])
		assert.Equal(t, "false", local["permissions.bypass_reviews_for_overrides"])
		assert.Equal(t, "false", local["permissions.bypass_precommit_webhook"])
		assert.Equal(t, "false", local["permissions.metric_management"])
		assert.Equal(t, "false", local["permissions.verify_metrics"])
		assert.Equal(t, "false", local["permissions.use_metrics_explorer"])
		assert.Equal(t, "false", local["permissions.local_metrics"])
		assert.Equal(t, "false", local["permissions.manage_alerts"])
		assert.Equal(t, "false", local["permissions.integrations_edit_access"])
		assert.Equal(t, "false", local["permissions.source_connection_and_creation"])
		assert.Equal(t, "false", local["permissions.data_warehouse_ingestion_and_exports_edit_access"])
		assert.Equal(t, "false", local["permissions.edit_and_tag_configs_with_core_tag"])
		assert.Equal(t, "false", local["permissions.reset_experiments"])
		assert.Equal(t, "false", local["permissions.event_dimensions_access"])
		assert.Equal(t, "false", local["permissions.manual_whn_reload"])
		assert.Equal(t, "false", local["permissions.whn_connection"])

		return nil
	}
}
