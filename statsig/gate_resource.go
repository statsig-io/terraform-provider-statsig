package statsig

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"reflect"
	"strconv"
)

func resourceGate() *schema.Resource {
	return statsigResource{
		endpoint:       "/gates",
		schema:         gateSchema(),
		toJsonData:     dataFromGateResource,
		fromJsonObject: populateGateResourceFromResponse,
	}.asTerraformResource()
}

func dataFromGateResource(rd *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":        rd.Get("name"),
		"description": rd.Get("description"),
		"isEnabled":   rd.Get("is_enabled"),
		"idType":      rd.Get("id_type"),
	}

	body["rules"] = formatRules(rd.Get("rules"), true)
	body["devRules"] = formatRules(rd.Get("dev_rules"), true)
	body["stagingRules"] = formatRules(rd.Get("staging_rules"), true)

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

func populateGateResourceFromResponse(rd *schema.ResourceData, r map[string]interface{}) {
	rd.Set("description", r["description"])
	rd.Set("is_enabled", r["isEnabled"])
	rd.Set("last_modifier_name", r["lastModifierName"])
	rd.Set("last_modifier_id", r["lastModifierID"])
	rd.Set("checks_per_hour", r["checksPerHour"])
	rd.Set("rules", formatRules(r["rules"], false))
	rd.Set("dev_rules", formatRules(r["devRules"], false))
	rd.Set("staging_rules", formatRules(r["stagingRules"], false))

	rd.SetId(r["id"].(string))
}
