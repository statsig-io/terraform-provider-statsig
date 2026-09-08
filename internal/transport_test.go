package statsig

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestTransport points a Transport at a local server instead of a tier and
// records the path of every request that reaches it.
func newTestTransport(t *testing.T, handler http.HandlerFunc) (*Transport, *[]string) {
	t.Helper()

	var seen []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.Method+" "+r.URL.Path)
		handler(w, r)
	}))
	t.Cleanup(srv.Close)

	return &Transport{
		api:    srv.URL + "/console/v1",
		apiKey: "console-test-key",
		client: srv.Client(),
	}, &seen
}

func writeConsoleResponse(w http.ResponseWriter, status int, body map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// An empty id used to build the URL sent an update aimed at one gate to the
// whole collection instead. Both verbs must refuse to send anything.
func TestTransportRefusesEmptyId(t *testing.T) {
	transport, seen := newTestTransport(t, func(w http.ResponseWriter, _ *http.Request) {
		writeConsoleResponse(w, http.StatusOK, map[string]interface{}{"message": "ok"})
	})

	var data map[string]interface{}

	_, err := transport.Patch("gates", "", map[string]string{"name": "a_gate"}, &data)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "resource id is empty")

	_, err = transport.Delete("gates", "", &data)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "resource id is empty")

	assert.Empty(t, *seen, "no request should reach the API")
}

func TestTransportSendsItemUrlForNonEmptyId(t *testing.T) {
	transport, seen := newTestTransport(t, func(w http.ResponseWriter, _ *http.Request) {
		writeConsoleResponse(w, http.StatusOK, map[string]interface{}{
			"message": "ok",
			"data":    map[string]interface{}{},
		})
	})

	var data map[string]interface{}

	_, err := transport.Patch("gates", "a_gate", map[string]string{"name": "a_gate"}, &data)
	require.NoError(t, err)

	_, err = transport.Delete("gates", "a_gate", &data)
	require.NoError(t, err)

	assert.Equal(t, []string{
		"PATCH /console/v1/gates/a_gate",
		"DELETE /console/v1/gates/a_gate",
	}, *seen)
}

// Get keeps its collection fallback: environments and the settings_* singletons
// pass an empty id on purpose.
func TestTransportGetFallsBackToCollectionForEmptyId(t *testing.T) {
	transport, seen := newTestTransport(t, func(w http.ResponseWriter, _ *http.Request) {
		writeConsoleResponse(w, http.StatusOK, map[string]interface{}{
			"message": "ok",
			"data":    map[string]interface{}{},
		})
	})

	var data map[string]interface{}
	_, err := transport.Get("environments", "", &data)
	require.NoError(t, err)

	assert.Equal(t, []string{"GET /console/v1/environments"}, *seen)
}
