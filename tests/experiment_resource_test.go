package tests

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccExperimentFull(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/experiment_full.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyFullExperimentSetup(t, "statsig_experiment.full_experiment"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/experiment_full.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccExperimentUpdating(t *testing.T) {
	key := "statsig_experiment.my_experiment"

	var testGroupID string
	var controlGroupID string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/experiment_basic.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(key, "id", "my_experiment"),
					resource.TestCheckResourceAttr(key, "status", "setup"),
					testAccExtractResourceAttr(key, "groups.0.id", &testGroupID),
					testAccExtractResourceAttr(key, "groups.1.id", &controlGroupID),
				),
			},
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_test_group_id", testGroupID)
					os.Setenv("TF_VAR_control_group_id", controlGroupID)
				},
				ConfigFile: config.StaticFile("test_resources/experiment_active.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(key, "id", "my_experiment"),
					resource.TestCheckResourceAttr(key, "status", "active"),
					testAccExtractResourceAttr(key, "groups.0.id", &testGroupID),
					testAccExtractResourceAttr(key, "groups.1.id", &controlGroupID),
				),
			},
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_test_group_id", testGroupID)
					os.Setenv("TF_VAR_control_group_id", controlGroupID)
				},
				ConfigFile: config.StaticFile("test_resources/experiment_decision_made.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyShippedExperimentSetup(t, key, &testGroupID),
				),
			},
		},
	})
}

func verifyShippedExperimentSetup(t *testing.T, name string, launchedGroupID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "my_experiment", local["id"])

		assert.Equal(t, "100", local["allocation"])

		assert.Equal(t, "decision_made", local["status"])

		assert.Equal(t, *launchedGroupID, local["launched_group_id"])

		return nil
	}
}

func verifyFullExperimentSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "full_experiment", local["id"])

		assert.Equal(t, "full_experiment", local["name"])

		assert.Equal(t, "A short description of what we are experimenting on.", local["description"])

		assert.Equal(t, "userID", local["id_type"])

		assert.Equal(t, "12.3", local["allocation"])

		assert.Equal(t, "setup", local["status"])

		assert.Equal(t, "80", local["default_confidence_interval"])

		assert.Equal(t, "true", local["bonferroni_correction"])

		assert.Equal(t, "10", local["duration"])

		assert.Equal(t, "a_layer", local["layer_id"])

		assert.Equal(t, "targeting_gate", local["targeting_gate_id"])

		assert.Equal(t, "", local["launched_group_id"])

		assert.Equal(t, "test-tag-a", local["tags.0"])
		assert.Equal(t, "test-tag-b", local["tags.1"])

		assert.Equal(t, "0", local["primary_metric_tags.#"])

		assert.Equal(t, "0", local["secondary_metric_tags.#"])

		assert.Equal(t, "user", local["primary_metrics.0.type"])
		assert.Equal(t, "d1_retention_rate", local["primary_metrics.0.name"])

		assert.Equal(t, "user", local["secondary_metrics.0.type"])
		assert.Equal(t, "dau", local["secondary_metrics.0.name"])

		assert.Equal(t, "user", local["secondary_metrics.1.type"])
		assert.Equal(t, "new_dau", local["secondary_metrics.1.name"])

		assert.Equal(t, "Test A", local["groups.0.name"])
		assert.Equal(t, "33.3", local["groups.0.size"])
		assert.Equal(t, "test_a", local["groups.0.parameter_values.a_string"])

		assert.Equal(t, "Test B", local["groups.1.name"])
		assert.Equal(t, "33.3", local["groups.1.size"])
		assert.Equal(t, "test_b", local["groups.1.parameter_values.a_string"])

		assert.Equal(t, "Control", local["groups.2.name"])
		assert.Equal(t, "33.4", local["groups.2.size"])
		assert.Equal(t, "control", local["groups.2.parameter_values.a_string"])

		return nil
	}
}
