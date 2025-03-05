package statsig

import (
	"context"
	"fmt"
	"terraform-provider-statsig/internal/models"
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
		var data models.MetricSourceAPIModel
		res, err := c.transport.Get(c.endpoint, metricSource.Name.ValueString(), &data)
		models.MetricSourceFromAPIModel(ctx, diags, metricSource, data)
		return res, err
	})
}
func (c *metricSourceClient) create(ctx context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricSourceAPIModel
		res, err := c.transport.Post(c.endpoint, models.MetricSourceToAPIModel(ctx, metricSource), &data)
		models.MetricSourceFromAPIModel(ctx, diags, metricSource, data)
		return res, err
	})
}
func (c *metricSourceClient) update(ctx context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	if metricSource.Name.IsUnknown() {
		metricSource.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricSourceAPIModel
		endpoint := fmt.Sprintf("%s/%s", c.endpoint, metricSource.Name.ValueString())
		res, err := c.transport.Post(endpoint, models.MetricSourceToAPIModel(ctx, metricSource), &data)
		models.MetricSourceFromAPIModel(ctx, diags, metricSource, data)
		return res, err
	})
}
func (c *metricSourceClient) delete(_ context.Context, metricSource *resource_metric_source.MetricSourceModel) diag.Diagnostics {
	if metricSource.Name.IsUnknown() {
		metricSource.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.MetricSourceAPIModel
		return c.transport.Delete(c.endpoint, metricSource.Name.ValueString(), &data)
	})
}
