package statsig

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider = Provider()
var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"statsig": func() (*schema.Provider, error) {
		return testAccProvider, nil
	},
}

func testAccPreCheck(t *testing.T) {
	_, ok := os.LookupEnv("statsig_console_key")
	if !ok {
		t.Fatal("statsig_console_key env var not provided")
	}
}
