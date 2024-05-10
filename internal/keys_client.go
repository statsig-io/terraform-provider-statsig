package statsig

import (
	"context"
	"terraform-provider-statsig/internal/resource_keys"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type keysClient struct {
	endpoint  string
	transport *transport
}

func newKeysCient(transport *transport) *keysClient {
	return &keysClient{
		endpoint:  "keys",
		transport: transport,
	}
}

// API data model for KeysModel (NOTE: see if we can get Terraform to also codegen this from OpenAPI)
type KeysAPIInputModel struct {
	Description           string   `json:"description"`
	Environments          []string `json:"environments"`
	Scopes                []string `json:"scopes"`
	SecondaryTargetAppIds []string `json:"secondary_target_app_ids"`
	TargetAppId           string   `json:"target_app_id"`
	Type                  string   `json:"type"`
}

type KeysAPIOutputModel struct {
	Description         string   `json:"description"`
	Environments        []string `json:"environments"`
	Key                 string   `json:"key"`
	Scopes              []string `json:"scopes"`
	SecondaryTargetApps []string `json:"secondaryTargetApps"`
	PrimaryTargetApp    string   `json:"primaryTargetApp"`
	Type                string   `json:"type"`
}

func toAPIInputModel(ctx context.Context, key *resource_keys.KeysModel) KeysAPIInputModel {
	return KeysAPIInputModel{
		Description:           key.Description.ValueString(),
		Environments:          ListToStringSlice(ctx, key.Environments),
		Scopes:                ListToStringSlice(ctx, key.Scopes),
		SecondaryTargetAppIds: ListToStringSlice(ctx, key.SecondaryTargetAppIds),
		TargetAppId:           key.TargetAppId.ValueString(),
		Type:                  key.Type.ValueString(),
	}
}

func fromAPIOutputModel(ctx context.Context, diags diag.Diagnostics, key *resource_keys.KeysModel, res KeysAPIOutputModel) {
	key.Key = StringToNilableValue(res.Key)
	key.Type = StringToNilableValue(res.Type)
	key.Description = StringToNilableValue(res.Description)
	key.TargetAppId = StringToNilableValue(res.PrimaryTargetApp)
	key.Environments = StringSliceToNilableValue(ctx, diags, res.Environments)
	key.Scopes = StringSliceToNilableValue(ctx, diags, res.Scopes)
	key.SecondaryTargetAppIds = StringSliceToNilableValue(ctx, diags, res.SecondaryTargetApps)
}

func (c *keysClient) read(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	if key.Key.IsUnknown() {
		key.Key = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data KeysAPIOutputModel
		res, err := c.transport.get(c.endpoint, key.Key.ValueString(), &data)
		fromAPIOutputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) create(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data KeysAPIOutputModel
		res, err := c.transport.post(c.endpoint, toAPIInputModel(ctx, key), &data)
		fromAPIOutputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) update(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	if key.Key.IsUnknown() {
		key.Key = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data KeysAPIOutputModel
		res, err := c.transport.patch(c.endpoint, key.Key.ValueString(), toAPIInputModel(ctx, key), &data)
		fromAPIOutputModel(ctx, diags, key, data)
		return res, err
	})
}

func (c *keysClient) delete(ctx context.Context, key *resource_keys.KeysModel) diag.Diagnostics {
	if key.Key.IsUnknown() {
		key.Key = types.StringNull()
	}

	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data KeysAPIOutputModel
		return c.transport.delete(c.endpoint, key.Key.ValueString(), toAPIInputModel(ctx, key), &data)
	})
}
