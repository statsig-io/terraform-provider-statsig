package tests

import (
	"errors"
	"os"
	"testing"

	provider "terraform-provider-statsig/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func protoV6ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"statsig": providerserver.NewProtocol6WithError(provider.New()),
	}
}

func testAccPreCheck(t *testing.T) {
	_, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		t.Fatal("statsig_console_key env var not provided")
	}
}

func testAccExtractResourceAttr(name string, attr string, res *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resources[name]
		if !ok || resource == nil {
			return errors.New("Resource not found")
		}
		*res = resource.Primary.Attributes[attr]
		return nil
	}
}
