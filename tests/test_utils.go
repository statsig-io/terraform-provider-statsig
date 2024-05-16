package tests

import (
	"context"
	"errors"
	"os"
	"testing"

	new "terraform-provider-statsig/internal"
	old "terraform-provider-statsig/statsig"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func protoV6ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
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
	}
}

func testAccPreCheck(t *testing.T) {
	_, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		t.Fatal("statsig_console_key env var not provided")
	}
}

func getImportStateIDFunc(name string, attr string, res *string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		resource, ok := s.RootModule().Resources[name]
		if !ok || resource == nil {
			return "", errors.New("Resource not found")
		}
		*res = resource.Primary.Attributes[attr]
		return *res, nil
	}
}
