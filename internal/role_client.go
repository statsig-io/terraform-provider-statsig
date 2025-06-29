package statsig

import (
	"context"
	"terraform-provider-statsig/internal/resource_role"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type roleClient struct {
	endpoint  string
	transport *Transport
}

func newRoleClient(transport *Transport) *roleClient {
	return &roleClient{
		endpoint:  "roles",
		transport: transport,
	}
}
func (c *roleClient) read(ctx context.Context, role *resource_role.RoleModel) diag.Diagnostics {
	if role.Name.IsUnknown() {
		role.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_role.RoleAPIModel
		res, err := c.transport.Get(c.endpoint, role.Name.ValueString(), &data)
		resource_role.RoleFromAPIModel(ctx, diags, role, data)
		return res, err
	})
}
func (c *roleClient) create(ctx context.Context, role *resource_role.RoleModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_role.RoleAPIModel
		res, err := c.transport.Post(c.endpoint, resource_role.RoleToAPIModel(ctx, role), &data)
		resource_role.RoleFromAPIModel(ctx, diags, role, data)
		return res, err
	})
}
func (c *roleClient) update(ctx context.Context, role *resource_role.RoleModel) diag.Diagnostics {
	if role.Name.IsUnknown() {
		role.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_role.RoleAPIModel
		res, err := c.transport.Patch(c.endpoint, role.Name.ValueString(), resource_role.RoleToAPIModel(ctx, role), &data)
		resource_role.RoleFromAPIModel(ctx, diags, role, data)
		return res, err
	})
}
func (c *roleClient) delete(_ context.Context, role *resource_role.RoleModel) diag.Diagnostics {
	if role.Name.IsUnknown() {
		role.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_role.RoleAPIModel
		return c.transport.Delete(c.endpoint, role.Name.ValueString(), &data)
	})
}
