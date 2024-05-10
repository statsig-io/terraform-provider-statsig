package statsig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type APIResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Errors     interface{}
}

func makeAPICallAndPopulateResource(ctx context.Context, k string, e string, m string, b []byte, r *schema.ResourceData, f func(ctx context.Context, r *schema.ResourceData, d map[string]interface{})) diag.Diagnostics {
	res, err := makeAPICall(ctx, k, e, m, b)
	if err != nil {
		return diag.FromErr(err)
	}

	if res.StatusCode != 201 && res.StatusCode != 200 {
		return diag.Errorf("Status %v, Message: %s, Errors: %v", res.StatusCode, res.Message, res.Errors)
	}

	if r == nil {
		return nil
	}

	if reflect.TypeOf(res.Data).Kind() != reflect.Map {
		return diag.Errorf("invalid type returned from API")
	}

	data := res.Data.(map[string]interface{})
	f(ctx, r, data)

	return nil
}

func makeAPICall(ctx context.Context, k string, e string, m string, b []byte) (*APIResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("https://api.statsig.com/console/v1%s", e)
	tflog.Debug(ctx, fmt.Sprintf("Making Request to %s", url))

	req, err := http.NewRequest(m, url, bytes.NewBuffer(b))

	if err != nil {
		return nil, err
	}

	req.Header.Set("statsig-api-key", k)
	if m == "POST" || m == "PATCH" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	req.Header.Set("statsig-sdk-type", "terraform-provider")
	req.Header.Set("statsig-sdk-version", getVersion())

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	response := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&response)

	tflog.Debug(ctx, fmt.Sprintf("Received Response %s", mapToJsonString(response)))

	if err != nil {
		return nil, err
	}

	if response["message"] == nil {
		return nil, errors.New("invalid response")
	}

	return &APIResponse{
		StatusCode: r.StatusCode,
		Message:    response["message"].(string),
		Data:       response["data"],
		Errors:     response["errors"],
	}, nil
}

func getVersion() string {
	versionBytes, err := os.ReadFile("version")
	version := "unknown"

	if err == nil && versionBytes != nil {
		version = string(versionBytes)
	}

	return version
}
