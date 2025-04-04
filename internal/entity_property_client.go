package statsig

import (
	"context"
	"fmt"
	"terraform-provider-statsig/internal/resource_entity_property"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type entityPropertyClient struct {
	endpoint  string
	transport *Transport
}

func newEntityPropertyClient(transport *Transport) *entityPropertyClient {
	return &entityPropertyClient{
		endpoint:  "experiments/entity_property",
		transport: transport,
	}
}
func (c *entityPropertyClient) read(ctx context.Context, entityProperty *resource_entity_property.EntityPropertyModel) diag.Diagnostics {
	if entityProperty.Name.IsUnknown() {
		entityProperty.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_entity_property.EntityPropertyAPIModel
		res, err := c.transport.Get(c.endpoint, entityProperty.Name.ValueString(), &data)
		resource_entity_property.EntityPropertyFromAPIModel(ctx, diags, entityProperty, data)
		return res, err
	})
}
func (c *entityPropertyClient) create(ctx context.Context, entityProperty *resource_entity_property.EntityPropertyModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_entity_property.EntityPropertyAPIModel
		res, err := c.transport.Post("experiments/entity_properties", resource_entity_property.EntityPropertyToAPIModel(ctx, entityProperty), &data)
		resource_entity_property.EntityPropertyFromAPIModel(ctx, diags, entityProperty, data)
		return res, err
	})
}
func (c *entityPropertyClient) update(ctx context.Context, entityProperty *resource_entity_property.EntityPropertyModel) diag.Diagnostics {
	if entityProperty.Name.IsUnknown() {
		entityProperty.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_entity_property.EntityPropertyAPIModel
		endpoint := fmt.Sprintf("%s/%s", c.endpoint, entityProperty.Name.ValueString())
		res, err := c.transport.Post(endpoint, resource_entity_property.EntityPropertyToAPIModel(ctx, entityProperty), &data)
		resource_entity_property.EntityPropertyFromAPIModel(ctx, diags, entityProperty, data)
		return res, err
	})
}
func (c *entityPropertyClient) delete(_ context.Context, entityProperty *resource_entity_property.EntityPropertyModel) diag.Diagnostics {
	if entityProperty.Name.IsUnknown() {
		entityProperty.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_entity_property.EntityPropertyAPIModel
		return c.transport.Delete(c.endpoint, entityProperty.Name.ValueString(), &data)
	})
}
