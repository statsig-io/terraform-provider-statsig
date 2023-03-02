package statsig

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"reflect"
)

func resourceExperiment() *schema.Resource {
	return statsigResource{
		endpoint:       "/experiments",
		schema:         experimentSchema(),
		toJsonData:     dataFromExperimentResource,
		fromJsonObject: populateExperimentResourceFromResponse,
	}.asTerraformResource()
}

func dataFromExperimentResource(rd *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":        rd.Get("name"),
		"description": rd.Get("description"),
		"idType":      rd.Get("id_type"),
	}

	body["groups"] = formatGroups(rd.Get("groups"), true)

	return json.Marshal(body)
}

func populateExperimentResourceFromResponse(rd *schema.ResourceData, r map[string]interface{}) {
	rd.Set("description", r["description"])
	rd.Set("last_modifier_name", r["lastModifierName"])
	rd.Set("last_modifier_id", r["lastModifierID"])

	rd.Set("groups", formatGroups(r["groups"], false))

	rd.SetId(r["id"].(string))
}

func formatGroups(in interface{}, forApi bool) []map[string]interface{} {
	if in == nil || reflect.TypeOf(in).Kind() != reflect.Slice {
		return []map[string]interface{}{}
	}

	groups := in.([]interface{})
	result := make([]map[string]interface{}, 0, len(groups))
	for _, v := range groups {
		val := v.(map[string]interface{})
		if forApi {
			result = append(result, map[string]interface{}{
				"name":            val["name"],
				"size":            val["size"],
				"parameterValues": jsonStringToMap(val["parameter_values_json"]),
			})
		} else {
			result = append(result, map[string]interface{}{
				"name":                  val["name"],
				"size":                  val["size"],
				"parameter_values_json": mapToJsonString(val["parameterValues"]),
			})
		}
	}

	return result
}

func jsonStringToMap(in interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	value, ok := in.(string)
	if !ok {
		return result
	}

	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return map[string]interface{}{}
	}

	return result
}

func mapToJsonString(in interface{}) string {
	empty := "{}"

	value, ok := in.(map[string]interface{})
	if !ok {
		return empty
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		return empty
	}

	return string(bytes)
}
