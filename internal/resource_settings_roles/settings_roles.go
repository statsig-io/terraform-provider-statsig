package resource_settings_roles

import (
	"context"
	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// API data model for SettingsRolesModel
type SettingsRolesAPIModel struct {
	DefaultProjectRole	string	`json:"default_project_role"`
}

func SettingsRolesToAPIModel(ctx context.Context, settingsRoles *SettingsRolesModel) SettingsRolesAPIModel {	
	return SettingsRolesAPIModel{
		DefaultProjectRole: utils.StringFromNilableValue(settingsRoles.DefaultProjectRole),
	}
}

func SettingsRolesFromAPIModel(ctx context.Context, diags diag.Diagnostics, settingsRoles *SettingsRolesModel, res SettingsRolesAPIModel) {
	settingsRoles.DefaultProjectRole = utils.StringToNilableValue(res.DefaultProjectRole)
}
