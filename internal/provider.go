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

type StatsigProviderVersion string

const (
	Test StatsigProviderVersion = "test"
	Prod StatsigProviderVersion = "1.0.0"
)

type ConsoleAPITier string

const (
	LatestTier  ConsoleAPITier = "latest"
	StagingTier ConsoleAPITier = "staging"
	ProdTier    ConsoleAPITier = ""
)

var _ provider.Provider = (*StatsigProvider)(nil)

func New() provider.Provider {
	return &StatsigProvider{
		Version: Prod,
	}
}

func NewTestProvider(apiKey string, tier ConsoleAPITier) provider.Provider {
	return &StatsigProvider{
		Version: Test,
		APIKey:  apiKey,
		Tier:    tier,
	}
}

type StatsigProvider struct {
	Version StatsigProviderVersion
	APIKey  string
	Tier    ConsoleAPITier
}

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
	if p.Version != Test {
		// Check environment variables
		p.APIKey = os.Getenv("STATSIG_CONSOLE_KEY")
		if p.APIKey == "" {
			p.APIKey = os.Getenv("statsig_console_key")
		}
	}

	var data StatsigProviderModel

	// Read configuration data into model
	// TODO: figure out how to mute this for acceptance tests which must provide this via env variable
	// resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	req.Config.Get(ctx, &data)

	// Check configuration data, which should take precedence over
	// environment variable data, if found.
	if data.apiKey.ValueString() != "" {
		p.APIKey = data.apiKey.ValueString()
	}

	if p.APIKey == "" {
		resp.Diagnostics.AddError(
			"Missing Console API Key Configuration",
			"While configuring the provider, the Console API key was not found in "+
				"the STATSIG_CONSOLE_KEY environment variable or provider "+
				"configuration block console_api_key attribute.",
		)
	}

	resp.ResourceData = &StatsigResourceData{
		transport: NewTransport(ctx, p.APIKey, p.Version, p.Tier),
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
		NewMetricResource,
	}
}
