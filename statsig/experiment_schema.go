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
	}
}
