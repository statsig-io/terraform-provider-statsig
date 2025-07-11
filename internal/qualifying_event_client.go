package statsig

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_qualifying_event"
)

type qualifyingEventClient struct {
	endpoint  string
	transport *Transport
}

func newQualifyingEventClient(transport *Transport) *qualifyingEventClient {
	return &qualifyingEventClient{
		endpoint:  "experiments/qualifying_events",
		transport: transport,
	}
}
func (c *qualifyingEventClient) read(ctx context.Context, qualifyingEvent *resource_qualifying_event.QualifyingEventModel) diag.Diagnostics {
	if qualifyingEvent.Name.IsUnknown() {
		qualifyingEvent.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_qualifying_event.QualifyingEventAPIModel
		res, err := c.transport.Get(c.endpoint, qualifyingEvent.Name.ValueString(), &data)
		resource_qualifying_event.QualifyingEventFromAPIModel(ctx, diags, qualifyingEvent, data)
		return res, err
	})
}
func (c *qualifyingEventClient) create(ctx context.Context, qualifyingEvent *resource_qualifying_event.QualifyingEventModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_qualifying_event.QualifyingEventAPIModel
		res, err := c.transport.Post(c.endpoint, resource_qualifying_event.QualifyingEventToAPIModel(ctx, qualifyingEvent), &data)
		resource_qualifying_event.QualifyingEventFromAPIModel(ctx, diags, qualifyingEvent, data)
		return res, err
	})
}
func (c *qualifyingEventClient) update(ctx context.Context, qualifyingEvent *resource_qualifying_event.QualifyingEventModel) diag.Diagnostics {
	if qualifyingEvent.Name.IsUnknown() {
		qualifyingEvent.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_qualifying_event.QualifyingEventAPIModel
		endpoint := fmt.Sprintf("%s/%s", c.endpoint, qualifyingEvent.Name.ValueString())
		res, err := c.transport.Post(endpoint, resource_qualifying_event.QualifyingEventToAPIModel(ctx, qualifyingEvent), &data)
		resource_qualifying_event.QualifyingEventFromAPIModel(ctx, diags, qualifyingEvent, data)
		return res, err
	})
}
func (c *qualifyingEventClient) delete(_ context.Context, qualifyingEvent *resource_qualifying_event.QualifyingEventModel) diag.Diagnostics {
	if qualifyingEvent.Name.IsUnknown() {
		qualifyingEvent.Name = types.StringNull()
	}
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_qualifying_event.QualifyingEventAPIModel
		return c.transport.Delete(c.endpoint, qualifyingEvent.Name.ValueString(), &data)
	})
}
