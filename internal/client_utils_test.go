package statsig

import (
	"context"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_gate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The Console API answers 4xx with {message, status} and never sends an
// "errors" key, so a status check is the only thing that can catch the failure.
// Without it the empty response decodes into a zero-value model and every
// attribute is written back to state as null.
func TestRunWithDiagnosticsSurfacesClientErrors(t *testing.T) {
	statuses := []int{http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound}

	for _, status := range statuses {
		t.Run(http.StatusText(status), func(t *testing.T) {
			transport, _ := newTestTransport(t, func(w http.ResponseWriter, _ *http.Request) {
				writeConsoleResponse(w, status, map[string]interface{}{
					"message": http.StatusText(status),
					"status":  status,
				})
			})

			var data map[string]interface{}
			diags := runWithDiagnostics(func(diag.Diagnostics) (*APIResponse, error) {
				return transport.Get("gates", "a_gate", &data)
			})

			require.True(t, diags.HasError(), "a %d must produce an error diagnostic", status)
			assert.Contains(t, diags.Errors()[0].Summary(), http.StatusText(status))
		})
	}
}

// A refresh of state whose gate id was nulled read the collection, and the
// array the Console API sends in "data" would not decode into the single-gate
// model. The reported failure was a JSON error that named neither the id nor the
// resource, so the id is what the read has to complain about.
func TestGateClientReadRefusesAnEmptyId(t *testing.T) {
	transport, seen := newTestTransport(t, func(w http.ResponseWriter, _ *http.Request) {
		writeConsoleResponse(w, http.StatusOK, map[string]interface{}{
			"message": "ok",
			"data":    []interface{}{map[string]interface{}{"id": "another_gate"}},
		})
	})

	gate := &resource_gate.GateModel{Id: types.StringNull()}
	diags := newGateClient(transport).read(context.Background(), gate)

	require.True(t, diags.HasError())
	detail := diags.Errors()[0].Detail()
	assert.Contains(t, detail, "cannot GET gates: resource id is empty")
	assert.NotContains(t, detail, "json: cannot unmarshal array into Go struct field Response.Data of type resource_gate.GateAPIModel")
	assert.Empty(t, *seen, "the collection must not be read in place of the resource")
}

func TestRunWithDiagnosticsAcceptsSuccess(t *testing.T) {
	transport, _ := newTestTransport(t, func(w http.ResponseWriter, _ *http.Request) {
		writeConsoleResponse(w, http.StatusOK, map[string]interface{}{
			"message": "ok",
			"data":    map[string]interface{}{"id": "a_gate"},
		})
	})

	var data map[string]interface{}
	diags := runWithDiagnostics(func(diag.Diagnostics) (*APIResponse, error) {
		return transport.Get("gates", "a_gate", &data)
	})

	assert.False(t, diags.HasError())
	assert.Equal(t, map[string]interface{}{"id": "a_gate"}, data)
}
