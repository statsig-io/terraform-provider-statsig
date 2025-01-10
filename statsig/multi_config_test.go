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

// TODO: Migrate test after finished migrating experiment resource to new provider
// func TestAccMultiConfig(t *testing.T) {
// 	multiConfig, _ := os.ReadFile("test_resources/multi_config.tf")

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: testAccProviderFactories,
// 		CheckDestroy:      testAccCheckGateDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: string(multiConfig),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("statsig_gate.simple_gate_a", "id", "simple_gate_a"),
// 					resource.TestCheckResourceAttr("statsig_gate.simple_gate_b", "id", "simple_gate_b"),
// 					resource.TestCheckResourceAttr("statsig_experiment.simple_experiment", "id", "simple_experiment"),
// 					verifyMultiConfigSetup(t),
// 				),
// 			},
// 			{
// 				Config:             string(multiConfig),
// 				ExpectNonEmptyPlan: false,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("statsig_gate.simple_gate_a", "id", "simple_gate_a"),
// 					resource.TestCheckResourceAttr("statsig_gate.simple_gate_b", "id", "simple_gate_b"),
// 					resource.TestCheckResourceAttr("statsig_experiment.simple_experiment", "id", "simple_experiment"),
// 					verifyMultiConfigSetup(t),
// 				),
// 			},
// 		},
// 	})
// }

func verifyMultiConfigSetup(t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		root := s.RootModule()
		gatesOs, _ := root.Outputs["gates"]
		experimentsOs, _ := root.Outputs["experiments"]

		assert.Equal(t, 2, len(gatesOs.Value.([]interface{})))
		assert.Equal(t, 1, len(experimentsOs.Value.([]interface{})))

		return nil
	}
}

func testAccCheckGateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "statsig_gate" {
			continue
		}

		gateID := rs.Primary.ID
		r, err := deleteGate(gateID)

		if err != nil {
			return err
		}

		// Should 404 as the gate was already deleted
		if r.StatusCode != 404 {
			return errors.New(fmt.Sprintf("Status %v, Message: %s, Errors: %v", r.StatusCode, r.Message, r.Errors))
		}
	}

	return nil
}

func deleteGate(name string) (*APIResponse, error) {
	k, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		panic("statsig_console_key env var not provided")
	}

	e := fmt.Sprintf("/gates/%s", name)
	ctx := context.Background()
	return makeAPICall(ctx, k, e, "DELETE", nil)
}
