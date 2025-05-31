package models

import (
	"context"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_entity_property"

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

func EntityPropertyToAPIModel(ctx context.Context, entity_property *resource_entity_property.EntityPropertyModel) EntityPropertyAPIModel {
	return EntityPropertyAPIModel{
		Description:     StringFromNilableValue(entity_property.Description),
		IdTypeMapping:   IdTypeMappingsToAPIModel(ctx, entity_property.IdTypeMapping),
		IsReadOnly:      NilableBoolFromBoolValue(entity_property.IsReadOnly),
		Name:            StringFromNilableValue(entity_property.Name),
		Owner:           OwnerToAPIModel(ctx, entity_property.Owner),
		Sql:             StringFromNilableValue(entity_property.Sql),
		Tags:            StringSliceFromListValue(ctx, entity_property.Tags),
		TimestampAsDay:  NilableBoolFromBoolValue(entity_property.TimestampAsDay),
		TimestampColumn: StringFromNilableValue(entity_property.TimestampColumn),
	}
}

func EntityPropertyFromAPIModel(ctx context.Context, diags diag.Diagnostics, entity_property *resource_entity_property.EntityPropertyModel, res EntityPropertyAPIModel) {
	entity_property.Description = StringToNilableValue(res.Description)
	entity_property.IdTypeMapping = IdTypeMappingsFromAPIModel(ctx, diags, res.IdTypeMapping)
	entity_property.IsReadOnly = NilableBoolToBoolValue(res.IsReadOnly)
	entity_property.Name = StringToNilableValue(res.Name)
	entity_property.Owner = OwnerFromAPIModel(ctx, diags, res.Owner)
	entity_property.Sql = StringToNilableValue(res.Sql)
	entity_property.Tags = StringSliceToListValue(ctx, diags, res.Tags)
	entity_property.TimestampAsDay = NilableBoolToBoolValue(res.TimestampAsDay)
	entity_property.TimestampColumn = StringToNilableValue(res.TimestampColumn)
}

type IdTypeMappingAPIModel struct {
	Column        string `json:"column"`
	StatsigUnitId string `json:"statsigUnitID"`
}

func IdTypeMappingToAPIModel(ctx context.Context, idTypeMapping *resource_entity_property.IdTypeMappingValue) IdTypeMappingAPIModel {
	return IdTypeMappingAPIModel{
		Column:        StringFromNilableValue(idTypeMapping.Column),
		StatsigUnitId: StringFromNilableValue(idTypeMapping.StatsigUnitId),
	}
}

func IdTypeMappingFromAPIModel(ctx context.Context, diags diag.Diagnostics, idTypeMapping *resource_entity_property.IdTypeMappingValue, res IdTypeMappingAPIModel) {
	idTypeMapping.Column = StringToNilableValue(res.Column)
	idTypeMapping.StatsigUnitId = StringToNilableValue(res.StatsigUnitId)
}

func IdTypeMappingsToAPIModel(ctx context.Context, list basetypes.ListValue) []IdTypeMappingAPIModel {
	var res []IdTypeMappingAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]IdTypeMappingAPIModel, 0)
	} else {
		res = make([]IdTypeMappingAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_entity_property.IdTypeMappingValue)
			if !ok {
				return nil
			}

			res[i] = IdTypeMappingToAPIModel(ctx, &obj)
		}
	}
	return res
}

func IdTypeMappingsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []IdTypeMappingAPIModel) basetypes.ListValue {
	attrTypes := resource_entity_property.IdTypeMappingValue{}.AttributeTypes(ctx)
	idTypeMappingType := resource_entity_property.IdTypeMappingType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(idTypeMappingType)
	} else {
		idTypeMappings := make([]attr.Value, len(list))
		for i, elem := range list {
			var idTypeMapping resource_entity_property.IdTypeMappingValue
			IdTypeMappingFromAPIModel(ctx, diags, &idTypeMapping, elem)
			obj, d := idTypeMapping.ToObjectValue(ctx)
			idTypeMappings[i] = resource_entity_property.NewIdTypeMappingValueMust(attrTypes, obj.Attributes())
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

func OwnerToAPIModel(ctx context.Context, owner resource_entity_property.OwnerValue) *OwnerAPIModel {
	if owner.IsNull() {
		return nil
	}
	return &OwnerAPIModel{
		OwnerEmail: StringFromNilableValue(owner.OwnerEmail),
		OwnerId:    StringFromNilableValue(owner.OwnerId),
		OwnerName:  StringFromNilableValue(owner.OwnerName),
		OwnerType:  StringFromNilableValue(owner.OwnerType),
	}
}

func OwnerFromAPIModel(ctx context.Context, diags diag.Diagnostics, owner *OwnerAPIModel) resource_entity_property.OwnerValue {
	if owner == nil {
		return resource_entity_property.NewOwnerValueNull()
	}

	var res resource_entity_property.OwnerValue
	res.OwnerEmail = StringToNilableValue(owner.OwnerEmail)
	res.OwnerId = StringToNilableValue(owner.OwnerId)
	res.OwnerName = StringToNilableValue(owner.OwnerName)
	res.OwnerType = StringToNilableValue(owner.OwnerType)
	obj, d := res.ToObjectValue(ctx)
	diags = append(diags, d...)
	return resource_entity_property.NewOwnerValueMust(
		resource_entity_property.OwnerValue{}.AttributeTypes(ctx),
		obj.Attributes(),
	)
}
