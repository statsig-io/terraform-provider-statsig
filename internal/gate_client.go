package statsig

import (
	"context"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_gate"

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
		var data resource_gate.GateAPIModel
		res, err := c.transport.Get(c.endpoint, gate.Id.ValueString(), &data)
		resource_gate.GateFromAPIModel(ctx, diags, gate, data)
		return res, err
	})
}

func (c *gateClient) create(ctx context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_gate.GateAPIModel
		res, err := c.transport.Post(c.endpoint, resource_gate.GateToAPIModel(ctx, gate), &data)
		resource_gate.GateFromAPIModel(ctx, diags, gate, data)
		return res, err
	})
}

func (c *gateClient) update(ctx context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	if gate.Id.IsUnknown() {
		gate.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_gate.GateAPIModel
		res, err := c.transport.Patch(c.endpoint, gate.Id.ValueString(), resource_gate.GateToAPIModel(ctx, gate), &data)
		resource_gate.GateFromAPIModel(ctx, diags, gate, data)
		return res, err
	})
}

func (c *gateClient) delete(_ context.Context, gate *resource_gate.GateModel) diag.Diagnostics {
	if gate.Id.IsUnknown() {
		gate.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_gate.GateAPIModel
		return c.transport.Delete(c.endpoint, gate.Id.ValueString(), &data)
	})
}
