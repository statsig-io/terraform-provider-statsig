package resource_keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// The keys endpoints take targetAppID / secondaryTargetAppIDs on the way in but
// answer with primaryTargetApp / secondaryTargetApps, which are display names.
// No response carries the IDs back, so those two fields are read but not mapped
// into state. See KeyFromAPIInputModel.
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
	key.Environments = utils.StringSliceToListValue(ctx, diags, res.Environments)
	key.Scopes = utils.StringSliceToListValue(ctx, diags, res.Scopes)

	// Keep whatever the plan holds for the target app attributes: writing the
	// response's display names over configured IDs makes the applied state
	// inconsistent with the plan and fails the apply. Both attributes are
	// computed, so an unknown still has to be resolved. An unset list resolves to
	// an empty list, the shape the API's empty secondaryTargetApps used to
	// produce, so length() and for_each over it keep working.
	if key.TargetAppId.IsUnknown() {
		key.TargetAppId = types.StringNull()
	}
	if key.SecondaryTargetAppIds.IsUnknown() {
		key.SecondaryTargetAppIds = types.ListValueMust(types.StringType, []attr.Value{})
	}
}
