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

func resourceGate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreateGate,
		ReadContext:   resourceReadGate,
		UpdateContext: resourceUpdateGate,
		DeleteContext: resourceDeleteGate,
		Schema:        gateSchema(),
	}
}

// region CreateContext

func resourceCreateGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	k := m.(string)
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	data, err := dataFromResource(d)

	//return diag.Errorf("%v", string(data))

	if err != nil {
		return diag.FromErr(err)
	}

	req, err := http.NewRequest("POST", "https://latest.api.statsig.com/console/v1/gates", bytes.NewBuffer(data))
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("statsig-api-key", k)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := runRequest(client, req)
	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != 201 {
		return diag.Errorf("%s, %v", res.Message, res.Errors)
	}
	if reflect.TypeOf(res.Data).Kind() != reflect.Map {
		return diag.Errorf("invalid type returned from /gates")
	}

	gateData := res.Data.(map[string]interface{})
	populateResourceFromResponse(d, gateData)

	return diags
}

// endregion

// region ReadContext

func resourceReadGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	k := m.(string)
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	//return diag.Errorf("Read")
	url := fmt.Sprintf("https://api.statsig.com/console/v1/gates/%s", d.Get("id"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("statsig-api-key", k)

	res, err := runRequest(client, req)

	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != 200 {
		return diag.Errorf("%s, %v", res.Message, res.Errors)
	}
	if reflect.TypeOf(res.Data).Kind() != reflect.Map {
		return diag.Errorf("invalid type returned from /gates")
	}

	val := res.Data.(map[string]interface{})
	populateResourceFromResponse(d, val)

	return diags
}

//endregion

// region UpdateContext

func resourceUpdateGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	k := m.(string)
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	data, err := dataFromResource(d)

	//return diag.Errorf("%v", string(data))

	if err != nil {
		return diag.FromErr(err)
	}

	url := fmt.Sprintf("https://api.statsig.com/console/v1/gates/%s", d.Get("id"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("statsig-api-key", k)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := runRequest(client, req)

	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != 200 {
		return diag.Errorf("%s, %v", res.Message, res.Errors)
	}
	if reflect.TypeOf(res.Data).Kind() != reflect.Map {
		return diag.Errorf("invalid type returned from /gates")
	}

	gateData := res.Data.(map[string]interface{})
	populateResourceFromResponse(d, gateData)

	return diags
}

// endregion

// region DeleteContext

func resourceDeleteGate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	k := m.(string)
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	//return diag.Errorf("Read")
	url := fmt.Sprintf("https://api.statsig.com/console/v1/gates/%s", d.Get("id"))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("statsig-api-key", k)

	res, err := runRequest(client, req)

	if err != nil {
		return diag.FromErr(err)
	}
	if res.StatusCode != 200 {
		return diag.Errorf(res.Message)
	}

	return diags
}

// endregion

type APIResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Errors     interface{}
}

func runRequest(client *http.Client, req *http.Request) (*APIResponse, error) {
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
	}, err
}

func dataFromResource(d *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":        d.Get("name"),
		"description": d.Get("description"),
		"isEnabled":   d.Get("is_enabled"),
		"idType":      d.Get("id_type"),
	}

	rules := d.Get("rules")
	if rules != nil {
		body["rules"] = rules
	} else {
		body["rules"] = []interface{}{}
	}

	return json.Marshal(body)
}

func populateResourceFromResponse(d *schema.ResourceData, r map[string]interface{}) {
	d.Set("description", r["description"])
	d.Set("is_enabled", r["isEnabled"])
	d.Set("last_modifier_name", r["lastModifierName"])
	d.Set("last_modifier_id", r["lastModifierID"])
	d.Set("checks_per_hour", r["checksPerHour"])
	d.Set("rules", []interface{}{})

	d.SetId(r["id"].(string))
}
