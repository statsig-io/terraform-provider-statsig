package resource_qualifying_event

import (
	"context"
	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for QualifyingEventModel
type QualifyingEventAPIModel struct {
	CustomFieldMapping []CustomFieldMappingAPIModel `json:"customFieldMapping,omitempty"`
	Description        string                       `json:"description"`
	IdTypeMapping      []IdTypeMappingAPIModel      `json:"idTypeMapping,omitempty"`
	IsReadOnly         *bool                        `json:"isReadOnly,omitempty"`
	Name               string                       `json:"name"`
	Owner              *OwnerAPIModel               `json:"owner,omitempty"`
	SourceType         string                       `json:"sourceType,omitempty"`
	Sql                string                       `json:"sql"`
	TableName          string                       `json:"tableName,omitempty"`
	Tags               []string                     `json:"tags"`
	TimestampAsDay     *bool                        `json:"timestampAsDay,omitempty"`
	TimestampColumn    string                       `json:"timestampColumn,omitempty"`
}

func QualifyingEventToAPIModel(ctx context.Context, metric_source *QualifyingEventModel) QualifyingEventAPIModel {
	return QualifyingEventAPIModel{
		CustomFieldMapping: CustomFieldMappingsToAPIModel(ctx, metric_source.CustomFieldMapping),
		Description:        utils.StringFromNilableValue(metric_source.Description),
		IdTypeMapping:      IdTypeMappingsToAPIModel(ctx, metric_source.IdTypeMapping),
		IsReadOnly:         utils.NilableBoolFromBoolValue(metric_source.IsReadOnly),
		Name:               utils.StringFromNilableValue(metric_source.Name),
		Owner:              OwnerToAPIModel(ctx, metric_source.Owner),
		SourceType:         utils.StringFromNilableValue(metric_source.SourceType),
		Sql:                utils.StringFromNilableValue(metric_source.Sql),
		TableName:          utils.StringFromNilableValue(metric_source.TableName),
		Tags:               utils.StringSliceFromListValue(ctx, metric_source.Tags),
		TimestampAsDay:     utils.NilableBoolFromBoolValue(metric_source.TimestampAsDay),
		TimestampColumn:    utils.StringFromNilableValue(metric_source.TimestampColumn),
	}
}

func QualifyingEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric_source *QualifyingEventModel, res QualifyingEventAPIModel) {
	metric_source.CustomFieldMapping = CustomFieldMappingsFromAPIModel(ctx, diags, res.CustomFieldMapping)
	metric_source.Description = utils.StringToNilableValue(res.Description)
	metric_source.IdTypeMapping = IdTypeMappingsFromAPIModel(ctx, diags, res.IdTypeMapping)
	metric_source.IsReadOnly = utils.NilableBoolToBoolValue(res.IsReadOnly)
	metric_source.Name = utils.StringToNilableValue(res.Name)
	metric_source.Owner = OwnerFromAPIModel(ctx, diags, res.Owner)
	metric_source.SourceType = utils.StringToNilableValue(res.SourceType)
	metric_source.Sql = utils.StringToStringValue(res.Sql)
	metric_source.TableName = utils.StringToNilableValue(res.TableName)
	metric_source.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	metric_source.TimestampAsDay = utils.NilableBoolToBoolValue(res.TimestampAsDay)
	metric_source.TimestampColumn = utils.StringToNilableValue(res.TimestampColumn)
}

type CustomFieldMappingAPIModel struct {
	Formula string `json:"formula"`
	Key     string `json:"key"`
}

func CustomFieldMappingToAPIModel(ctx context.Context, customFieldMapping *CustomFieldMappingValue) CustomFieldMappingAPIModel {
	return CustomFieldMappingAPIModel{
		Formula: utils.StringFromNilableValue(customFieldMapping.Formula),
		Key:     utils.StringFromNilableValue(customFieldMapping.Key),
	}
}

func CustomFieldMappingFromAPIModel(ctx context.Context, diags diag.Diagnostics, customFieldMapping *CustomFieldMappingValue, res CustomFieldMappingAPIModel) {
	customFieldMapping.Formula = utils.StringToNilableValue(res.Formula)
	customFieldMapping.Key = utils.StringToNilableValue(res.Key)
}

func CustomFieldMappingsToAPIModel(ctx context.Context, list types.List) []CustomFieldMappingAPIModel {
	var res []CustomFieldMappingAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]CustomFieldMappingAPIModel, 0)
	} else {
		res = make([]CustomFieldMappingAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(CustomFieldMappingValue)
			if !ok {
				return nil
			}

			res[i] = CustomFieldMappingToAPIModel(ctx, &obj)
		}
	}
	return res
}

func CustomFieldMappingsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []CustomFieldMappingAPIModel) basetypes.ListValue {
	attrTypes := CustomFieldMappingValue{}.AttributeTypes(ctx)
	customFieldMappingType := CustomFieldMappingType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(customFieldMappingType)
	} else {
		customFieldMappings := make([]attr.Value, len(list))
		for i, elem := range list {
			var customFieldMapping CustomFieldMappingValue
			CustomFieldMappingFromAPIModel(ctx, diags, &customFieldMapping, elem)
			obj, d := customFieldMapping.ToObjectValue(ctx)
			customFieldMappings[i] = NewCustomFieldMappingValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(customFieldMappingType, customFieldMappings)
		diags = append(diags, d...)
		return v
	}
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
