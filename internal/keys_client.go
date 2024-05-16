package statsig

import (
	"context"
	"terraform-provider-statsig/internal/models"
	"terraform-provider-statsig/internal/resource_keys"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type keysClient struct {
	endpoint  string
	transport *transport
}

func newKeysCient(transport *transport) *keysClient {
	return &keysClient{
		endpoint:  "keys",
		transport: transport,
	}
}

func (c *keysClient) read(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	if key.Key.IsUnknown() {
		key.Key = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.KeysAPIOutputModel
		res, err := c.transport.get(c.endpoint, key.Key.ValueString(), &data)
		models.KeyFromAPIInputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) create(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.KeysAPIOutputModel
		res, err := c.transport.post(c.endpoint, models.KeyToAPIInputModel(ctx, key), &data)
		models.KeyFromAPIInputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) update(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	if key.Key.IsUnknown() {
		key.Key = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.KeysAPIOutputModel
		res, err := c.transport.patch(c.endpoint, key.Key.ValueString(), models.KeyToAPIInputModel(ctx, key), &data)
		models.KeyFromAPIInputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) delete(_ context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	if key.Key.IsUnknown() {
		key.Key = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.KeysAPIOutputModel
		return c.transport.delete(c.endpoint, key.Key.ValueString(), &data)
	})
}
