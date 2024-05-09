package tests

import (
	"context"
	"testing"

	new "terraform-provider-statsig/internal"
	old "terraform-provider-statsig/statsig"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
		Steps: []resource.TestStep{}, // TODO: migrate tests for resources still using old provider
	})
}
