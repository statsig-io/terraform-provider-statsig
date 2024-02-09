package statsig

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dynamicConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"id_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_modifier_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modifier_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"target_apps": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"holdout_ids": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"tags": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		"is_enabled": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: ruleSchema(),
			},
		},
	}
}
