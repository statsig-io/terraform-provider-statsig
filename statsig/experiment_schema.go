package statsig

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func experimentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"id_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"layer_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_modifier_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modifier_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"launched_group_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"hypothesis": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"primary_metrics": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"primary_metric_tags": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"secondary_metrics": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"secondary_metric_tags": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: groupsSchema(),
			},
		},
		"allocation": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"targeting_gate_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"default_confidence_interval": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"bonferroni_correction": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeString,
			Computed: true,
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
			Type:     schema.TypeInt,
			Required: true,
		},
		"parameter_values_json": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
