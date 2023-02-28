package statsig

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExperimentBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: basicExperiment,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExperimentExists("statsig_experiment.my_experiment"),
				),
			},
		},
	})
}

func testAccCheckExperimentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "statsig_experiment" {
			continue
		}

		experimentID := rs.Primary.ID
		k := testAccProvider.Meta().(string)

		e := fmt.Sprintf("/experiments/%s", experimentID)
		r, err := makeAPICall(k, e, "DELETE", nil)

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

func testAccCheckExperimentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no OrderID set")
		}

		return nil
	}
}

const basicExperiment = `
resource "statsig_experiment" "my_experiment" {
  name        = "my_experiment"
  description = "A short description of what this Gate is used for."
  id_type     = "userID"
}
`
