package statsig

import (
	"context"
	"terraform-provider-statsig/internal/resource_environments"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type environmentsClient struct {
	endpoint  string
	transport *Transport
}

func newEnvironmentsClient(transport *Transport) *environmentsClient {
	return &environmentsClient{
		endpoint:  "environments",
		transport: transport,
	}
}
func (c *environmentsClient) read(ctx context.Context, environments *resource_environments.EnvironmentsModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_environments.EnvironmentsAPIModel
		res, err := c.transport.Get(c.endpoint, "", &data)
		resource_environments.EnvironmentsFromAPIModel(ctx, diags, environments, data)
		return res, err
	})
}
func (c *environmentsClient) update(ctx context.Context, environments *resource_environments.EnvironmentsModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_environments.EnvironmentsAPIModel
		res, err := c.transport.Post(c.endpoint, resource_environments.EnvironmentsToAPIModel(ctx, environments), &data)
		resource_environments.EnvironmentsFromAPIModel(ctx, diags, environments, data)
		return res, err
	})
}
