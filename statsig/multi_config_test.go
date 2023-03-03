package statsig

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccMultiConfig(t *testing.T) {
	multiConfig, _ := os.ReadFile("test_resources/multi_config.tf")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckGateDestroy,
		Steps: []resource.TestStep{
			{
				Config: string(multiConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("statsig_gate.simple_gate_a", "id", "simple_gate_a"),
					resource.TestCheckResourceAttr("statsig_gate.simple_gate_b", "id", "simple_gate_b"),
					resource.TestCheckResourceAttr("statsig_experiment.simple_experiment", "id", "simple_experiment"),
					verifyMultiConfigSetup(t),
				),
			},
			{
				Config:             string(multiConfig),
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("statsig_gate.simple_gate_a", "id", "simple_gate_a"),
					resource.TestCheckResourceAttr("statsig_gate.simple_gate_b", "id", "simple_gate_b"),
					resource.TestCheckResourceAttr("statsig_experiment.simple_experiment", "id", "simple_experiment"),
					verifyMultiConfigSetup(t),
				),
			},
		},
	})
}

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
