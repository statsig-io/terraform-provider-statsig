package statsig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_dynamic_config"
)

type dynamicConfigClient struct {
	endpoint  string
	transport *Transport
}

func newDynamicConfigClient(transport *Transport) *dynamicConfigClient {
	return &dynamicConfigClient{
		endpoint:  "dynamic_configs",
		transport: transport,
	}
}

func (c *dynamicConfigClient) read(ctx context.Context, dynamicConfig *resource_dynamic_config.DynamicConfigModel) diag.Diagnostics {
	if dynamicConfig.Id.IsUnknown() {
		dynamicConfig.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_dynamic_config.DynamicConfigAPIModel
		res, err := c.transport.Get(c.endpoint, dynamicConfig.Id.ValueString(), &data)
		resource_dynamic_config.DynamicConfigFromAPIModel(ctx, diags, dynamicConfig, data)
		return res, err
	})
}

func (c *dynamicConfigClient) create(ctx context.Context, dynamicConfig *resource_dynamic_config.DynamicConfigModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_dynamic_config.DynamicConfigAPIModel
		res, err := c.transport.Post(c.endpoint, resource_dynamic_config.DynamicConfigToAPIModel(ctx, dynamicConfig), &data)
		resource_dynamic_config.DynamicConfigFromAPIModel(ctx, diags, dynamicConfig, data)
		return res, err
	})
}

func (c *dynamicConfigClient) update(ctx context.Context, dynamicConfig *resource_dynamic_config.DynamicConfigModel) diag.Diagnostics {
	if dynamicConfig.Id.IsUnknown() {
		dynamicConfig.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_dynamic_config.DynamicConfigAPIModel
		res, err := c.transport.Patch(c.endpoint, dynamicConfig.Id.ValueString(), resource_dynamic_config.DynamicConfigToAPIModel(ctx, dynamicConfig), &data)
		resource_dynamic_config.DynamicConfigFromAPIModel(ctx, diags, dynamicConfig, data)
		return res, err
	})
}

func (c *dynamicConfigClient) delete(_ context.Context, dynamicConfig *resource_dynamic_config.DynamicConfigModel) diag.Diagnostics {
	if dynamicConfig.Id.IsUnknown() {
		dynamicConfig.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_dynamic_config.DynamicConfigAPIModel
		return c.transport.Delete(c.endpoint, dynamicConfig.Id.ValueString(), &data)
	})
}
