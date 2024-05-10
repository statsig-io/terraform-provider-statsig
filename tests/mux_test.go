package tests

import (
	"context"
	"os"
	"testing"

	new "terraform-provider-statsig/internal"
	old "terraform-provider-statsig/statsig"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type TestSteps struct {
	steps []resource.TestStep
}

func (t *TestSteps) Append(in ...[]resource.TestStep) {
	for _, test := range in {
		if test == nil {
			continue
		}

		t.steps = append(t.steps, test...)
	}
}

func BuildTestSteps(t *testing.T) []resource.TestStep {
	testSteps := TestSteps{
		steps: make([]resource.TestStep, 0),
	}
	testSteps.Append(TestAccServerKey(t), TestAccClientKey(t), TestAccConsoleKey(t))
	// testSteps.Append(TestAccServerKey(t))
	return testSteps.steps
}

func TestMuxServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"statsig": func() (tfprotov6.ProviderServer, error) {
				ctx := context.Background()

				oldProvider := old.Provider() // terraform-plugin-sdk provider
				newProvider := new.New()      // terraform-plugin-framework provider

				upgradedSdkServer, err := tf5to6server.UpgradeServer(
					ctx,
					oldProvider.GRPCProvider,
				)

				if err != nil {
					return nil, err
				}

				providers := []func() tfprotov6.ProviderServer{
					providerserver.NewProtocol6(newProvider),
					func() tfprotov6.ProviderServer {
						return upgradedSdkServer
					},
				}

				muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)

				if err != nil {
					return nil, err
				}

				return muxServer.ProviderServer(), nil
			},
		},
		PreCheck: func() { testAccPreCheck(t) },
		Steps:    BuildTestSteps(t), // TODO: migrate tests for resources still using old provider
	})
}

func testAccPreCheck(t *testing.T) {
	_, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		t.Fatal("statsig_console_key env var not provided")
	}
}
