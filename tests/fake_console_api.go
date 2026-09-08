package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"testing"

	provider "github.com/statsig-io/terraform-provider-statsig/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// fakeConsoleAPI stands in for the Statsig Console API so a test can drive real
// Terraform without credentials. It records every request, which is how the
// update-path tests check the URL the provider built and how the keys tests
// check which fields the body carries.
type fakeConsoleAPI struct {
	mu       sync.Mutex
	requests []recordedRequest
	records  map[string]map[string]map[string]interface{}
}

type recordedRequest struct {
	method string
	path   string
	body   []byte
}

func (r recordedRequest) entry() string {
	return r.method + " " + r.path
}

// createFunc turns a POST body into a stored record and the id it is filed under.
type createFunc func(body map[string]interface{}) (string, map[string]interface{})

// fakeCollection describes one Console API collection the fake serves.
type fakeCollection struct {
	create createFunc
	// requestOnly names fields the endpoint accepts but never echoes back.
	requestOnly []string
}

var fakeCollections = map[string]fakeCollection{
	"gates":    {create: echoCreate},
	"segments": {create: echoCreate},
	"metrics":  {create: echoCreate},
	"tags":     {create: nameKeyedCreate},
	"keys":     {create: keysCreate, requestOnly: []string{"targetAppID", "secondaryTargetAppIDs"}},
}

// fakeGeneratedKey is the key POST /console/v1/keys hands back. The real API
// generates it, so a config can never supply it.
const fakeGeneratedKey = "secret-fake-key-0001"

// echoCreate stores the request body as the record and files it under the id the
// API assigns: the body's own id when set, otherwise its name. The gate and
// segment endpoints work this way, since their request and response are the
// same model and the response always carries id.
func echoCreate(body map[string]interface{}) (string, map[string]interface{}) {
	id, _ := body["id"].(string)
	if id == "" {
		id, _ = body["name"].(string)
	}
	body["id"] = id
	return id, body
}

// nameKeyedCreate files the record under its name. statsig_tag has no id of its
// own: the item URL carries the name.
func nameKeyedCreate(body map[string]interface{}) (string, map[string]interface{}) {
	name, _ := body["name"].(string)
	return name, body
}

// keysCreate reproduces the asymmetry in POST /console/v1/keys. The endpoint
// takes targetAppID and secondaryTargetAppIDs, and answers with the display
// names primaryTargetApp and secondaryTargetApps. No response carries the IDs.
func keysCreate(body map[string]interface{}) (string, map[string]interface{}) {
	record := map[string]interface{}{
		"key":                 fakeGeneratedKey,
		"type":                body["type"],
		"description":         body["description"],
		"environments":        body["environments"],
		"scopes":              body["scopes"],
		"primaryTargetApp":    "",
		"secondaryTargetApps": []interface{}{},
	}
	if id, ok := body["targetAppID"].(string); ok && id != "" {
		record["primaryTargetApp"] = "My Edge App"
	}
	if ids, ok := body["secondaryTargetAppIDs"].([]interface{}); ok && len(ids) > 0 {
		record["secondaryTargetApps"] = []interface{}{"My Other App"}
	}
	return fakeGeneratedKey, record
}

func startFakeConsoleAPI(t *testing.T) *fakeConsoleAPI {
	t.Helper()

	f := &fakeConsoleAPI{records: map[string]map[string]map[string]interface{}{}}

	mux := http.NewServeMux()
	for name, collection := range fakeCollections {
		f.records[name] = map[string]map[string]interface{}{}
		mux.HandleFunc("/console/v1/"+name, f.collectionHandler(name, collection))
		mux.HandleFunc("/console/v1/"+name+"/", f.itemHandler(name, collection))
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f.record(r)
		writeNotFound(w)
	})

	address := localTierAddress(t)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatalf("cannot bind %s, the address LocalTier points at: %v", address, err)
	}
	server := &http.Server{Handler: mux}
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(func() { _ = server.Close() })

	return f
}

// localTierAddress derives the listen address from the provider's own LocalTier
// URL so the two cannot drift apart.
func localTierAddress(t *testing.T) string {
	t.Helper()

	parsed, err := url.Parse(provider.LocalAPI)
	if err != nil {
		t.Fatalf("cannot parse %s: %v", provider.LocalAPI, err)
	}
	return parsed.Host
}

// testAccLocalProviders points the provider at the fake Console API. Unlike
// testAccProviders it needs no Statsig credentials.
func testAccLocalProviders() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"statsig": providerserver.NewProtocol6WithError(
			provider.NewTestProvider("console-fake-key", provider.LocalTier),
		),
	}
}

func (f *fakeConsoleAPI) collectionHandler(name string, collection fakeCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.record(r)

		if r.Method == http.MethodGet {
			writeData(w, f.list(name))
			return
		}

		id, record := collection.create(decodeBody(r))
		f.mu.Lock()
		f.records[name][id] = record
		f.mu.Unlock()
		writeData(w, record)
	}
}

func (f *fakeConsoleAPI) itemHandler(name string, collection fakeCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.record(r)

		path := strings.TrimPrefix(r.URL.Path, "/console/v1/"+name+"/")
		// statsig_segment writes its rules through a sub-resource.
		id := strings.TrimSuffix(path, "/conditional")

		f.mu.Lock()
		defer f.mu.Unlock()
		record, found := f.records[name][id]
		if id == "" || !found {
			writeNotFound(w)
			return
		}

		switch r.Method {
		case http.MethodGet:
			writeData(w, record)
		case http.MethodPost, http.MethodPatch:
			for field, value := range decodeBody(r) {
				if !slices.Contains(collection.requestOnly, field) {
					record[field] = value
				}
			}
			writeData(w, record)
		case http.MethodDelete:
			delete(f.records[name], id)
			writeData(w, nil)
		}
	}
}

// requestLog returns every request the provider has made, in order, as
// "METHOD /path".
func (f *fakeConsoleAPI) requestLog() []string {
	f.mu.Lock()
	defer f.mu.Unlock()

	log := make([]string, 0, len(f.requests))
	for _, request := range f.requests {
		log = append(log, request.entry())
	}
	return log
}

// requestBodiesFor returns the body of every request matching "METHOD /path",
// in order.
func (f *fakeConsoleAPI) requestBodiesFor(entry string) [][]byte {
	f.mu.Lock()
	defer f.mu.Unlock()

	var bodies [][]byte
	for _, request := range f.requests {
		if request.entry() == entry {
			bodies = append(bodies, request.body)
		}
	}
	return bodies
}

func (f *fakeConsoleAPI) record(r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(body))

	f.mu.Lock()
	defer f.mu.Unlock()

	f.requests = append(f.requests, recordedRequest{method: r.Method, path: r.URL.Path, body: body})
}

func (f *fakeConsoleAPI) list(name string) []interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()

	records := []interface{}{}
	for _, record := range f.records[name] {
		records = append(records, record)
	}
	return records
}

func decodeBody(r *http.Request) map[string]interface{} {
	body := map[string]interface{}{}
	_ = json.NewDecoder(r.Body).Decode(&body)
	return body
}

func writeData(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"message": "ok", "data": data})
}

// writeNotFound uses the body the Console API actually sends for 4xx: a message
// and a status, with no "errors" key for the provider to key off.
func writeNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Not Found",
		"status":  http.StatusNotFound,
	})
}
