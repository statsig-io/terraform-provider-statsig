package statsig

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExperiment() *schema.Resource {
	return statsigResource{
		endpoint:       "/experiments",
		schema:         experimentSchema(),
		toJsonData:     dataFromExperimentResource,
		fromJsonObject: populateExperimentResourceFromResponse,
	}.asTerraformResource()
}

func dataFromExperimentResource(ctx context.Context, rd *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":                      rd.Get("name"),
		"description":               rd.Get("description"),
		"idType":                    rd.Get("id_type"),
		"status":                    rd.Get("status"),
		"allocation":                rd.Get("allocation"),
		"hypothesis":                rd.Get("hypothesis"),
		"layerID":                   rd.Get("layer_id"),
		"targetingGateID":           rd.Get("targeting_gate_id"),
		"launchedGroupID":           rd.Get("launched_group_id"),
		"defaultConfidenceInterval": rd.Get("default_confidence_interval"),
		"bonferroniCorrection":      rd.Get("bonferroni_correction"),
		"tags":                      rd.Get("tags"),
		"duration":                  rd.Get("duration"),
		"primaryMetricTags":         rd.Get("primary_metric_tags"),
		"secondaryMetricTags":       rd.Get("secondary_metric_tags"),
		"primaryMetrics":            jsonStringToArray(rd.Get("primary_metrics_json")),
		"secondaryMetrics":          jsonStringToArray(rd.Get("secondary_metrics_json")),
	}

	body["groups"] = formatGroups(rd.Get("groups"), true)

	return json.Marshal(body)
}

func populateExperimentResourceFromResponse(ctx context.Context, rd *schema.ResourceData, r map[string]interface{}) {
	rd.Set("description", r["description"])
	rd.Set("id_type", r["idType"])
	rd.Set("name", r["id"])
	rd.Set("last_modifier_name", r["lastModifierName"])
	rd.Set("last_modifier_id", r["lastModifierID"])
	rd.Set("status", r["status"])
	rd.Set("hypothesis", r["hypothesis"])
	rd.Set("layer_id", r["layerID"])
	rd.Set("allocation", r["allocation"])
	rd.Set("targeting_gate_id", r["targetingGateID"])
	rd.Set("launched_group_id", r["launchedGroupID"])
	rd.Set("bonferroni_correction", r["bonferroniCorrection"])
	rd.Set("default_confidence_interval", r["defaultConfidenceInterval"])
	rd.Set("tags", r["tags"])
	rd.Set("duration", r["duration"])
	rd.Set("primary_metric_tags", r["primaryMetricTags"])
	rd.Set("secondary_metric_tags", r["secondaryMetricTags"])
	rd.Set("primary_metrics_json", arrayToJsonString(r["primaryMetrics"]))
	rd.Set("secondary_metrics_json", arrayToJsonString(r["secondaryMetrics"]))

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
				"id":              val["id"],
				"name":            val["name"],
				"size":            val["size"],
				"parameterValues": jsonStringToMap(val["parameter_values_json"]),
			})
		} else {
			result = append(result, map[string]interface{}{
				"id":                    val["id"],
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

func jsonStringToArray(in interface{}) []interface{} {
	var result []interface{}

	value, ok := in.(string)
	if !ok {
		return result
	}

	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return []interface{}{}
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

func arrayToJsonString(in interface{}) string {
	empty := "[]"

	value, ok := in.([]interface{})
	if !ok {
		return empty
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		return empty
	}

	return string(bytes)
}
