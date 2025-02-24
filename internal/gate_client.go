package statsig

import (
	"context"
	"terraform-provider-statsig/internal/models"
	"terraform-provider-statsig/internal/resource_gate"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type gateClient struct {
	endpoint  string
	transport *Transport
}

func newGateClient(transport *Transport) *gateClient {
	return &gateClient{
		endpoint:  "gates",
		transport: transport,
	}
}

func (c *gateClient) read(ctx context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	if gate.Id.IsUnknown() {
		gate.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.GateAPIModel
		res, err := c.transport.Get(c.endpoint, gate.Id.ValueString(), &data)
		models.GateFromAPIModel(ctx, diags, gate, data)
		return res, err
	})
}

func (c *gateClient) create(ctx context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.GateAPIModel
		res, err := c.transport.Post(c.endpoint, models.GateToAPIModel(ctx, gate), &data)
		models.GateFromAPIModel(ctx, diags, gate, data)
		return res, err
	})
}

func (c *gateClient) update(ctx context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	if gate.Id.IsUnknown() {
		gate.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.GateAPIModel
		res, err := c.transport.Patch(c.endpoint, gate.Id.ValueString(), models.GateToAPIModel(ctx, gate), &data)
		models.GateFromAPIModel(ctx, diags, gate, data)
		return res, err
	})
}

func (c *gateClient) delete(_ context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	if gate.Id.IsUnknown() {
		gate.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.GateAPIModel
		return c.transport.Delete(c.endpoint, gate.Id.ValueString(), &data)
	})
}
