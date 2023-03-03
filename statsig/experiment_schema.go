package statsig

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func experimentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modifier_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modifier_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modifier_time": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"id_type": {
			Type:     schema.TypeString,
			ForceNew: true,
			Optional: true,
			Default:  "userID",
		},
		"layer_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"hypothesis": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"primary_metrics_json": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "[]",
		},
		"primary_metric_tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"secondary_metrics_json": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "[]",
		},
		"secondary_metric_tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"groups": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: groupsSchema(),
			},
		},
		"allocation": {
			Type:             schema.TypeFloat,
			Required:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.FloatBetween(0, 100)),
		},
		"targeting_gate_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"default_confidence_interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "95",
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"80", "90", "95", "98", "99"}, true)),
		},
		"bonferroni_correction": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"duration": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  14,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"status": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "setup",
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"setup", "active", "decision_made", "abandoned"}, true)),
		},
		"launched_group_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func groupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"size": {
			Type:             schema.TypeFloat,
			Required:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.FloatBetween(0, 100)),
		},
		"parameter_values_json": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
