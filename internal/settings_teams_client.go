package statsig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_settings_teams"
)

type settingsTeamsClient struct {
	endpoint  string
	transport *Transport
}

func newSettingsTeamsClient(transport *Transport) *settingsTeamsClient {
	return &settingsTeamsClient{
		endpoint:  "settings/teams",
		transport: transport,
	}
}

func (c *settingsTeamsClient) read(ctx context.Context, settingsTeams *resource_settings_teams.SettingsTeamsModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_teams.SettingsTeamsAPIModel
		res, err := c.transport.Get(c.endpoint, "", &data)
		resource_settings_teams.SettingsTeamsFromAPIModel(ctx, diags, settingsTeams, data)
		return res, err
	})
}

func (c *settingsTeamsClient) update(ctx context.Context, settingsTeams *resource_settings_teams.SettingsTeamsModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_teams.SettingsTeamsAPIModel
		res, err := c.transport.Post(c.endpoint, resource_settings_teams.SettingsTeamsToAPIModel(ctx, settingsTeams), &data)
		resource_settings_teams.SettingsTeamsFromAPIModel(ctx, diags, settingsTeams, data)
		return res, err
	})
}
