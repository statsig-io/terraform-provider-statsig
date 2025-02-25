package models

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func StringSliceToListValue(ctx context.Context, diags diag.Diagnostics, list []string) basetypes.ListValue {
	if list == nil {
		return types.ListNull(types.StringType)
	} else {
		v, d := types.ListValueFrom(ctx, types.StringType, list)
		diags = append(diags, d...)
		return v
	}
}

func StringSliceFromListValue(ctx context.Context, list basetypes.ListValue) []string {
	var res []string
	if list.IsNull() || list.IsUnknown() {
		res = make([]string, 0)
	} else {
		res = make([]string, len(list.Elements()))
		list.ElementsAs(ctx, &res, false)
	}
	return res
}

func StringToNilableValue(str string) basetypes.StringValue {
	if str == "" {
		return types.StringNull()
	} else {
		return types.StringValue(str)
	}
}

func StringFromNilableValue(str basetypes.StringValue) string {
	if str.IsNull() || str.IsUnknown() {
		return ""
	} else {
		return str.ValueString()
	}
}

func IntToNumberValue(value int) basetypes.NumberValue {
	return types.NumberValue(big.NewFloat(float64(value)))
}

func IntFromNumberValue(num basetypes.NumberValue) int {
	if num.IsNull() || num.IsUnknown() {
		return 0
	}
	val, _ := num.ValueBigFloat().Int64()
	return int(val)
}

func BoolToBoolValue(value bool) basetypes.BoolValue {
	return types.BoolValue(value)
}

func BoolFromBoolValue(value basetypes.BoolValue) bool {
	return value.ValueBool()
}

func NilableBoolToBoolValue(value *bool) basetypes.BoolValue {
	if value == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*value)
}

func NilableBoolFromBoolValue(value basetypes.BoolValue) *bool {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	return value.ValueBoolPointer()
}
