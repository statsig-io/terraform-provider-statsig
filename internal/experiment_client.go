package statsig

import (
	"context"
	"terraform-provider-statsig/internal/models"
	"terraform-provider-statsig/internal/resource_experiment"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type experimentClient struct {
	endpoint  string
	transport *Transport
}

func newExperimentClient(transport *Transport) *experimentClient {
	return &experimentClient{
		endpoint:  "experiments",
		transport: transport,
	}
}

func (c *experimentClient) read(ctx context.Context, experiment *resource_experiment.ExperimentModel) diag.Diagnostics {
	if experiment.Id.IsUnknown() {
		experiment.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.ExperimentAPIModel
		res, err := c.transport.Get(c.endpoint, experiment.Id.ValueString(), &data)
		models.ExperimentFromAPIModel(ctx, diags, experiment, data)
		return res, err
	})
}

func (c *experimentClient) create(ctx context.Context, experiment *resource_experiment.ExperimentModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.ExperimentAPIModel
		res, err := c.transport.Post(c.endpoint, models.ExperimentToAPIModel(ctx, experiment), &data)
		models.ExperimentFromAPIModel(ctx, diags, experiment, data)
		return res, err
	})
}

func (c *experimentClient) update(ctx context.Context, experiment *resource_experiment.ExperimentModel) diag.Diagnostics {
	if experiment.Id.IsUnknown() {
		experiment.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.ExperimentAPIModel
		res, err := c.transport.Patch(c.endpoint, experiment.Id.ValueString(), models.ExperimentToAPIModel(ctx, experiment), &data)
		models.ExperimentFromAPIModel(ctx, diags, experiment, data)
		return res, err
	})
}

func (c *experimentClient) delete(_ context.Context, experiment *resource_experiment.ExperimentModel) diag.Diagnostics {
	if experiment.Id.IsUnknown() {
		experiment.Id = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data models.ExperimentAPIModel
		return c.transport.Delete(c.endpoint, experiment.Id.ValueString(), &data)
	})
}
