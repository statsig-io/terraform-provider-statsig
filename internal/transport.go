package statsig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type transport struct {
	ctx      context.Context
	api      string
	apiKey   string
	metadata statsigMetadata
	client   *http.Client
}

type Response struct {
	Message string
	Data    interface{}
	Errors  interface{}
}

type APIResponse struct {
	StatusCode int
	Response
}

func newTransport(_ context.Context, apiKey string) *transport {
	return &transport{
		api:      "https://api.statsig.com/console/v1",
		apiKey:   apiKey,
		metadata: getStatsigMetadata(),
		client:   &http.Client{Timeout: time.Second * 3},
	}
}

func (t *transport) get(endpoint string, id string, resp interface{}) (*APIResponse, error) {
	return t.doRequest("GET", fmt.Sprintf("%s/%s", endpoint, id), nil, resp)
}

func (t *transport) post(endpoint string, body interface{}, resp interface{}) (*APIResponse, error) {
	return t.doRequest("POST", endpoint, body, resp)
}

func (t *transport) patch(endpoint string, id string, body interface{}, resp interface{}) (*APIResponse, error) {
	return t.doRequest("PATCH", fmt.Sprintf("%s/%s", endpoint, id), body, resp)
}

func (t *transport) delete(endpoint string, id string, body interface{}, resp interface{}) (*APIResponse, error) {
	return t.doRequest("DELETE", fmt.Sprintf("%s/%s", endpoint, id), body, resp)
}

func (t *transport) doRequest(method, endpoint string, body interface{}, resp interface{}) (*APIResponse, error) {
	req, err := t.buildRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	r, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("Failed %s request to %s with status code %d.", method, req.URL, r.StatusCode))
	}
	defer r.Body.Close()

	var response Response
	response.Data = resp
	err = json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &APIResponse{
		StatusCode: r.StatusCode,
		Response:   response,
	}, nil
}

func (t *transport) buildRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	var bodyBuf io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyBuf = bytes.NewBuffer(bodyBytes)
	} else {
		if method == "POST" {
			bodyBuf = bytes.NewBufferString("{}")
		}
	}
	url := fmt.Sprintf("%s/%s", t.api, endpoint)
	req, err := http.NewRequest(method, url, bodyBuf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("STATSIG-API-KEY", t.apiKey)
	if method == "POST" || method == "PATCH" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	req.Header.Add("STATSIG-SDK-TYPE", t.metadata.SDKType)
	req.Header.Add("STATSIG-SDK-VERSION", t.metadata.SDKVersion)
	return req, nil
}
