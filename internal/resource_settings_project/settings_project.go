package resource_settings_project

import (
	"context"
	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// API data model for SettingsProjectModel
// Mirrors SettingsProjectContractDto from OpenAPI
type SettingsProjectAPIModel struct {
	Name            string `json:"name"`
	Visibility      string `json:"visibility"`
	DefaultUnitType string `json:"default_unit_type,omitempty"`
}

func SettingsProjectToAPIModel(ctx context.Context, settingsProject *SettingsProjectModel) SettingsProjectAPIModel {
	return SettingsProjectAPIModel{
		Name:            utils.StringFromNilableValue(settingsProject.Name),
		Visibility:      utils.StringFromNilableValue(settingsProject.Visibility),
		DefaultUnitType: utils.StringFromNilableValue(settingsProject.DefaultUnitType),
	}
}

func SettingsProjectFromAPIModel(ctx context.Context, diags diag.Diagnostics, settingsProject *SettingsProjectModel, res SettingsProjectAPIModel) {
	settingsProject.Name = utils.StringToStringValue(res.Name)
	settingsProject.Visibility = utils.StringToStringValue(res.Visibility)
	settingsProject.DefaultUnitType = utils.StringToNilableValue(res.DefaultUnitType)
}
