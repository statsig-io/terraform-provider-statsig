package statsig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ListToStringSlice(ctx context.Context, list basetypes.ListValue) []string {
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

func StringSliceToNilableValue(ctx context.Context, diags diag.Diagnostics, list []string) basetypes.ListValue {
	if list == nil {
		return types.ListNull(types.StringType)
	} else {
		v, d := types.ListValueFrom(ctx, types.StringType, list)
		diags = append(diags, d...)
		return v
	}
}
