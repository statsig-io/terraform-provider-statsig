package statsig

import (
	"context"
	"fmt"
	"terraform-provider-statsig/internal/resource_metric_source"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type metricSourceClient struct {
	endpoint  string
	transport *Transport
}

func newMetricSourceClient(transport *Transport) *metricSourceClient {
	return &metricSourceClient{
		endpoint:  "metrics/metric_source",
		transport: transport,
	}
}
func (c *metricSourceClient) read(ctx context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	if metricSource.Name.IsUnknown() {
		metricSource.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_metric_source.MetricSourceAPIModel
		res, err := c.transport.Get(c.endpoint, metricSource.Name.ValueString(), &data)
		resource_metric_source.MetricSourceFromAPIModel(ctx, diags, metricSource, data)
		return res, err
	})
}
func (c *metricSourceClient) create(ctx context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_metric_source.MetricSourceAPIModel
		res, err := c.transport.Post(c.endpoint, resource_metric_source.MetricSourceToAPIModel(ctx, metricSource), &data)
		resource_metric_source.MetricSourceFromAPIModel(ctx, diags, metricSource, data)
		return res, err
	})
}
func (c *metricSourceClient) update(ctx context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	if metricSource.Name.IsUnknown() {
		metricSource.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_metric_source.MetricSourceAPIModel
		endpoint := fmt.Sprintf("%s/%s", c.endpoint, metricSource.Name.ValueString())
		res, err := c.transport.Post(endpoint, resource_metric_source.MetricSourceToAPIModel(ctx, metricSource), &data)
		resource_metric_source.MetricSourceFromAPIModel(ctx, diags, metricSource, data)
		return res, err
	})
}
func (c *metricSourceClient) delete(_ context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	if metricSource.Name.IsUnknown() {
		metricSource.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_metric_source.MetricSourceAPIModel
		return c.transport.Delete(c.endpoint, metricSource.Name.ValueString(), &data)
	})
}
