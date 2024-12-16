package statsig

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("experiment_resource", &resource.Sweeper{
		Name: "experiment_resource",
		F: func(region string) error {

			for _, s := range []string{"my_experiment", "simple_experiment", "full_experiment"} {
				_, err := deleteExperiment(s)
				if err != nil {
					return err
				}
			}

			return nil
		},
	})
}

func TestAccExperimentFull(t *testing.T) {
	fullExperiment, _ := os.ReadFile("test_resources/experiment_full.tf")

	key := "statsig_experiment.full_experiment"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      verifyExperimentDestroyed,
		Steps: []resource.TestStep{
			{
				Config: string(fullExperiment),
				Check: resource.ComposeTestCheckFunc(
					verifyFullExperimentSetup(t, key),
				),
			},
		},
	})
}

func TestAccExperimentUpdating(t *testing.T) {
	basicExperiment, _ := os.ReadFile("test_resources/experiment_basic.tf")
	activeExperiment, _ := os.ReadFile("test_resources/experiment_active.tf")
	shippedExperiment, _ := os.ReadFile("test_resources/experiment_decision_made.tf")

	key := "statsig_experiment.my_experiment"

	var groupIDToLaunch string

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      verifyExperimentDestroyed,
		Steps: []resource.TestStep{
			{
				Config: string(basicExperiment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(key, "id", "my_experiment"),
					resource.TestCheckResourceAttr(key, "status", "setup"),
					extractGroupID(key, &groupIDToLaunch),
				),
			},
			{
				Config: string(activeExperiment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(key, "id", "my_experiment"),
					resource.TestCheckResourceAttr(key, "status", "active"),
					extractGroupID(key, &groupIDToLaunch),
				),
			},
			{
				PreConfig: func() {
					os.Setenv("TF_VAR_launched_group_id", groupIDToLaunch)
				},
				Config: string(shippedExperiment),
				Check: resource.ComposeTestCheckFunc(
					verifyShippedExperimentSetup(t, key, &groupIDToLaunch),
				),
			},
		},
	})
}

func verifyExperimentDestroyed(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "statsig_experiment" {
			continue
		}

		experimentID := rs.Primary.ID

		r, err := deleteExperiment(experimentID)

		if err != nil {
			return err
		}

		// Should 404 as the experiment was already deleted
		if r.StatusCode != 404 {
			return errors.New(fmt.Sprintf("Status %v, Message: %s, Errors: %v", r.StatusCode, r.Message, r.Errors))
		}
	}

	return nil
}

func deleteExperiment(name string) (*APIResponse, error) {
	k, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		panic("statsig_console_key env var not provided")
	}

	e := fmt.Sprintf("/experiments/%s", name)
	ctx := context.Background()
	return makeAPICall(ctx, k, e, "DELETE", nil)
}

func extractGroupID(name string, out *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		*out = rs.Primary.Attributes["groups.0.id"]

		return nil
	}
}

func verifyShippedExperimentSetup(t *testing.T, name string, launchedGroupID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		remote, err := getExperimentDataFromServer(rs.Primary.ID)
		local := rs.Primary.Attributes
		if err != nil {
			return err
		}

		assert.Equal(t, "my_experiment", remote["id"])
		assert.Equal(t, "my_experiment", local["id"])

		assert.Equal(t, 100.0, remote["allocation"])
		assert.Equal(t, "100", local["allocation"])

		assert.Equal(t, "decision_made", remote["status"])
		assert.Equal(t, "decision_made", local["status"])

		assert.Equal(t, *launchedGroupID, remote["launchedGroupID"])
		assert.Equal(t, *launchedGroupID, local["launched_group_id"])

		return nil
	}
}

func verifyFullExperimentSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		remote, err := getExperimentDataFromServer(rs.Primary.ID)
		local := rs.Primary.Attributes
		if err != nil {
			return err
		}

		assert.Equal(t, "full_experiment", remote["id"])
		assert.Equal(t, "full_experiment", local["id"])

		assert.Equal(t, "full_experiment", local["name"])

		assert.Equal(t, "A short description of what we are experimenting on.", local["description"])
		assert.Equal(t, "A short description of what we are experimenting on.", remote["description"])

		assert.Equal(t, "userID", local["id_type"])
		assert.Equal(t, "userID", remote["idType"])

		assert.Equal(t, "12.3", local["allocation"])
		assert.Equal(t, 12.3, remote["allocation"])

		assert.Equal(t, "setup", local["status"])
		assert.Equal(t, "setup", remote["status"])

		assert.Equal(t, "80", local["default_confidence_interval"])
		assert.Equal(t, "80", remote["defaultConfidenceInterval"])

		assert.Equal(t, "true", local["bonferroni_correction"])
		assert.Equal(t, true, remote["bonferroniCorrection"])

		assert.Equal(t, "10", local["duration"])
		assert.Equal(t, 10.0, remote["duration"])

		assert.Equal(t, "a_layer", local["layer_id"])
		assert.Equal(t, "a_layer", remote["layerID"])

		assert.Equal(t, "targeting_gate", local["targeting_gate_id"])
		assert.Equal(t, "targeting_gate", remote["targetingGateID"])

		assert.Equal(t, "", local["launched_group_id"])
		assert.Equal(t, nil, remote["launchedGroupID"])

		assert.Equal(t, "test-tag-a", local["tags.0"])
		assert.Equal(t, "test-tag-b", local["tags.1"])
		assert.Equal(t, []interface{}{"test-tag-a", "test-tag-b"}, remote["tags"])

		assert.Equal(t, "0", local["primary_metric_tags.#"])
		assert.Equal(t, []interface{}{}, remote["primaryMetricTags"])

		assert.Equal(t, "0", local["secondary_metric_tags.#"])
		assert.Equal(t, []interface{}{}, remote["secondaryMetricTags"])

		primary := jsonStringToArray(local["primary_metrics_json"])[0].(map[string]interface{})
		assert.Equal(t, "user", primary["type"])
		assert.Equal(t, "d1_retention_rate", primary["name"])

		core := jsonStringToArray(local["secondary_metrics_json"])[0].(map[string]interface{})
		assert.Equal(t, "user", core["type"])
		assert.Equal(t, "dau", core["name"])

		secondary := jsonStringToArray(local["secondary_metrics_json"])[1].(map[string]interface{})
		assert.Equal(t, "user", secondary["type"])
		assert.Equal(t, "new_dau", secondary["name"])

		groups := remote["groups"].([]interface{})
		assert.Equal(t, 3, len(groups))

		checkA, checkB, checkControl := false, false, false

		for index, elem := range groups {
			group := elem.(map[string]interface{})
			switch group["name"] {
			case "Test A":
				checkA = true
				assert.Equal(t, 33.3, group["size"])
				assert.Equal(t, "33.3", local[fmt.Sprintf("groups.%d.size", index)])

				params := jsonStringToMap(local[fmt.Sprintf("groups.%d.parameter_values_json", index)])
				assert.Equal(t, "test_a", params["a_string"])
				break

			case "Test B":
				checkB = true
				assert.Equal(t, 33.3, group["size"])
				assert.Equal(t, "33.3", local[fmt.Sprintf("groups.%d.size", index)])

				params := jsonStringToMap(local[fmt.Sprintf("groups.%d.parameter_values_json", index)])
				assert.Equal(t, "test_b", params["a_string"])
				break

			case "Control":
				checkControl = true
				assert.Equal(t, 33.4, group["size"])
				assert.Equal(t, "33.4", local[fmt.Sprintf("groups.%d.size", index)])

				params := jsonStringToMap(local[fmt.Sprintf("groups.%d.parameter_values_json", index)])
				assert.Equal(t, "control", params["a_string"])
				break
			}
		}

		assert.True(t, checkA)
		assert.True(t, checkB)
		assert.True(t, checkControl)

		return nil
	}
}

func getExperimentDataFromServer(eid string) (map[string]interface{}, error) {
	sdkKey, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		return nil, fmt.Errorf("no sdk key found")
	}

	endpoint := fmt.Sprintf("/experiments/%s", eid)

	ctx := context.Background()
	response, err := makeAPICall(ctx, sdkKey, endpoint, "GET", nil)

	if err != nil {
		return nil, err
	}

	data := response.Data.(map[string]interface{})
	return data, nil
}
