package statsig

import (
	"context"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_keys"

	"github.com/statsig-io/terraform-provider-statsig/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type keysClient struct {
	endpoint  string
	transport *Transport
}

func newKeysClient(transport *Transport) *keysClient {
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
		res, err := c.transport.Get(c.endpoint, key.Key.ValueString(), &data)
		models.KeyFromAPIInputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) create(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.KeysAPIOutputModel
		res, err := c.transport.Post(c.endpoint, models.KeyToAPIInputModel(ctx, key), &data)
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
		res, err := c.transport.Patch(c.endpoint, key.Key.ValueString(), models.KeyToAPIInputModel(ctx, key), &data)
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
		return c.transport.Delete(c.endpoint, key.Key.ValueString(), &data)
	})
}
