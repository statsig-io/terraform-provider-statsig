package statsig

import (
	"context"
	"os"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*StatsigProvider)(nil)

func New() provider.Provider {
	return &StatsigProvider{}
}

type StatsigProvider struct{}

type StatsigProviderModel struct {
	apiKey types.String `tfsdk:"console_api_key"`
}

type StatsigResourceData struct {
	transport *Transport
}

func (p *StatsigProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"console_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "A Statsig Console API Key",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^console-.*"), "Provided key is not a valid Console API key"),
				},
			},
		},
	}
}

func (p *StatsigProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Check environment variables
	apiKey := os.Getenv("STATSIG_CONSOLE_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("statsig_console_key")
	}

	var data StatsigProviderModel

	// Read configuration data into model
	// TODO: figure out how to mute this for acceptance tests which must provide this via env variable
	// resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	req.Config.Get(ctx, &data)

	// Check configuration data, which should take precedence over
	// environment variable data, if found.
	if data.apiKey.ValueString() != "" {
		apiKey = data.apiKey.ValueString()
	}

	if apiKey == "" {
		resp.Diagnostics.AddError(
			"Missing Console API Key Configuration",
			"While configuring the provider, the Console API key was not found in "+
				"the STATSIG_CONSOLE_KEY environment variable or provider "+
				"configuration block console_api_key attribute.",
		)
	}

	resp.ResourceData = &StatsigResourceData{
		transport: NewTransport(ctx, apiKey),
	}
}

func (p *StatsigProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "statsig"
}

func (p *StatsigProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *StatsigProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewKeysResource,
		NewGateResource,
		NewExperimentResource,
		NewEntityPropertyResource,
		NewMetricSourceResource,
	}
}
