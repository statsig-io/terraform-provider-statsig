package statsig

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type statsigResource struct {
	endpoint       string
	schema         map[string]*schema.Schema
	fromJsonObject func(rd *schema.ResourceData, r map[string]interface{})
	toJsonData     func(rd *schema.ResourceData) ([]byte, error)
}

func (b statsigResource) resourceCreate(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := b.toJsonData(rd)
	if err != nil {
		return diag.FromErr(err)
	}
	result := makeAPICallAndPopulateResource(m.(string), b.endpoint, "POST", data, rd, b.fromJsonObject)
	return result
}

func (b statsigResource) resourceRead(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("%s/%s", b.endpoint, rd.Get("id"))
	return makeAPICallAndPopulateResource(m.(string), e, "GET", nil, rd, b.fromJsonObject)
}

func (b statsigResource) resourceUpdate(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	data, err := dataFromExperimentResource(rd)
	if err != nil {
		return diag.FromErr(err)
	}
	e := fmt.Sprintf("%s/%s", b.endpoint, rd.Get("id"))
	return makeAPICallAndPopulateResource(m.(string), e, "POST", data, rd, b.fromJsonObject)
}

func (b statsigResource) resourceDelete(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	e := fmt.Sprintf("%s/%s", b.endpoint, rd.Get("id"))
	d := makeAPICallAndPopulateResource(m.(string), e, "DELETE", nil, nil, b.fromJsonObject)
	if d == nil {
		rd.SetId("")
	}
	return d
}

func (b statsigResource) asTerraformResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: b.resourceCreate,
		ReadContext:   b.resourceRead,
		UpdateContext: b.resourceUpdate,
		DeleteContext: b.resourceDelete,
		Schema:        b.schema,
	}
}
