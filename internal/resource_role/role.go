package resource_role

import (
	"context"
	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// API data model for RoleModel
type RoleAPIModel struct {
	Name        string					`json:"name"`
	Permissions map[string]interface{}	`json:"permissions"`
}

func RoleToAPIModel(ctx context.Context, role *RoleModel) RoleAPIModel {	
	return RoleAPIModel{
		Name:        utils.StringFromNilableValue(role.Name),
		Permissions: utils.MapFromMapValue(ctx, role.Permissions),
	}
}

func RoleFromAPIModel(ctx context.Context, diags diag.Diagnostics, role *RoleModel, res RoleAPIModel) {
	role.Name = utils.StringToNilableValue(res.Name)
	role.Permissions = utils.MapToMapValue(ctx, diags, res.Permissions)
}
