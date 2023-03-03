package statsig

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGateFull(t *testing.T) {
	fullGate, _ := os.ReadFile("test_resources/gate_full.tf")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckGateDestroy,
		Steps: []resource.TestStep{
			{
				Config: string(fullGate),
				Check: resource.ComposeTestCheckFunc(
					verifyFullGateSetup(t, "statsig_gate.my_gate"),
					verifyFullGateOutput(t),
				),
			},
			{
				ImportStateId: "my_gate",
				ImportState:   true,
				ResourceName:  "statsig_gate.my_gate",
			},
			{
				Config: string(fullGate),
				Check: resource.ComposeTestCheckFunc(
					verifyFullGateSetup(t, "statsig_gate.my_gate"),
					verifyFullGateOutput(t),
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

		ctx := context.Background()

		e := fmt.Sprintf("/gates/%s", gateID)
		r, err := makeAPICall(ctx, k, e, "DELETE", nil)

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

func verifyFullGateOutput(t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		o, _ := s.RootModule().Outputs["my_gate"]
		rs, _ := s.RootModule().Resources["statsig_gate.my_gate"]
		local := rs.Primary.Attributes

		conditions := o.Value.(map[string]interface{})["rules"].([]interface{})[0].(map[string]interface{})["conditions"].([]interface{})

		assert.Equal(t, "public", conditions[0].(map[string]interface{})["type"])
		assert.Equal(t, "user_id", conditions[1].(map[string]interface{})["type"])
		assert.Equal(t, "email", conditions[2].(map[string]interface{})["type"])
		assert.Equal(t, "custom_field", conditions[3].(map[string]interface{})["type"])
		assert.Equal(t, "app_version", conditions[4].(map[string]interface{})["type"])
		assert.Equal(t, "browser_name", conditions[5].(map[string]interface{})["type"])
		assert.Equal(t, "browser_version", conditions[6].(map[string]interface{})["type"])
		assert.Equal(t, "os_name", conditions[7].(map[string]interface{})["type"])
		assert.Equal(t, "os_version", conditions[8].(map[string]interface{})["type"])
		assert.Equal(t, "country", conditions[9].(map[string]interface{})["type"])
		assert.Equal(t, "passes_gate", conditions[10].(map[string]interface{})["type"])
		assert.Equal(t, "fails_gate", conditions[11].(map[string]interface{})["type"])
		assert.Equal(t, "time", conditions[12].(map[string]interface{})["type"])
		assert.Equal(t, "environment_tier", conditions[13].(map[string]interface{})["type"])
		assert.Equal(t, "passes_segment", conditions[14].(map[string]interface{})["type"])
		assert.Equal(t, "fails_segment", conditions[15].(map[string]interface{})["type"])
		assert.Equal(t, "ip_address", conditions[16].(map[string]interface{})["type"])

		for index, elem := range conditions {
			condition := elem.(map[string]interface{})

			assert.Equal(t, fmt.Sprint(index), fmt.Sprintf("%s", condition["index"]))
			assert.Equal(t, fmt.Sprint(index), local[fmt.Sprintf("rules.0.conditions.%d.index", index)])
		}

		return nil
	}
}

func verifyFullGateSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		remote, err := getGateDataFromServer(rs.Primary.ID)
		local := rs.Primary.Attributes
		if err != nil {
			return err
		}

		assert.Equal(t, "my_gate", remote["id"])
		assert.Equal(t, "my_gate", local["id"])

		assert.Equal(t, "my_gate", local["name"])

		assert.Equal(t, "A short description of what this Gate is used for.", local["description"])
		assert.Equal(t, "A short description of what this Gate is used for.", remote["description"])

		assert.Equal(t, "userID", local["id_type"])
		assert.Equal(t, "userID", remote["idType"])

		rule := remote["rules"].([]interface{})[0].(map[string]interface{})
		conditions := rule["conditions"].([]interface{})

		assert.Equal(t, fmt.Sprintf("%d", len(conditions)), local["rules.0.conditions.#"])

		assert.Equal(t, "public", local["rules.0.conditions.0.type"])
		assert.Equal(t, "", local["rules.0.conditions.0.target_value"])
		assert.Equal(t, "", local["rules.0.conditions.0.operator"])
		assert.Equal(t, "", local["rules.0.conditions.0.field"])

		public := conditions[0].(map[string]interface{})
		assert.Nil(t, public["target_value"])
		assert.Nil(t, public["operator"])
		assert.Nil(t, public["field"])

		assert.Equal(t, "user_id", local["rules.0.conditions.1.type"])
		assert.Equal(t, "1", local["rules.0.conditions.1.target_value.0"])
		assert.Equal(t, "2", local["rules.0.conditions.1.target_value.1"])
		assert.Equal(t, "any", local["rules.0.conditions.1.operator"])
		assert.Equal(t, "", local["rules.0.conditions.1.field"])

		userID := conditions[1].(map[string]interface{})
		assert.Equal(t, []interface{}{"1", "2"}, userID["targetValue"])
		assert.Equal(t, "any", userID["operator"])
		assert.Nil(t, userID["field"])

		return nil
	}
}

func getGateDataFromServer(gid string) (map[string]interface{}, error) {
	sdkKey, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		return nil, fmt.Errorf("no sdk key found")
	}

	endpoint := fmt.Sprintf("/gates/%s", gid)
	ctx := context.Background()
	response, err := makeAPICall(ctx, sdkKey, endpoint, "GET", nil)

	if err != nil {
		return nil, err
	}

	data := response.Data.(map[string]interface{})
	return data, nil
}
