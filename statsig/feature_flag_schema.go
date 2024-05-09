package statsig

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func featureFlagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"is_enabled": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"id_type": {
			Type:     schema.TypeString,
			Required: true,
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
		"checks_per_hour": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: ruleSchema(),
			},
		},
		"dev_rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: ruleSchema(),
			},
		},
		"staging_rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: ruleSchema(),
			},
		},
	}
}

func ruleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pass_percentage": {
			Type:             schema.TypeInt,
			Required:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
		},
		"conditions": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: conditionSchema(),
			},
		},
	}
}

func conditionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"target_value": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"operator": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"field": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
