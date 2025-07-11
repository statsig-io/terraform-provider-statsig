package statsig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_tag"
)

type tagClient struct {
	endpoint  string
	transport *Transport
}

func newTagClient(transport *Transport) *tagClient {
	return &tagClient{
		endpoint:  "tags",
		transport: transport,
	}
}
func (c *tagClient) read(ctx context.Context, tag *resource_tag.TagModel) diag.Diagnostics {
	if tag.Name.IsUnknown() {
		tag.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_tag.TagAPIModel
		res, err := c.transport.Get(c.endpoint, tag.Name.ValueString(), &data)
		resource_tag.TagFromAPIModel(ctx, diags, tag, data)
		return res, err
	})
}
func (c *tagClient) create(ctx context.Context, tag *resource_tag.TagModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_tag.TagAPIModel
		res, err := c.transport.Post(c.endpoint, resource_tag.TagToAPIModel(ctx, tag), &data)
		resource_tag.TagFromAPIModel(ctx, diags, tag, data)
		return res, err
	})
}
func (c *tagClient) update(ctx context.Context, tag *resource_tag.TagModel) diag.Diagnostics {
	if tag.Name.IsUnknown() {
		tag.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_tag.TagAPIModel
		res, err := c.transport.Patch(c.endpoint, tag.Name.ValueString(), resource_tag.TagToAPIModel(ctx, tag), &data)
		resource_tag.TagFromAPIModel(ctx, diags, tag, data)
		return res, err
	})
}
func (c *tagClient) delete(_ context.Context, tag *resource_tag.TagModel) diag.Diagnostics {
	if tag.Name.IsUnknown() {
		tag.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_tag.TagAPIModel
		return c.transport.Delete(c.endpoint, tag.Name.ValueString(), &data)
	})
}
