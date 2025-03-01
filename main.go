package main

import (
	"context"
	"flag"
	"log"
	provider "terraform-provider-statsig/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address:         "registry.terraform.io/statsig/console_api",
		Debug:           debug,
		ProtocolVersion: 6,
	}

	err := providerserver.Serve(
		context.Background(),
		provider.New,
		opts,
	)

	if err != nil {
		log.Fatal(err)
	}
}
