package utils

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// APIField renders one field of a Console API request body.
//
// The Console API leaves a field alone only when the body leaves it out. Any
// present value, "" and false and [] included, is read as an instruction to set
// the field. Terraform marks an Optional+Computed attribute unknown when the
// configuration does not mention it, so an unknown plan value is what
// "unspecified" means here, and serializing its Go zero value would clear a
// value the user never asked to change.
//
// One rule, three states: an unspecified attribute is left out of the body, a
// known empty value is sent and clears the field, and a known value is sent and
// sets it.
//
// A nil result is what leaves the field out, so every request field built this
// way is a json.RawMessage tagged omitempty.
func APIField(planValue attr.Value, value any) json.RawMessage {
	if planValue == nil || planValue.IsUnknown() || planValue.IsNull() {
		return nil
	}

	encoded, _ := json.Marshal(value)
	return encoded
}

func StringAPIField(value basetypes.StringValue) json.RawMessage {
	return APIField(value, value.ValueString())
}

func BoolAPIField(value basetypes.BoolValue) json.RawMessage {
	return APIField(value, value.ValueBool())
}

func Int64APIField(value basetypes.Int64Value) json.RawMessage {
	return APIField(value, value.ValueInt64())
}

func FloatAPIField(value basetypes.Float64Value) json.RawMessage {
	return APIField(value, value.ValueFloat64())
}

func NumberAPIField(value basetypes.NumberValue) json.RawMessage {
	return APIField(value, IntFromNumberValue(value))
}

func StringSliceAPIField(ctx context.Context, value basetypes.ListValue) json.RawMessage {
	return APIField(value, StringSliceFromListValue(ctx, value))
}

func MapAPIField(ctx context.Context, value basetypes.MapValue) json.RawMessage {
	return APIField(value, MapFromMapValue(ctx, value))
}
