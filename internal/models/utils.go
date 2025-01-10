package models

import (
	"context"
	"math/big"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func SliceFromListValue(ctx context.Context, list types.List) []interface{} {
	var res []interface{}
	if list.IsNull() || list.IsUnknown() {
		res = make([]interface{}, 0)
	} else {
		res = make([]interface{}, len(list.Elements()))
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

func Int64ToInt64Value(value int64) basetypes.Int64Value {
	return types.Int64Value(value)
}

func Int64FromInt64Value(value basetypes.Int64Value) int64 {
	return value.ValueInt64()
}

func FloatToFloatValue(value float64) basetypes.Float64Value {
	return types.Float64Value(value)
}

func FloatFromFloatValue(value types.Float64) float64 {
	return value.ValueFloat64()
}

func FloatToNumberValue(value float64) types.Number {
	return types.NumberValue(big.NewFloat(value))
}

func FloatFromNumberValue(value types.Number) float64 {
	if value.IsNull() || value.IsUnknown() {
		return 0
	}
	val, _ := value.ValueBigFloat().Float64()
	return val
}

func BoolToBoolValue(value bool) basetypes.BoolValue {
	return types.BoolValue(value)
}

func BoolFromBoolValue(value basetypes.BoolValue) bool {
	return value.ValueBool()
}

func MapFromMapValue(ctx context.Context, value basetypes.MapValue) map[string]interface{} {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}
	attributes := value.Elements()
	res := make(map[string]interface{}, len(attributes))
	for attrName, attribute := range attributes {
		res[attrName] = InterfaceFromValue(ctx, attribute)
	}
	return res
}

func MapToMapValue(ctx context.Context, diags diag.Diagnostics, value map[string]interface{}) basetypes.MapValue {
	var attrType attr.Type
	attributes := make(map[string]attr.Value, len(value))
	for key, val := range value {
		v := InterfaceToValue(ctx, diags, val)
		attrType = v.Type(ctx)
		attributes[key] = v
	}

	objValues, diag := types.MapValue(attrType, attributes)
	diags = append(diags, diag...)
	return objValues
}

func InterfaceToValue(ctx context.Context, diags diag.Diagnostics, value interface{}) attr.Value {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return StringToNilableValue(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int64ToInt64Value(v.Int())
	case reflect.Float32, reflect.Float64:
		return FloatToFloatValue(v.Float())
	case reflect.Bool:
		return BoolToBoolValue(v.Bool())
	case reflect.Slice, reflect.Array:
		var elems []attr.Value
		for _, item := range v.Interface().([]interface{}) {
			elems = append(elems, InterfaceToValue(ctx, diags, item))
		}
		val, _ := types.ListValue(types.StringType, elems)
		return val
	case reflect.Map:
		return MapToMapValue(ctx, diags, v.Interface().(map[string]interface{}))
	default:
		return types.StringUnknown()
	}
}

func InterfaceFromValue(ctx context.Context, value attr.Value) interface{} {
	if value.Type(ctx).Equal(types.StringType) {
		return StringFromNilableValue(value.(types.String))
	} else if value.Type(ctx).Equal(types.Int64Type) {
		return Int64FromInt64Value(value.(types.Int64))
	} else if value.Type(ctx).Equal(types.Float64Type) {
		return FloatFromFloatValue(value.(types.Float64))
	} else if value.Type(ctx).Equal(types.NumberType) {
		return FloatFromNumberValue(value.(types.Number))
	} else if value.Type(ctx).Equal(types.BoolType) {
		return BoolFromBoolValue(value.(types.Bool))
	} else if value.Type(ctx).Equal(types.ListType{}) {
		return SliceFromListValue(ctx, value.(types.List))
	} else if value.Type(ctx).Equal(types.ObjectType{}) {
		return MapFromMapValue(ctx, value.(types.Map))
	} else {
		return nil
	}
}
