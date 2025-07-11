package statsig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_settings_roles"
)

type settingsRolesClient struct {
	endpoint  string
	transport *Transport
}

func newSettingsRolesClient(transport *Transport) *settingsRolesClient {
	return &settingsRolesClient{
		endpoint:  "settings/roles",
		transport: transport,
	}
}
func (c *settingsRolesClient) read(ctx context.Context, settingsRoles *resource_settings_roles.SettingsRolesModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_roles.SettingsRolesAPIModel
		res, err := c.transport.Get(c.endpoint, "", &data)
		resource_settings_roles.SettingsRolesFromAPIModel(ctx, diags, settingsRoles, data)
		return res, err
	})
}
func (c *settingsRolesClient) update(ctx context.Context, settingsRoles *resource_settings_roles.SettingsRolesModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_roles.SettingsRolesAPIModel
		res, err := c.transport.Post(c.endpoint, resource_settings_roles.SettingsRolesToAPIModel(ctx, settingsRoles), &data)
		resource_settings_roles.SettingsRolesFromAPIModel(ctx, diags, settingsRoles, data)
		return res, err
	})
}
