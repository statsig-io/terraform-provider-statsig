package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccMultiConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/multi_config.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("statsig_gate.simple_gate_a", "id", "simple_gate_a"),
					resource.TestCheckResourceAttr("statsig_gate.simple_gate_b", "id", "simple_gate_b"),
					resource.TestCheckResourceAttr("statsig_experiment.simple_experiment", "id", "simple_experiment"),
					verifyMultiConfigSetup(t),
				),
			},
			{
				ConfigFile:         config.StaticFile("test_resources/multi_config.tf"),
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
