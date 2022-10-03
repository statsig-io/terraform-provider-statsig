package statsig

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	key := d.Get("console_api_key").(string)

	if !strings.HasPrefix(key, "console-") {
		return nil, diag.Errorf("Must provide a valid CONSOLE api key via the environment variable STATSIG_CONSOLE_KEY.")
	}

	return key, nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"console_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"STATSIG_CONSOLE_KEY", "statsig_console_key"}, nil),
				Description: "A Statsig Console API Key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"statsig_gate": resourceGate(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}
