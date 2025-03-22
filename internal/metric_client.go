package statsig

import (
	"context"
	"fmt"
	"terraform-provider-statsig/internal/models"
	"terraform-provider-statsig/internal/resource_metric"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type metricClient struct {
	endpoint  string
	transport *Transport
}

func newMetricClient(transport *Transport) *metricClient {
	return &metricClient{
		endpoint:  "metrics",
		transport: transport,
	}
}
func (c *metricClient) read(ctx context.Context, metric *resource_metric.MetricModel) diag.Diagnostics {
	if metric.Id.IsUnknown() {
		metric.Id = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricAPIModel
		res, err := c.transport.Get(c.endpoint, metric.Id.ValueString(), &data)
		models.MetricFromAPIModel(ctx, diags, metric, data)
		return res, err
	})
}
func (c *metricClient) create(ctx context.Context, metric *resource_metric.MetricModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricAPIModel
		res, err := c.transport.Post(c.endpoint, models.MetricToAPIModel(ctx, metric), &data)
		models.MetricFromAPIModel(ctx, diags, metric, data)
		return res, err
	})
}
func (c *metricClient) update(ctx context.Context, metric *resource_metric.MetricModel) diag.Diagnostics {
	if metric.Id.IsUnknown() {
		metric.Id = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricAPIModel
		endpoint := fmt.Sprintf("%s/%s", c.endpoint, metric.Id.ValueString())
		res, err := c.transport.Post(endpoint, models.MetricToAPIModel(ctx, metric), &data)
		models.MetricFromAPIModel(ctx, diags, metric, data)
		return res, err
	})
}
func (c *metricClient) delete(_ context.Context, metric *resource_metric.MetricModel) diag.Diagnostics {
	if metric.Id.IsUnknown() {
		metric.Id = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricAPIModel
		return c.transport.Delete(c.endpoint, metric.Id.ValueString(), &data)
	})
}
