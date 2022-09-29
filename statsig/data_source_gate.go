package statsig

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatesRead,
		Schema: map[string]*schema.Schema{
			"gates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: gateSchema(),
				},
			},
		},
	}
}

// region ReadContext

func dataSourceGatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	k := m.(string)
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", "https://api.statsig.com/console/v1/gates", nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("statsig-api-key", k)

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	response := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	if reflect.TypeOf(response["data"]).Kind() != reflect.Slice {
		return diag.Errorf("invalid type returned from /gates")
	}

	gateData := response["data"].([]interface{})
	gates := make([]map[string]interface{}, 0, len(gateData))

	for _, v := range gateData {
		val := v.(map[string]interface{})
		gate := make(map[string]interface{})
		gate["id"] = val["id"]
		gate["is_enabled"] = val["isEnabled"]
		gate["description"] = val["description"]
		gate["last_modifier_name"] = val["lastModifierName"]
		gate["last_modifier_id"] = val["lastModifierID"]
		gate["checks_per_hour"] = val["checksPerHour"]

		//ruleData := val["rules"].([]interface{})
		//rules := rulesRead(ruleData)
		//
		//gate["rules"] = rules

		gates = append(gates, gate)
	}

	if err := d.Set("gates", gates); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func rulesRead(d []interface{}) []map[string]interface{} {
	rules := make([]map[string]interface{}, 0, len(d))

	for _, v := range d {
		val := v.(map[string]interface{})
		rule := make(map[string]interface{})
		rule["name"] = val["name"]
		rule["pass_percentage"] = val["passPercentage"]
		rule["conditions"] = conditionsRead(val["conditions"].([]interface{}))
		rules = append(rules, rule)
	}

	return rules
}

func conditionsRead(i []interface{}) []map[string]interface{} {
	conditions := make([]map[string]interface{}, 0, len(i))

	for _, v := range i {
		val := v.(map[string]interface{})
		cond := make(map[string]interface{})
		cond["type"] = val["type"]

		if val["targetValue"] != nil {
			tvData := val["targetValue"]
			targetValues := make([]string, 0)
			switch reflect.TypeOf(tvData).Kind() {
			case reflect.Slice:
				fallthrough
			case reflect.Array:
				for _, t := range tvData.([]interface{}) {
					targetValues = append(targetValues, fmt.Sprintf("%v", t))
				}
				break
			default:
				targetValues = append(targetValues, fmt.Sprintf("%v", tvData))
			}

			cond["target_value"] = targetValues
		}

		cond["operator"] = val["operator"]
		cond["field"] = val["field"]

		conditions = append(conditions, cond)
	}

	return conditions
}

//endregion
