package tests

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Tests that plan is empty after terraform refresh
func RefreshNoopPlanCheck() resource.TestStep {
	return resource.TestStep{
		RefreshState: true,
		RefreshPlanChecks: resource.RefreshPlanChecks{
			PostRefresh: []plancheck.PlanCheck{
				plancheck.ExpectEmptyPlan(),
			},
		},
	}
}

func DebugResourceCheckFunc(name string, localName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		o, _ := s.RootModule().Outputs[localName]
		local := rs.Primary.Attributes
		fmt.Printf("local: %+v\n", local)
		fmt.Printf("output: %+v\n", o)
		return nil
	}
}

var _ plancheck.PlanCheck = debugPlan{}

type debugPlan struct {
	step string
}

func (e debugPlan) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	rd, err := json.Marshal(req.Plan)
	if err != nil {
		fmt.Println("error marshalling machine-readable plan output:", err)
	}
	fmt.Printf("[%s] %s\n", e.step, string(rd))
}

func DebugPlan(step string) plancheck.PlanCheck {
	return debugPlan{step: step}
}
