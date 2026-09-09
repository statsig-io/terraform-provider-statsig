package resource_keys

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
)

// API data model for KeysModel (NOTE: see if we can get Terraform to also codegen this from OpenAPI)
//
// Every field is raw JSON so the request can leave out an attribute the
// configuration never mentioned. See utils.APIField.
type KeysAPIInputModel struct {
	Description           json.RawMessage `json:"description,omitempty"`
	Environments          json.RawMessage `json:"environments,omitempty"`
	Scopes                json.RawMessage `json:"scopes,omitempty"`
	SecondaryTargetAppIds json.RawMessage `json:"secondaryTargetAppIDs,omitempty"`
	TargetAppId           json.RawMessage `json:"targetAppID,omitempty"`
	Type                  json.RawMessage `json:"type,omitempty"`
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
		Description:           utils.StringAPIField(key.Description),
		Environments:          utils.StringSliceAPIField(ctx, key.Environments),
		Scopes:                utils.StringSliceAPIField(ctx, key.Scopes),
		SecondaryTargetAppIds: utils.StringSliceAPIField(ctx, key.SecondaryTargetAppIds),
		TargetAppId:           targetAppIdToAPIField(key.TargetAppId),
		Type:                  utils.StringAPIField(key.Type),
	}
}

// target_app_id is the one attribute whose clear is not its own empty value. The
// keys endpoints store an empty string as an identifier and write a reverse
// association keyed on it, so "" must never be sent; the clear is a JSON null.
// The absent and set states follow the shared rule in utils.APIField.
func targetAppIdToAPIField(value types.String) json.RawMessage {
	if !value.IsUnknown() && !value.IsNull() && value.ValueString() == "" {
		return json.RawMessage("null")
	}
	return utils.StringAPIField(value)
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
