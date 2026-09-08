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
// The two target app fields are raw JSON because the keys endpoints read three
// distinct requests from them, and a plain Go string can only express two. See
// targetAppIdToAPIField.
type KeysAPIInputModel struct {
	Description           string          `json:"description"`
	Environments          []string        `json:"environments"`
	Scopes                []string        `json:"scopes"`
	SecondaryTargetAppIds json.RawMessage `json:"secondaryTargetAppIDs,omitempty"`
	TargetAppId           json.RawMessage `json:"targetAppID,omitempty"`
	Type                  string          `json:"type"`
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
		SecondaryTargetAppIds: secondaryTargetAppIdsToAPIField(ctx, key.SecondaryTargetAppIds),
		TargetAppId:           targetAppIdToAPIField(key.TargetAppId),
		Type:                  key.Type.ValueString(),
	}
}

// The keys endpoints read the target app fields three ways: an absent field
// leaves the current assignment alone, a null clears it, and a value sets it. An
// empty string is not a clear. The API stores it as an identifier and writes a
// reverse association keyed on it, so it must never be sent.
//
// A config that omits the attribute leaves the plan value unknown, which is the
// absent case. An attribute set to "" is the clear.
func targetAppIdToAPIField(value types.String) json.RawMessage {
	if value.IsUnknown() || value.IsNull() {
		return nil
	}
	if value.ValueString() == "" {
		return json.RawMessage("null")
	}
	encoded, _ := json.Marshal(value.ValueString())
	return encoded
}

// An empty array clears the assignment, so the list needs no null case.
func secondaryTargetAppIdsToAPIField(ctx context.Context, value types.List) json.RawMessage {
	if value.IsUnknown() || value.IsNull() {
		return nil
	}
	encoded, _ := json.Marshal(utils.StringSliceFromListValue(ctx, value))
	return encoded
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
