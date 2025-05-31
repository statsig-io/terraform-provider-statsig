package resource_keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
)

// API data model for KeysModel (NOTE: see if we can get Terraform to also codegen this from OpenAPI)
type KeysAPIInputModel struct {
	Description           string   `json:"description"`
	Environments          []string `json:"environments"`
	Scopes                []string `json:"scopes"`
	SecondaryTargetAppIds []string `json:"secondaryTargetAppIDs,omitempty"`
	TargetAppId           string   `json:"targetAppID,omitempty"`
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

func KeyToAPIInputModel(ctx context.Context, key *KeysModel) KeysAPIInputModel {
	return KeysAPIInputModel{
		Description:           key.Description.ValueString(),
		Environments:          utils.StringSliceFromListValue(ctx, key.Environments),
		Scopes:                utils.StringSliceFromListValue(ctx, key.Scopes),
		SecondaryTargetAppIds: utils.StringSliceFromListValue(ctx, key.SecondaryTargetAppIds),
		TargetAppId:           key.TargetAppId.ValueString(),
		Type:                  key.Type.ValueString(),
	}
}

func KeyFromAPIInputModel(ctx context.Context, diags diag.Diagnostics, key *KeysModel, res KeysAPIOutputModel) {
	key.Key = utils.StringToNilableValue(res.Key)
	key.Type = utils.StringToNilableValue(res.Type)
	key.Description = utils.StringToNilableValue(res.Description)
	key.TargetAppId = utils.StringToNilableValue(res.PrimaryTargetApp)
	key.Environments = utils.StringSliceToListValue(ctx, diags, res.Environments)
	key.Scopes = utils.StringSliceToListValue(ctx, diags, res.Scopes)
	key.SecondaryTargetAppIds = utils.StringSliceToListValue(ctx, diags, res.SecondaryTargetApps)
}
