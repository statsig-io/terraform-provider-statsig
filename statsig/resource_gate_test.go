package statsig

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGateBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckGateDestroy,
		Steps: []resource.TestStep{
			{
				Config: basicGate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGateExists("statsig_gate.my_gate"),
				),
			},
		},
	})
}

func testAccCheckGateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "statsig_gate" {
			continue
		}

		gateID := rs.Primary.ID
		k := testAccProvider.Meta().(string)

		e := fmt.Sprintf("/gates/%s", gateID)
		r, err := makeAPICall(k, e, "DELETE", nil)

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

func testAccCheckGateExists(n string) resource.TestCheckFunc {
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

const basicGate = `
resource "statsig_gate" "my_gate" {
  name        = "my_gate"
  description = "A short description of what this Gate is used for."
  is_enabled  = true
  id_type     = "userID"
  rules {
    name            = "All Conditions"
    pass_percentage = 10
    conditions {
      type = "public"
    }
    conditions {
      type         = "user_id"
      target_value = [
        "1", "2"
      ]
      operator = "any"
    }
    conditions {
      type         = "email"
      target_value = ["@outlook.com", "@gmail.com"]
      operator     = "str_contains_any"
    }
    conditions {
      type         = "custom_field"
      target_value = [31]
      operator     = "gt"
      field        = "age"
    }
    conditions {
      type         = "app_version"
      target_value = ["1.1.1"]
      operator     = "version_gt"
    }
    conditions {
      type         = "browser_name"
      target_value = ["Firefox", "Chrome"]
      operator     = "any"
    }
    conditions {
      type         = "browser_version"
      target_value = ["94.0.4606.81", "94.0.4606.92"]
      operator     = "any"
    }
    conditions {
      type         = "os_name"
      target_value = ["Android", "Windows"]
      operator     = "none"
    }
    conditions {
      type         = "os_version"
      target_value = ["11.0.0"]
      operator     = "version_lte"
    }
    conditions {
      type         = "country"
      target_value = ["NZ", "US"]
      operator     = "any"
    }
    conditions {
      type         = "passes_gate"
      target_value = ["my_gate_2"]
    }
    conditions {
      type         = "fails_gate"
      target_value = ["a_failing_gate"]
    }
    conditions {
      type         = "time"
      target_value = [1643070357193]
      operator     = "after"
    }
    conditions {
      type         = "environment_tier"
      target_value = ["production"]
      operator     = "any"
    }
    conditions {
      type         = "passes_segment"
      target_value = ["growth_org"]
    }
    conditions {
      type         = "fails_segment"
      target_value = ["promo_id_list"]
    }
    conditions {
      type         = "ip_address"
      target_value = ["1.1.1.1", "8.8.8.8"]
      operator     = "any"
    }
  }
  dev_rules {
    name            = "All Conditions"
    pass_percentage = 10
    conditions {
      type = "public"
    }
  }
}
`
