package main

import (
	"context"
	"flag"
	"log"
	new "terraform-provider-statsig/internal"
	old "terraform-provider-statsig/statsig"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

func main() {
	ctx := context.Background()

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	oldProvider := old.Provider() // terraform-plugin-sdk provider
	newProvider := new.New()      // terraform-plugin-framework provider

	providers := []func() tfprotov5.ProviderServer{
		providerserver.NewProtocol5(newProvider),
		oldProvider.GRPCProvider,
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf5server.ServeOpt

	if debug {
		serveOpts = append(serveOpts, tf5server.WithManagedDebug())
	}

	err = tf5server.Serve(
		"registry.terraform.io/statsig/console_api",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
