package resource_entity_property

import (
	"context"

	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for EntityPropertyModel
type EntityPropertyAPIModel struct {
	Description     string                  `json:"description"`
	IdTypeMapping   []IdTypeMappingAPIModel `json:"idTypeMapping,omitempty"`
	IsReadOnly      *bool                   `json:"isReadOnly,omitempty"`
	Name            string                  `json:"name"`
	Owner           *OwnerAPIModel          `json:"owner,omitempty"`
	Sql             string                  `json:"sql"`
	Tags            []string                `json:"tags"`
	TimestampAsDay  *bool                   `json:"timestampAsDay,omitempty"`
	TimestampColumn string                  `json:"timestampColumn,omitempty"`
}

func EntityPropertyToAPIModel(ctx context.Context, entity_property *EntityPropertyModel) EntityPropertyAPIModel {
	return EntityPropertyAPIModel{
		Description:     utils.StringFromNilableValue(entity_property.Description),
		IdTypeMapping:   IdTypeMappingsToAPIModel(ctx, entity_property.IdTypeMapping),
		IsReadOnly:      utils.NilableBoolFromBoolValue(entity_property.IsReadOnly),
		Name:            utils.StringFromNilableValue(entity_property.Name),
		Owner:           OwnerToAPIModel(ctx, entity_property.Owner),
		Sql:             utils.StringFromNilableValue(entity_property.Sql),
		Tags:            utils.StringSliceFromListValue(ctx, entity_property.Tags),
		TimestampAsDay:  utils.NilableBoolFromBoolValue(entity_property.TimestampAsDay),
		TimestampColumn: utils.StringFromNilableValue(entity_property.TimestampColumn),
	}
}

func EntityPropertyFromAPIModel(ctx context.Context, diags diag.Diagnostics, entity_property *EntityPropertyModel, res EntityPropertyAPIModel) {
	entity_property.Description = utils.StringToNilableValue(res.Description)
	entity_property.IdTypeMapping = IdTypeMappingsFromAPIModel(ctx, diags, res.IdTypeMapping)
	entity_property.IsReadOnly = utils.NilableBoolToBoolValue(res.IsReadOnly)
	entity_property.Name = utils.StringToNilableValue(res.Name)
	entity_property.Owner = OwnerFromAPIModel(ctx, diags, res.Owner)
	entity_property.Sql = utils.StringToNilableValue(res.Sql)
	entity_property.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	entity_property.TimestampAsDay = utils.NilableBoolToBoolValue(res.TimestampAsDay)
	entity_property.TimestampColumn = utils.StringToNilableValue(res.TimestampColumn)
}

type IdTypeMappingAPIModel struct {
	Column        string `json:"column"`
	StatsigUnitId string `json:"statsigUnitID"`
}

func IdTypeMappingToAPIModel(ctx context.Context, idTypeMapping *IdTypeMappingValue) IdTypeMappingAPIModel {
	return IdTypeMappingAPIModel{
		Column:        utils.StringFromNilableValue(idTypeMapping.Column),
		StatsigUnitId: utils.StringFromNilableValue(idTypeMapping.StatsigUnitId),
	}
}

func IdTypeMappingFromAPIModel(ctx context.Context, diags diag.Diagnostics, idTypeMapping *IdTypeMappingValue, res IdTypeMappingAPIModel) {
	idTypeMapping.Column = utils.StringToNilableValue(res.Column)
	idTypeMapping.StatsigUnitId = utils.StringToNilableValue(res.StatsigUnitId)
}

func IdTypeMappingsToAPIModel(ctx context.Context, list basetypes.ListValue) []IdTypeMappingAPIModel {
	var res []IdTypeMappingAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]IdTypeMappingAPIModel, 0)
	} else {
		res = make([]IdTypeMappingAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(IdTypeMappingValue)
			if !ok {
				return nil
			}

			res[i] = IdTypeMappingToAPIModel(ctx, &obj)
		}
	}
	return res
}

func IdTypeMappingsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []IdTypeMappingAPIModel) basetypes.ListValue {
	attrTypes := IdTypeMappingValue{}.AttributeTypes(ctx)
	idTypeMappingType := IdTypeMappingType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(idTypeMappingType)
	} else {
		idTypeMappings := make([]attr.Value, len(list))
		for i, elem := range list {
			var idTypeMapping IdTypeMappingValue
			IdTypeMappingFromAPIModel(ctx, diags, &idTypeMapping, elem)
			obj, d := idTypeMapping.ToObjectValue(ctx)
			idTypeMappings[i] = NewIdTypeMappingValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(idTypeMappingType, idTypeMappings)
		diags = append(diags, d...)
		return v
	}
}

type OwnerAPIModel struct {
	OwnerEmail string `json:"ownerEmail,omitempty"`
	OwnerId    string `json:"ownerID,omitempty"`
	OwnerName  string `json:"ownerName,omitempty"`
	OwnerType  string `json:"ownerType,omitempty"`
}

func OwnerToAPIModel(ctx context.Context, owner OwnerValue) *OwnerAPIModel {
	if owner.IsNull() {
		return nil
	}
	return &OwnerAPIModel{
		OwnerEmail: utils.StringFromNilableValue(owner.OwnerEmail),
		OwnerId:    utils.StringFromNilableValue(owner.OwnerId),
		OwnerName:  utils.StringFromNilableValue(owner.OwnerName),
		OwnerType:  utils.StringFromNilableValue(owner.OwnerType),
	}
}

func OwnerFromAPIModel(ctx context.Context, diags diag.Diagnostics, owner *OwnerAPIModel) OwnerValue {
	if owner == nil {
		return NewOwnerValueNull()
	}

	var res OwnerValue
	res.OwnerEmail = utils.StringToNilableValue(owner.OwnerEmail)
	res.OwnerId = utils.StringToNilableValue(owner.OwnerId)
	res.OwnerName = utils.StringToNilableValue(owner.OwnerName)
	res.OwnerType = utils.StringToNilableValue(owner.OwnerType)
	obj, d := res.ToObjectValue(ctx)
	diags = append(diags, d...)
	return NewOwnerValueMust(
		OwnerValue{}.AttributeTypes(ctx),
		obj.Attributes(),
	)
}
