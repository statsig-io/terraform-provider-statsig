package resource_keys

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The three requests the keys endpoints read from the target app fields are told
// apart by the serialized body, not by the Go value, so this asserts the bytes.
// An absent field leaves the assignment alone, a null or an empty array clears
// it, and a value sets it.
func TestKeyToAPIInputModelSendsTargetAppsThreeWays(t *testing.T) {
	const targetAppId = "4SRgGcr8uWNVW3c2OGWFZC"
	const secondaryTargetAppId = "2Kd9hLpQzXcVbNmR4TsYuI"

	cases := []struct {
		name string
		key  *KeysModel
		// An empty want means the field must not appear in the body at all.
		wantTargetAppId string
		wantSecondaries string
	}{
		{
			name: "a configured value is sent",
			key: &KeysModel{
				TargetAppId: types.StringValue(targetAppId),
				SecondaryTargetAppIds: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(secondaryTargetAppId),
				}),
			},
			wantTargetAppId: `"` + targetAppId + `"`,
			wantSecondaries: `["` + secondaryTargetAppId + `"]`,
		},
		{
			name: "an explicitly empty value clears the assignment",
			key: &KeysModel{
				TargetAppId:           types.StringValue(""),
				SecondaryTargetAppIds: types.ListValueMust(types.StringType, []attr.Value{}),
			},
			wantTargetAppId: "null",
			wantSecondaries: "[]",
		},
		{
			name: "an attribute the config omits is left out of the request",
			key: &KeysModel{
				TargetAppId:           types.StringUnknown(),
				SecondaryTargetAppIds: types.ListUnknown(types.StringType),
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			encoded, err := json.Marshal(KeyToAPIInputModel(context.Background(), testCase.key))
			require.NoError(t, err)

			var body map[string]json.RawMessage
			require.NoError(t, json.Unmarshal(encoded, &body))

			assertAPIField(t, body, "targetAppID", testCase.wantTargetAppId)
			assertAPIField(t, body, "secondaryTargetAppIDs", testCase.wantSecondaries)

			assert.NotEqual(t, `""`, string(body["targetAppID"]),
				"the API stores an empty string as an identifier instead of clearing the assignment")
		})
	}
}

func assertAPIField(t *testing.T, body map[string]json.RawMessage, field string, want string) {
	t.Helper()

	raw, present := body[field]
	if want == "" {
		assert.False(t, present, "%s must be absent so the API leaves the assignment alone", field)
		return
	}

	require.True(t, present, "%s must be sent", field)
	assert.JSONEq(t, want, string(raw))
}

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
// has to be resolved before the value is written to state. secondary_target_app_ids
// resolves to an empty list, not null: the API's empty secondaryTargetApps array
// used to produce an empty list, and length() and for_each error on a null one.
//
// terraform import leaves every attribute but key null rather than unknown, so
// an unset list has to resolve the same way from either state. Otherwise
// length() and for_each work after an apply and error after an import.
func TestKeyFromAPIInputModelResolvesUnsetTargetApps(t *testing.T) {
	cases := []struct {
		name                  string
		targetAppId           types.String
		secondaryTargetAppIds types.List
	}{
		{
			name:                  "the plan left them unknown",
			targetAppId:           types.StringUnknown(),
			secondaryTargetAppIds: types.ListUnknown(types.StringType),
		},
		{
			name:                  "an import left them null",
			targetAppId:           types.StringNull(),
			secondaryTargetAppIds: types.ListNull(types.StringType),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			key := &KeysModel{
				TargetAppId:           testCase.targetAppId,
				SecondaryTargetAppIds: testCase.secondaryTargetAppIds,
			}

			KeyFromAPIInputModel(context.Background(), diag.Diagnostics{}, key, KeysAPIOutputModel{
				Key:              "secret-abc",
				PrimaryTargetApp: "My Edge App",
			})

			assert.True(t, key.TargetAppId.IsNull())
			assert.False(t, key.SecondaryTargetAppIds.IsNull(), "a null list breaks length() and for_each")
			assert.Equal(t, types.ListValueMust(types.StringType, []attr.Value{}), key.SecondaryTargetAppIds)
		})
	}
}
