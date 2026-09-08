package statsig

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"
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

			assert.True(t, diags.HasError(), "a %d must produce an error diagnostic", status)
			assert.Contains(t, diags.Errors()[0].Summary(), http.StatusText(status))
		})
	}
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
