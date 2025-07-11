package resource_settings_teams

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
)

// API data model for SettingsTeamsModel
// Mirrors SettingsTeamsContractDto from OpenAPI
// Only one field: require_teams_on_configs

type SettingsTeamsAPIModel struct {
	RequireTeamsOnConfigs bool `json:"require_teams_on_configs"`
}

func SettingsTeamsToAPIModel(ctx context.Context, settingsTeams *SettingsTeamsModel) SettingsTeamsAPIModel {
	return SettingsTeamsAPIModel{
		RequireTeamsOnConfigs: utils.BoolFromBoolValue(settingsTeams.RequireTeamsOnConfigs),
	}
}

func SettingsTeamsFromAPIModel(ctx context.Context, diags diag.Diagnostics, settingsTeams *SettingsTeamsModel, res SettingsTeamsAPIModel) {
	settingsTeams.RequireTeamsOnConfigs = utils.BoolToBoolValue(res.RequireTeamsOnConfigs)
}
