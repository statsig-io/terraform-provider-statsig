package statsig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type APIResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Errors     interface{}
}

func resourceGate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreateGate,
		ReadContext:   resourceReadGate,
		UpdateContext: resourceUpdateGate,
		DeleteContext: resourceDeleteGate,
		Schema:        gateSchema(),
	}
}

func resourceCreateGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := dataFromResource(d)

	if err != nil {
		return diag.FromErr(err)
	}
	res, errDiag := makeAPICall(m.(string), "/gates", "POST", data)

	return handleGatesResponse(errDiag, res, d)
}

func resourceReadGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("/gates/%s", d.Get("id"))
	res, errDiag := makeAPICall(m.(string), e, "GET", nil)
	return handleGatesResponse(errDiag, res, d)
}

func resourceUpdateGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := dataFromResource(d)
	if err != nil {
		return diag.FromErr(err)
	}

	e := fmt.Sprintf("/gates/%s", d.Get("id"))
	res, errDiag := makeAPICall(m.(string), e, "POST", data)

	return handleGatesResponse(errDiag, res, d)
}

func resourceDeleteGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("/gates/%s", d.Get("id"))
	res, errDiag := makeAPICall(m.(string), e, "DELETE", nil)
	return handleResponse(errDiag, res)
}

func makeAPICall(k string, e string, m string, b []byte) (*APIResponse, diag.Diagnostics) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://latest.api.statsig.com/console/v1%s", e)

	req, err := http.NewRequest(m, url, bytes.NewBuffer(b))

	if err != nil {
		return nil, diag.FromErr(err)
	}

	req.Header.Set("statsig-api-key", k)
	if m == "POST" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	defer r.Body.Close()

	response := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	if response["message"] == nil {
		return nil, diag.FromErr(errors.New("gates response is invalid"))
	}

	return &APIResponse{
		StatusCode: r.StatusCode,
		Message:    response["message"].(string),
		Data:       response["data"],
		Errors:     response["errors"],
	}, nil
}

func handleGatesResponse(e diag.Diagnostics, r *APIResponse, d *schema.ResourceData) diag.Diagnostics {
	e = handleResponse(e, r)
	if e != nil {
		return e
	}

	if reflect.TypeOf(r.Data).Kind() != reflect.Map {
		return diag.Errorf("invalid type returned from /gates")
	}

	gateData := r.Data.(map[string]interface{})
	populateResourceFromResponse(d, gateData)

	return nil
}

func handleResponse(e diag.Diagnostics, r *APIResponse) diag.Diagnostics {
	if e != nil {
		return e
	}
	if r.StatusCode != 201 && r.StatusCode != 200 {
		return diag.Errorf("Status %v, Message: %s, Errors: %v", r.StatusCode, r.Message, r.Errors)
	}

	return nil
}

func dataFromResource(d *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":        d.Get("name"),
		"description": d.Get("description"),
		"isEnabled":   d.Get("is_enabled"),
		"idType":      d.Get("id_type"),
	}

	body["rules"] = formatRules(d.Get("rules"), true)
	body["devRules"] = formatRules(d.Get("dev_rules"), true)
	body["stagingRules"] = formatRules(d.Get("staging_rules"), true)

	return json.Marshal(body)
}

func formatRules(in interface{}, forApi bool) []map[string]interface{} {
	if in == nil || reflect.TypeOf(in).Kind() != reflect.Slice {
		return []map[string]interface{}{}
	}

	rules := in.([]interface{})
	result := make([]map[string]interface{}, 0, len(rules))
	for _, v := range rules {
		val := v.(map[string]interface{})
		if forApi {
			result = append(result, map[string]interface{}{
				"name":           val["name"],
				"passPercentage": val["pass_percentage"],
				"conditions":     formatConditions(val["conditions"], forApi),
			})
		} else {
			result = append(result, map[string]interface{}{
				"name":            val["name"],
				"pass_percentage": val["passPercentage"],
				"conditions":      formatConditions(val["conditions"], forApi),
			})
		}
	}

	return result
}

func formatConditions(in interface{}, forApi bool) []map[string]interface{} {
	if in == nil || reflect.TypeOf(in).Kind() != reflect.Slice {
		return []map[string]interface{}{}
	}

	conditions := in.([]interface{})
	result := make([]map[string]interface{}, 0, len(conditions))
	for _, v := range conditions {
		val := v.(map[string]interface{})
		if forApi {
			result = append(result, map[string]interface{}{
				"type":        val["type"],
				"targetValue": val["target_value"],
				"operator":    val["operator"],
				"field":       val["field"],
			})
		} else {
			result = append(result, map[string]interface{}{
				"type":         val["type"],
				"target_value": val["targetValue"],
				"operator":     val["operator"],
				"field":        val["field"],
			})
		}
	}

	return result
}

func populateResourceFromResponse(d *schema.ResourceData, r map[string]interface{}) {
	d.Set("description", r["description"])
	d.Set("is_enabled", r["isEnabled"])
	d.Set("last_modifier_name", r["lastModifierName"])
	d.Set("last_modifier_id", r["lastModifierID"])
	d.Set("checks_per_hour", r["checksPerHour"])
	d.Set("rules", formatRules(r["rules"], false))
	d.Set("dev_rules", formatRules(r["devRules"], false))
	d.Set("staging_rules", formatRules(r["stagingRules"], false))
	d.SetId(r["id"].(string))
}
