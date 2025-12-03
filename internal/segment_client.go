package statsig

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_segment"
)

type segmentClient struct {
	endpoint  string
	transport *Transport
}

func newSegmentClient(transport *Transport) *segmentClient {
	return &segmentClient{
		endpoint:  "segments",
		transport: transport,
	}
}

func (c *segmentClient) read(ctx context.Context, segment *resource_segment.SegmentModel) diag.Diagnostics {
	if segment.Id.IsUnknown() {
		segment.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_segment.SegmentAPIModel
		res, err := c.transport.Get(c.endpoint, segment.Id.ValueString(), &data)
		resource_segment.SegmentFromAPIModel(ctx, diags, segment, data)
		return res, err
	})
}

func (c *segmentClient) create(ctx context.Context, segment *resource_segment.SegmentModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_segment.SegmentAPIModel
		res, err := c.transport.Post(c.endpoint, resource_segment.SegmentToAPIModel(ctx, segment), &data)
		resource_segment.SegmentFromAPIModel(ctx, diags, segment, data)
		return res, err
	})
}

func (c *segmentClient) update(ctx context.Context, segment *resource_segment.SegmentModel) diag.Diagnostics {
	if segment.Id.IsUnknown() {
		segment.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_segment.SegmentAPIModel
		endpoint := fmt.Sprintf("%s/%s/conditional", c.endpoint, segment.Id.ValueString())
		res, err := c.transport.Post(endpoint, resource_segment.SegmentToRulesAPIModel(ctx, segment), &data)
		resource_segment.SegmentFromAPIModel(ctx, diags, segment, data)
		return res, err
	})
}

func (c *segmentClient) delete(_ context.Context, segment *resource_segment.SegmentModel) diag.Diagnostics {
	if segment.Id.IsUnknown() {
		segment.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_segment.SegmentAPIModel
		return c.transport.Delete(c.endpoint, segment.Id.ValueString(), &data)
	})
}
