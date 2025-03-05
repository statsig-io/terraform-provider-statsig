package models

import (
	"context"
	"terraform-provider-statsig/internal/resource_metric_source"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for MetricSourceModel
type MetricSourceAPIModel struct {
	CustomFieldMapping []CustomFieldMappingAPIModel        `json:"customFieldMapping,omitempty"`
	Description        string                              `json:"description"`
	IdTypeMapping      []MetricSourceIdTypeMappingAPIModel `json:"idTypeMapping,omitempty"`
	IsReadOnly         *bool                               `json:"isReadOnly,omitempty"`
	Name               string                              `json:"name"`
	Owner              *MetricSourceOwnerAPIModel          `json:"owner,omitempty"`
	SourceType         string                              `json:"sourceType,omitempty"`
	Sql                string                              `json:"sql"`
	TableName          string                              `json:"tableName,omitempty"`
	Tags               []string                            `json:"tags"`
	TimestampAsDay     *bool                               `json:"timestampAsDay,omitempty"`
	TimestampColumn    string                              `json:"timestampColumn,omitempty"`
}

func MetricSourceToAPIModel(ctx context.Context, metric_source *resource_metric_source.MetricSourceModel) MetricSourceAPIModel {
	return MetricSourceAPIModel{
		CustomFieldMapping: CustomFieldMappingsToAPIModel(ctx, metric_source.CustomFieldMapping),
		Description:        StringFromNilableValue(metric_source.Description),
		IdTypeMapping:      MetricSourceIdTypeMappingsToAPIModel(ctx, metric_source.IdTypeMapping),
		IsReadOnly:         NilableBoolFromBoolValue(metric_source.IsReadOnly),
		Name:               StringFromNilableValue(metric_source.Name),
		Owner:              MetricSourceOwnerToAPIModel(ctx, metric_source.Owner),
		SourceType:         StringFromNilableValue(metric_source.SourceType),
		Sql:                StringFromNilableValue(metric_source.Sql),
		TableName:          StringFromNilableValue(metric_source.TableName),
		Tags:               StringSliceFromListValue(ctx, metric_source.Tags),
		TimestampAsDay:     NilableBoolFromBoolValue(metric_source.TimestampAsDay),
		TimestampColumn:    StringFromNilableValue(metric_source.TimestampColumn),
	}
}

func MetricSourceFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric_source *resource_metric_source.MetricSourceModel, res MetricSourceAPIModel) {
	metric_source.CustomFieldMapping = CustomFieldMappingsFromAPIModel(ctx, diags, res.CustomFieldMapping)
	metric_source.Description = StringToNilableValue(res.Description)
	metric_source.IdTypeMapping = MetricSourceIdTypeMappingsFromAPIModel(ctx, diags, res.IdTypeMapping)
	metric_source.IsReadOnly = NilableBoolToBoolValue(res.IsReadOnly)
	metric_source.Name = StringToNilableValue(res.Name)
	metric_source.Owner = MetricSourceOwnerFromAPIModel(ctx, diags, res.Owner)
	metric_source.SourceType = StringToNilableValue(res.SourceType)
	metric_source.Sql = StringToStringValue(res.Sql)
	metric_source.TableName = StringToNilableValue(res.TableName)
	metric_source.Tags = StringSliceToListValue(ctx, diags, res.Tags)
	metric_source.TimestampAsDay = NilableBoolToBoolValue(res.TimestampAsDay)
	metric_source.TimestampColumn = StringToNilableValue(res.TimestampColumn)
}

type CustomFieldMappingAPIModel struct {
	Formula string `json:"formula"`
	Key     string `json:"key"`
}

func CustomFieldMappingToAPIModel(ctx context.Context, customFieldMapping *resource_metric_source.CustomFieldMappingValue) CustomFieldMappingAPIModel {
	return CustomFieldMappingAPIModel{
		Formula: StringFromNilableValue(customFieldMapping.Formula),
		Key:     StringFromNilableValue(customFieldMapping.Key),
	}
}

func CustomFieldMappingFromAPIModel(ctx context.Context, diags diag.Diagnostics, customFieldMapping *resource_metric_source.CustomFieldMappingValue, res CustomFieldMappingAPIModel) {
	customFieldMapping.Formula = StringToNilableValue(res.Formula)
	customFieldMapping.Key = StringToNilableValue(res.Key)
}

func CustomFieldMappingsToAPIModel(ctx context.Context, list basetypes.ListValue) []CustomFieldMappingAPIModel {
	var res []CustomFieldMappingAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]CustomFieldMappingAPIModel, 0)
	} else {
		res = make([]CustomFieldMappingAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric_source.CustomFieldMappingValue)
			if !ok {
				return nil
			}

			res[i] = CustomFieldMappingToAPIModel(ctx, &obj)
		}
	}
	return res
}

func CustomFieldMappingsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []CustomFieldMappingAPIModel) basetypes.ListValue {
	attrTypes := resource_metric_source.CustomFieldMappingValue{}.AttributeTypes(ctx)
	customFieldMappingType := resource_metric_source.CustomFieldMappingType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(customFieldMappingType)
	} else {
		customFieldMappings := make([]attr.Value, len(list))
		for i, elem := range list {
			var customFieldMapping resource_metric_source.CustomFieldMappingValue
			CustomFieldMappingFromAPIModel(ctx, diags, &customFieldMapping, elem)
			obj, d := customFieldMapping.ToObjectValue(ctx)
			customFieldMappings[i] = resource_metric_source.NewCustomFieldMappingValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(customFieldMappingType, customFieldMappings)
		diags = append(diags, d...)
		return v
	}
}

type MetricSourceIdTypeMappingAPIModel struct {
	Column        string `json:"column"`
	StatsigUnitId string `json:"statsigUnitID"`
}

func MetricSourceIdTypeMappingToAPIModel(ctx context.Context, idTypeMapping *resource_metric_source.IdTypeMappingValue) MetricSourceIdTypeMappingAPIModel {
	return MetricSourceIdTypeMappingAPIModel{
		Column:        StringFromNilableValue(idTypeMapping.Column),
		StatsigUnitId: StringFromNilableValue(idTypeMapping.StatsigUnitId),
	}
}

func MetricSourceIdTypeMappingFromAPIModel(ctx context.Context, diags diag.Diagnostics, idTypeMapping *resource_metric_source.IdTypeMappingValue, res MetricSourceIdTypeMappingAPIModel) {
	idTypeMapping.Column = StringToNilableValue(res.Column)
	idTypeMapping.StatsigUnitId = StringToNilableValue(res.StatsigUnitId)
}

func MetricSourceIdTypeMappingsToAPIModel(ctx context.Context, list basetypes.ListValue) []MetricSourceIdTypeMappingAPIModel {
	var res []MetricSourceIdTypeMappingAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricSourceIdTypeMappingAPIModel, 0)
	} else {
		res = make([]MetricSourceIdTypeMappingAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric_source.IdTypeMappingValue)
			if !ok {
				return nil
			}

			res[i] = MetricSourceIdTypeMappingToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MetricSourceIdTypeMappingsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricSourceIdTypeMappingAPIModel) basetypes.ListValue {
	attrTypes := resource_metric_source.IdTypeMappingValue{}.AttributeTypes(ctx)
	idTypeMappingType := resource_metric_source.IdTypeMappingType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(idTypeMappingType)
	} else {
		idTypeMappings := make([]attr.Value, len(list))
		for i, elem := range list {
			var idTypeMapping resource_metric_source.IdTypeMappingValue
			MetricSourceIdTypeMappingFromAPIModel(ctx, diags, &idTypeMapping, elem)
			obj, d := idTypeMapping.ToObjectValue(ctx)
			idTypeMappings[i] = resource_metric_source.NewIdTypeMappingValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(idTypeMappingType, idTypeMappings)
		diags = append(diags, d...)
		return v
	}
}

type MetricSourceOwnerAPIModel struct {
	OwnerEmail string `json:"ownerEmail,omitempty"`
	OwnerId    string `json:"ownerID,omitempty"`
	OwnerName  string `json:"ownerName,omitempty"`
	OwnerType  string `json:"ownerType,omitempty"`
}

func MetricSourceOwnerToAPIModel(ctx context.Context, owner resource_metric_source.OwnerValue) *MetricSourceOwnerAPIModel {
	if owner.IsNull() {
		return nil
	}
	return &MetricSourceOwnerAPIModel{
		OwnerEmail: StringFromNilableValue(owner.OwnerEmail),
		OwnerId:    StringFromNilableValue(owner.OwnerId),
		OwnerName:  StringFromNilableValue(owner.OwnerName),
		OwnerType:  StringFromNilableValue(owner.OwnerType),
	}
}

func MetricSourceOwnerFromAPIModel(ctx context.Context, diags diag.Diagnostics, owner *MetricSourceOwnerAPIModel) resource_metric_source.OwnerValue {
	if owner == nil {
		return resource_metric_source.NewOwnerValueNull()
	}

	var res resource_metric_source.OwnerValue
	res.OwnerEmail = StringToNilableValue(owner.OwnerEmail)
	res.OwnerId = StringToNilableValue(owner.OwnerId)
	res.OwnerName = StringToNilableValue(owner.OwnerName)
	res.OwnerType = StringToNilableValue(owner.OwnerType)
	obj, d := res.ToObjectValue(ctx)
	diags = append(diags, d...)
	return resource_metric_source.NewOwnerValueMust(
		resource_metric_source.OwnerValue{}.AttributeTypes(ctx),
		obj.Attributes(),
	)
}
