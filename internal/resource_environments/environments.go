package resource_environments

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
)

// API data model for EnvironmentsModel
type EnvironmentsAPIModel struct {
	Environments	[]EnvironmentAPIModel	`json:"environments"`
}

func EnvironmentsToAPIModel(ctx context.Context, environments *EnvironmentsModel) EnvironmentsAPIModel {
	var res EnvironmentsAPIModel
	res.Environments = EnvironmentListToAPIModel(ctx, environments.Environments)
	return res
}

func EnvironmentsFromAPIModel(ctx context.Context, diags diag.Diagnostics, environments *EnvironmentsModel, res EnvironmentsAPIModel) {
	environments.Environments = EnvironmentListFromAPIModel(ctx, diags, res.Environments)
}

type EnvironmentAPIModel struct {
	ID						string			`json:"id"`
	Name					string			`json:"name"`
	IsProduction			bool			`json:"isProduction"`
	RequiredReviewGroupID	string			`json:"requiredReviewGroupID"`
	RequiresReleasePipeline	bool			`json:"requiresReleasePipeline"`
	RequiresReview			bool			`json:"requiresReview"`
}

func EnvironmentToAPIModel(ctx context.Context, environment *EnvironmentsValue) EnvironmentAPIModel {
	return EnvironmentAPIModel{
		ID:							utils.StringFromNilableValue(environment.Id),
		Name:						utils.StringFromNilableValue(environment.Name),
		IsProduction:				utils.BoolFromBoolValue(environment.IsProduction),
		RequiredReviewGroupID:		utils.StringFromNilableValue(environment.RequiredReviewGroupId),
		RequiresReleasePipeline:	utils.BoolFromBoolValue(environment.RequiresReleasePipeline),
		RequiresReview:				utils.BoolFromBoolValue(environment.RequiresReview),
	}
}

func EnvironmentFromAPIModel(ctx context.Context, diags diag.Diagnostics, environment *EnvironmentsValue, res EnvironmentAPIModel) {
	environment.Id = utils.StringToNilableValue(res.ID)
	environment.Name = utils.StringToNilableValue(res.Name)
	environment.IsProduction = utils.BoolToBoolValue(res.IsProduction)
	environment.RequiredReviewGroupId = utils.StringToNilableValue(res.RequiredReviewGroupID)
	environment.RequiresReleasePipeline = utils.BoolToBoolValue(res.RequiresReleasePipeline)
	environment.RequiresReview = utils.BoolToBoolValue(res.RequiresReview)
}

func EnvironmentListToAPIModel(ctx context.Context, list basetypes.ListValue) []EnvironmentAPIModel {
	var res []EnvironmentAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]EnvironmentAPIModel, 0)
	} else {
		res = make([]EnvironmentAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(EnvironmentsValue)
			if !ok {
				return nil
			}

			res[i] = EnvironmentToAPIModel(ctx, &obj)
		}
	}
	return res
}

func EnvironmentListFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []EnvironmentAPIModel) basetypes.ListValue {
	attrTypes := EnvironmentsValue{}.AttributeTypes(ctx)
	environmentType := EnvironmentsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(environmentType)
	} else {
		environments := make([]attr.Value, len(list))
		for i, elem := range list {
			var environment EnvironmentsValue
			EnvironmentFromAPIModel(ctx, diags, &environment, elem)
			obj, d := environment.ToObjectValue(ctx)
			environments[i] = NewEnvironmentsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(environmentType, environments)
		diags = append(diags, d...)
		return v
	}
}
