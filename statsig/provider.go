package statsig

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	key := d.Get("key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if !strings.HasPrefix(key, "console-") {
		return nil, diag.Errorf("Must provide a valid CONSOLE api key via the environment variable STATSIG_CONSOLE_KEY.")
	}

	return key, diags
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"STATSIG_CONSOLE_KEY", "statsig_console_key"}, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"statsig_gates": dataSourceGates(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
