package statsig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
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
	return makeAPICallAndHandleResponse(m.(string), "/gates", "POST", data, d)
}

func resourceReadGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("/gates/%s", d.Get("id"))
	return makeAPICallAndHandleResponse(m.(string), e, "GET", nil, d)
}

func resourceUpdateGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := dataFromResource(d)
	if err != nil {
		return diag.FromErr(err)
	}
	e := fmt.Sprintf("/gates/%s", d.Get("id"))
	return makeAPICallAndHandleResponse(m.(string), e, "POST", data, d)
}

func resourceDeleteGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("/gates/%s", d.Get("id"))
	diag := makeAPICallAndHandleResponse(m.(string), e, "DELETE", nil, nil)
	if diag == nil {
		d.SetId("")
	}
	return diag
}

func makeAPICallAndHandleResponse(k string, e string, m string, b []byte, d *schema.ResourceData) diag.Diagnostics {
	res, err := makeAPICall(k, e, m, b)
	if err != nil {
		return diag.FromErr(err)
	}
	return handleResponse(res, d)
}

func makeAPICall(k string, e string, m string, b []byte) (*APIResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.statsig.com/console/v1%s", e)

	req, err := http.NewRequest(m, url, bytes.NewBuffer(b))

	if err != nil {
		return nil, err
	}

	req.Header.Set("statsig-api-key", k)
	if m == "POST" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	response := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	if response["message"] == nil {
		return nil, errors.New("gates response is invalid")
	}

	return &APIResponse{
		StatusCode: r.StatusCode,
		Message:    response["message"].(string),
		Data:       response["data"],
		Errors:     response["errors"],
	}, nil
}

func handleResponse(r *APIResponse, d *schema.ResourceData) diag.Diagnostics {
	if r.StatusCode != 201 && r.StatusCode != 200 {
		return diag.Errorf("Status %v, Message: %s, Errors: %v", r.StatusCode, r.Message, r.Errors)
	}

	if d == nil {
		return nil
	}

	if reflect.TypeOf(r.Data).Kind() != reflect.Map {
		return diag.Errorf("invalid type returned from /gates")
	}

	gateData := r.Data.(map[string]interface{})
	populateResourceFromResponse(d, gateData)

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
				"type":         formatConditionValue(val["type"]),
				"target_value": formatTargetValue(val["targetValue"]),
				"operator":     formatConditionValue(val["operator"]),
				"field":        formatConditionValue(val["field"]),
			})
		}
	}

	return result
}

func formatConditionValue(v interface{}) interface{} {
	if v == nil {
		return ""
	}
	return v
}

func formatTargetValue(v interface{}) []string {
	r := make([]string, 0)
	if v == nil {
		return r
	}
	t := reflect.TypeOf(v).Kind()
	switch t {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for _, i := range v.([]interface{}) {
			r = append(r, toString(i))
		}
	default:
		r = append(r, toString(v))
	}

	return r
}

func toString(i interface{}) string {
	asInt, ok := i.(int64)
	if ok {
		return fmt.Sprintf("%d", asInt)
	}

	asFloat, ok := i.(float64)
	if ok {
		asInt := int(asFloat)
		if asFloat == float64(asInt) {
			return fmt.Sprintf("%d", asInt)
		}
		return strconv.FormatFloat(asFloat, 'G', -1, 64)
	}

	asBool, ok := i.(bool)
	if ok {
		if asBool {
			return "true"
		}
		return "false"
	}

	asString, ok := i.(string)
	if ok {
		return asString
	}

	return ""
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

	print(d.Get("staging_rules"))

	d.SetId(r["id"].(string))
}
