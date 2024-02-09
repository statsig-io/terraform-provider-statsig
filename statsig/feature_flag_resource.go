package statsig

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFeatureFlag(flagType string) *schema.Resource {
	return statsigResource{
		endpoint:       fmt.Sprintf("/%s", flagType),
		schema:         featureFlagSchema(),
		toJsonData:     dataFromFeatureFlagResource,
		fromJsonObject: populateFeatureFlagResourceFromResponse,
	}.asTerraformResource()
}

func dataFromFeatureFlagResource(ctx context.Context, rd *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":        rd.Get("name"),
		"description": rd.Get("description"),
		"isEnabled":   rd.Get("is_enabled"),
		"idType":      rd.Get("id_type"),
	}

	body["rules"] = formatRules(ctx, rd.Get("rules"), true)
	body["devRules"] = formatRules(ctx, rd.Get("dev_rules"), true)
	body["stagingRules"] = formatRules(ctx, rd.Get("staging_rules"), true)

	return json.Marshal(body)
}

func formatRules(ctx context.Context, in interface{}, forApi bool) []map[string]interface{} {
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
				"conditions":     formatConditions(ctx, val["conditions"], forApi),
			})
		} else {
			result = append(result, map[string]interface{}{
				"name":            val["name"],
				"pass_percentage": val["passPercentage"],
				"conditions":      formatConditions(ctx, val["conditions"], forApi),
			})
		}
	}

	return result
}

func formatConditions(ctx context.Context, in interface{}, forApi bool) []map[string]interface{} {
	if in == nil || reflect.TypeOf(in).Kind() != reflect.Slice {
		return []map[string]interface{}{}
	}

	conditions := in.([]interface{})
	result := make([]map[string]interface{}, 0, len(conditions))

	tflog.Debug(ctx, fmt.Sprintf("Processing Condtions. For API %t. Conditions %s", forApi, arrayToJsonString(conditions)))

	for index, v := range conditions {
		val := v.(map[string]interface{})
		if forApi {
			result = append(result, map[string]interface{}{
				"type":        val["type"],
				"targetValue": val["target_value"],
				"operator":    val["operator"],
				"field":       val["field"],
			})
		} else {
			condition := map[string]interface{}{
				"index": index,
				"type":  formatConditionTypeField(val["type"]),
			}

			maybeAddConditionValue(val["operator"], "operator", &condition)
			maybeAddConditionValue(val["field"], "field", &condition)

			maybeAddTargetValue(val["targetValue"], "target_value", &condition)

			result = append(result, condition)
		}
	}

	return result
}

func formatConditionTypeField(v interface{}) interface{} {
	if v == nil {
		return ""
	}
	return v
}

func maybeAddConditionValue(v interface{}, key string, condition *map[string]interface{}) {
	if v == nil {
		return
	}

	(*condition)[key] = v
}

func maybeAddTargetValue(v interface{}, key string, condition *map[string]interface{}) {
	if v == nil {
		return
	}

	r := make([]string, 0)
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

	(*condition)[key] = r
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

func populateFeatureFlagResourceFromResponse(ctx context.Context, rd *schema.ResourceData, r map[string]interface{}) {
	rd.Set("description", r["description"])
	rd.Set("is_enabled", r["isEnabled"])
	rd.Set("last_modifier_name", r["lastModifierName"])
	rd.Set("last_modifier_id", r["lastModifierID"])
	rd.Set("checks_per_hour", r["checksPerHour"])
	rd.Set("rules", formatRules(ctx, r["rules"], false))
	rd.Set("dev_rules", formatRules(ctx, r["devRules"], false))
	rd.Set("staging_rules", formatRules(ctx, r["stagingRules"], false))

	rd.SetId(r["id"].(string))
}
