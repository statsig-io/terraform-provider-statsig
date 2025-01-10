package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccExperimentFull_MUX(t *testing.T) {
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

func TestAccExperimentUpdating_MUX(t *testing.T) {
	key := "statsig_experiment.my_experiment"

	var groupIDToLaunch string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/experiment_basic.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(key, "id", "my_experiment"),
					resource.TestCheckResourceAttr(key, "status", "setup"),
					testAccExtractResourceAttr(key, "groups.0.id", &groupIDToLaunch),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/experiment_active.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(key, "id", "my_experiment"),
					resource.TestCheckResourceAttr(key, "status", "active"),
					testAccExtractResourceAttr(key, "groups.0.id", &groupIDToLaunch),
				),
			},
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_launched_group_id", groupIDToLaunch)
				},
				ConfigFile: config.StaticFile("test_resources/experiment_decision_made.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifyShippedExperimentSetup(t, key, &groupIDToLaunch),
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

		primary := jsonStringToArray(local["primary_metrics_json"])[0].(map[string]interface{})
		assert.Equal(t, "user", primary["type"])
		assert.Equal(t, "d1_retention_rate", primary["name"])

		core := jsonStringToArray(local["secondary_metrics_json"])[0].(map[string]interface{})
		assert.Equal(t, "user", core["type"])
		assert.Equal(t, "dau", core["name"])

		secondary := jsonStringToArray(local["secondary_metrics_json"])[1].(map[string]interface{})
		assert.Equal(t, "user", secondary["type"])
		assert.Equal(t, "new_dau", secondary["name"])

		assert.Equal(t, "Test A", local["groups.0.name"])
		assert.Equal(t, "33.3", local["groups.0.size"])
		params := jsonStringToMap(local["groups.0.parameter_values_json"])
		assert.Equal(t, "test_a", params["a_string"])

		assert.Equal(t, "Test B", local["groups.1.name"])
		assert.Equal(t, "33.3", local["groups.1.size"])
		params = jsonStringToMap(local["groups.1.parameter_values_json"])
		assert.Equal(t, "test_b", params["a_string"])

		assert.Equal(t, "Control", local["groups.2.name"])
		assert.Equal(t, "33.4", local["groups.2.size"])
		params = jsonStringToMap(local["groups.2.parameter_values_json"])
		assert.Equal(t, "control", params["a_string"])

		return nil
	}
}

func jsonStringToMap(in interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	value, ok := in.(string)
	if !ok {
		return result
	}

	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return map[string]interface{}{}
	}

	return result
}

func jsonStringToArray(in interface{}) []interface{} {
	var result []interface{}

	value, ok := in.(string)
	if !ok {
		return result
	}

	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return []interface{}{}
	}

	return result
}
