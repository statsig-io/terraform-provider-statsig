package statsig

import (
	"context"
	"terraform-provider-statsig/internal/resource_unit_id_type"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type unitIdTypeClient struct {
	endpoint  string
	transport *Transport
}

func newUnitIdTypeClient(transport *Transport) *unitIdTypeClient {
	return &unitIdTypeClient{
		endpoint:  "unit_id_types",
		transport: transport,
	}
}

func (c *unitIdTypeClient) read(ctx context.Context, unitIdType *resource_unit_id_type.UnitIdTypeModel) diag.Diagnostics {
	if unitIdType.Name.IsUnknown() {
		unitIdType.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_unit_id_type.UnitIdTypeAPIModel
		res, err := c.transport.Get(c.endpoint, unitIdType.Name.ValueString(), &data)
		resource_unit_id_type.UnitIdTypeFromAPIModel(ctx, diags, unitIdType, data)
		return res, err
	})
}

func (c *unitIdTypeClient) create(ctx context.Context, unitIdType *resource_unit_id_type.UnitIdTypeModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_unit_id_type.UnitIdTypeAPIModel
		res, err := c.transport.Post(c.endpoint, resource_unit_id_type.UnitIdTypeToAPIModel(ctx, unitIdType), &data)
		resource_unit_id_type.UnitIdTypeFromAPIModel(ctx, diags, unitIdType, data)
		return res, err
	})
}

func (c *unitIdTypeClient) update(ctx context.Context, unitIdType *resource_unit_id_type.UnitIdTypeModel) diag.Diagnostics {
	if unitIdType.Name.IsUnknown() {
		unitIdType.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_unit_id_type.UnitIdTypeAPIModel
		res, err := c.transport.Patch(c.endpoint, unitIdType.Name.ValueString(), resource_unit_id_type.UnitIdTypeToAPIModel(ctx, unitIdType), &data)
		resource_unit_id_type.UnitIdTypeFromAPIModel(ctx, diags, unitIdType, data)
		return res, err
	})
}

func (c *unitIdTypeClient) delete(_ context.Context, unitIdType *resource_unit_id_type.UnitIdTypeModel) diag.Diagnostics {
	if unitIdType.Name.IsUnknown() {
		unitIdType.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_unit_id_type.UnitIdTypeAPIModel
		return c.transport.Delete(c.endpoint, unitIdType.Name.ValueString(), &data)
	})
}
