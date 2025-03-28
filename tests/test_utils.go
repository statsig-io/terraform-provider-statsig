package tests

import (
	"errors"
	"fmt"
	"os"
	"testing"

	provider "terraform-provider-statsig/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type TestOptions struct {
	isWHN bool
}

func testAccProviders(t *testing.T, opts TestOptions) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"statsig": providerserver.NewProtocol6WithError(
			provider.NewTestProvider(getTestAPIKey(t, opts)),
		),
	}
}

func getTestAPIKey(t *testing.T, opts TestOptions) string {
	var apiKeyEnv string
	if opts.isWHN {
		apiKeyEnv = "statsig_whn_console_key"
	} else {
		apiKeyEnv = "statsig_console_key"
	}

	apiKey, ok := os.LookupEnv(apiKeyEnv)
	if !ok {
		t.Fatal(fmt.Sprintf("%s env var not provided", apiKeyEnv))
	}
	return apiKey
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
