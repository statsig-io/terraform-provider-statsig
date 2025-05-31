package models

import (
	"context"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_keys"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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

func KeyToAPIInputModel(ctx context.Context, key *resource_keys.KeysModel) KeysAPIInputModel {
	return KeysAPIInputModel{
		Description:           key.Description.ValueString(),
		Environments:          StringSliceFromListValue(ctx, key.Environments),
		Scopes:                StringSliceFromListValue(ctx, key.Scopes),
		SecondaryTargetAppIds: StringSliceFromListValue(ctx, key.SecondaryTargetAppIds),
		TargetAppId:           key.TargetAppId.ValueString(),
		Type:                  key.Type.ValueString(),
	}
}

func KeyFromAPIInputModel(ctx context.Context, diags diag.Diagnostics, key *resource_keys.KeysModel, res KeysAPIOutputModel) {
	key.Key = StringToNilableValue(res.Key)
	key.Type = StringToNilableValue(res.Type)
	key.Description = StringToNilableValue(res.Description)
	key.TargetAppId = StringToNilableValue(res.PrimaryTargetApp)
	key.Environments = StringSliceToListValue(ctx, diags, res.Environments)
	key.Scopes = StringSliceToListValue(ctx, diags, res.Scopes)
	key.SecondaryTargetAppIds = StringSliceToListValue(ctx, diags, res.SecondaryTargetApps)
}
