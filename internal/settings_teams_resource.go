package statsig

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_settings_teams"
)

var _ resource.Resource = (*settingsTeamsResource)(nil)

func NewSettingsTeamsResource() resource.Resource {
	return &settingsTeamsResource{}
}

type settingsTeamsResource struct{
	data   *StatsigResourceData
	client *settingsTeamsClient
}

func (r *settingsTeamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	data, ok := req.ProviderData.(*StatsigResourceData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *httpClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.data = data
	r.client = newSettingsTeamsClient(data.transport)
}

func (r *settingsTeamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_teams"
}

func (r *settingsTeamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_settings_teams.SettingsTeamsResourceSchema(ctx)
}

func (r *settingsTeamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resource_settings_teams.SettingsTeamsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...) 
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.client.update(ctx, &data)...) 
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) 
}

func (r *settingsTeamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resource_settings_teams.SettingsTeamsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...) 
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.client.read(ctx, &data)...) 
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) 
}

func (r *settingsTeamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_settings_teams.SettingsTeamsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...) 
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.client.update(ctx, &data)...) 
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...) 
}

func (r *settingsTeamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data resource_settings_teams.SettingsTeamsModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	// NO-OP
}
