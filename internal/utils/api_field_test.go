package utils

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

// The Console API tells "leave this alone" from "set this" by whether the field
// is in the body at all, so the rule is about the encoded bytes rather than the
// Go value. A nil result is what omitempty drops.
func TestAPIFieldSendsThreeStates(t *testing.T) {
	ctx := context.Background()
	filled := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("production")})

	cases := []struct {
		name string
		text types.String
		flag types.Bool
		list types.List
		// An empty want means the field must be left out of the request.
		wantText string
		wantFlag string
		wantList string
	}{
		{
			name: "an attribute the config omits is left out",
			text: types.StringUnknown(),
			flag: types.BoolUnknown(),
			list: types.ListUnknown(types.StringType),
		},
		{
			name: "a null attribute is left out",
			text: types.StringNull(),
			flag: types.BoolNull(),
			list: types.ListNull(types.StringType),
		},
		{
			name: "a known empty value is sent and clears the field",
			text: types.StringValue(""),
			flag: types.BoolValue(false),
			list: types.ListValueMust(types.StringType, []attr.Value{}),

			wantText: `""`,
			wantFlag: "false",
			wantList: "[]",
		},
		{
			name: "a known value is sent",
			text: types.StringValue("edge server key"),
			flag: types.BoolValue(true),
			list: filled,

			wantText: `"edge server key"`,
			wantFlag: "true",
			wantList: `["production"]`,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			assertField(t, StringAPIField(testCase.text), testCase.wantText)
			assertField(t, BoolAPIField(testCase.flag), testCase.wantFlag)
			assertField(t, StringSliceAPIField(ctx, testCase.list), testCase.wantList)
		})
	}
}

func assertField(t *testing.T, field json.RawMessage, want string) {
	t.Helper()

	if want == "" {
		assert.Nil(t, field, "an unspecified field must be nil so omitempty leaves it out")
		return
	}
	assert.JSONEq(t, want, string(field))
}
