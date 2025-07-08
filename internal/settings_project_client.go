package statsig

import (
	"context"
	"terraform-provider-statsig/internal/resource_settings_project"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type settingsProjectClient struct {
	endpoint  string
	transport *Transport
}

func newSettingsProjectClient(transport *Transport) *settingsProjectClient {
	return &settingsProjectClient{
		endpoint:  "settings/project",
		transport: transport,
	}
}

func (c *settingsProjectClient) read(ctx context.Context, settingsProject *resource_settings_project.SettingsProjectModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_project.SettingsProjectAPIModel
		res, err := c.transport.Get(c.endpoint, "", &data)
		resource_settings_project.SettingsProjectFromAPIModel(ctx, diags, settingsProject, data)
		return res, err
	})
}

func (c *settingsProjectClient) update(ctx context.Context, settingsProject *resource_settings_project.SettingsProjectModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_project.SettingsProjectAPIModel
		res, err := c.transport.Post(c.endpoint, resource_settings_project.SettingsProjectToAPIModel(ctx, settingsProject), &data)
		resource_settings_project.SettingsProjectFromAPIModel(ctx, diags, settingsProject, data)
		return res, err
	})
}
