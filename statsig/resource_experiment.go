package statsig

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExperiment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreateExperiment,
		ReadContext:   resourceReadExperiment,
		UpdateContext: resourceUpdateExperiment,
		DeleteContext: resourceDeleteExperiment,
		Schema:        experimentSchema(),
	}
}

func resourceCreateExperiment(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := dataFromExperimentResource(rd)
	if err != nil {
		return diag.FromErr(err)
	}
	return makeAPICallAndPopulateResource(m.(string), "/experiments", "POST", data, rd, populateExperimentResourceFromResponse)
}

func resourceReadExperiment(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("/experiments/%s", rd.Get("id"))
	return makeAPICallAndPopulateResource(m.(string), e, "GET", nil, rd, populateExperimentResourceFromResponse)
}

func resourceUpdateExperiment(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := dataFromExperimentResource(rd)
	if err != nil {
		return diag.FromErr(err)
	}
	e := fmt.Sprintf("/experiments/%s", rd.Get("id"))
	return makeAPICallAndPopulateResource(m.(string), e, "POST", data, rd, populateExperimentResourceFromResponse)
}

func resourceDeleteExperiment(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("/experiments/%s", rd.Get("id"))
	d := makeAPICallAndPopulateResource(m.(string), e, "DELETE", nil, nil, populateExperimentResourceFromResponse)
	if d == nil {
		rd.SetId("")
	}
	return d
}

func dataFromExperimentResource(rd *schema.ResourceData) ([]byte, error) {
	body := map[string]interface{}{
		"name":        rd.Get("name"),
		"description": rd.Get("description"),
		"idType":      rd.Get("id_type"),
	}

	return json.Marshal(body)
}

func populateExperimentResourceFromResponse(rd *schema.ResourceData, r map[string]interface{}) {
	rd.Set("description", r["description"])
	rd.Set("last_modifier_name", r["lastModifierName"])
	rd.Set("last_modifier_id", r["lastModifierID"])

	rd.SetId(r["id"].(string))
}
