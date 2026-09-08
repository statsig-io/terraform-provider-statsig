package resource_keys

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

// The keys endpoints take targetAppID but answer with primaryTargetApp, a
// display name. Writing that name into target_app_id made the applied state
// differ from the plan, which aborts the apply.
func TestKeyFromAPIInputModelKeepsConfiguredTargetApps(t *testing.T) {
	secondary := types.ListValueMust(types.StringType, []attr.Value{
		types.StringValue("2Kd9hLpQzXcVbNmR4TsYuI"),
	})
	key := &KeysModel{
		TargetAppId:           types.StringValue("4SRgGcr8uWNVW3c2OGWFZC"),
		SecondaryTargetAppIds: secondary,
	}

	KeyFromAPIInputModel(context.Background(), diag.Diagnostics{}, key, KeysAPIOutputModel{
		Key:                 "secret-abc",
		Type:                "SERVER",
		Description:         "edge server key",
		PrimaryTargetApp:    "My Edge App",
		SecondaryTargetApps: []string{"My Other App"},
	})

	assert.Equal(t, "4SRgGcr8uWNVW3c2OGWFZC", key.TargetAppId.ValueString())
	assert.Equal(t, secondary, key.SecondaryTargetAppIds)
	assert.Equal(t, "secret-abc", key.Key.ValueString())
}

// Both target app attributes are computed, so an unknown left by the plan still
// has to be resolved before the value is written to state.
func TestKeyFromAPIInputModelResolvesUnknownTargetApps(t *testing.T) {
	key := &KeysModel{
		TargetAppId:           types.StringUnknown(),
		SecondaryTargetAppIds: types.ListUnknown(types.StringType),
	}

	KeyFromAPIInputModel(context.Background(), diag.Diagnostics{}, key, KeysAPIOutputModel{
		Key:              "secret-abc",
		PrimaryTargetApp: "My Edge App",
	})

	assert.True(t, key.TargetAppId.IsNull())
	assert.True(t, key.SecondaryTargetAppIds.IsNull())
}
