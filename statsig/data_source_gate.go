package statsig

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	k := m.(string)
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
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
		gates = append(gates, gate)
	}

	if err := d.Set("gates", gates); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceGates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatesRead,
		Schema: map[string]*schema.Schema{
			"gates": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modifier_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modifier_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"checks_per_hour": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
