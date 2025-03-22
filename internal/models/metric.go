package models

import (
	"context"
	"terraform-provider-statsig/internal/resource_metric"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for MetricModel
type MetricAPIModel struct {
	CustomRollUpEnd        *float64                        `json:"customRollUpEnd,omitempty"`
	CustomRollUpStart      *float64                        `json:"customRollUpStart,omitempty"`
	Description            string                          `json:"description,omitempty"`
	Directionality         string                          `json:"directionality,omitempty"`
	DryRun                 *bool                           `json:"dryRun,omitempty"`
	FunnelCountDistinct    string                          `json:"funnelCountDistinct,omitempty"`
	FunnelEventList        []FunnelEventAPIModel           `json:"funnelEventList"`
	Id                     string                          `json:"id,omitempty"`
	IsPermanent            *bool                           `json:"isPermanent,omitempty"`
	IsReadOnly             *bool                           `json:"isReadOnly,omitempty"`
	IsVerified             *bool                           `json:"isVerified,omitempty"`
	MetricComponentMetrics []MetricComponentMetricAPIModel `json:"metricComponentMetrics"`
	MetricEvents           []MetricEventAPIModel           `json:"metricEvents"`
	Name                   string                          `json:"name"`
	RollupTimeWindow       string                          `json:"rollupTimeWindow,omitempty"`
	Tags                   []string                        `json:"tags"`
	Team                   string                          `json:"team,omitempty"`
	TeamId                 string                          `json:"teamID,omitempty"`
	Type                   string                          `json:"type"`
	UnitTypes              []string                        `json:"unitTypes"`
	WarehouseNative        *WarehouseNativeAPIModel        `json:"warehouseNative,omitempty"`
}

func MetricToAPIModel(ctx context.Context, metric *resource_metric.MetricModel) MetricAPIModel {
	return MetricAPIModel{
		CustomRollUpEnd:        NilableFloatFromFloatValue(metric.CustomRollUpEnd),
		CustomRollUpStart:      NilableFloatFromFloatValue(metric.CustomRollUpStart),
		Description:            StringFromNilableValue(metric.Description),
		Directionality:         StringFromNilableValue(metric.Directionality),
		DryRun:                 NilableBoolFromBoolValue(metric.DryRun),
		FunnelCountDistinct:    StringFromNilableValue(metric.FunnelCountDistinct),
		FunnelEventList:        FunnelEventsToAPIModel(ctx, metric.FunnelEventList),
		Id:                     StringFromNilableValue(metric.Id),
		IsPermanent:            NilableBoolFromBoolValue(metric.IsPermanent),
		IsReadOnly:             NilableBoolFromBoolValue(metric.IsReadOnly),
		IsVerified:             NilableBoolFromBoolValue(metric.IsVerified),
		MetricComponentMetrics: MetricComponentMetricsToAPIModel(ctx, metric.MetricComponentMetrics),
		MetricEvents:           MetricEventsToAPIModel(ctx, metric.MetricEvents),
		Name:                   StringFromNilableValue(metric.Name),
		RollupTimeWindow:       StringFromNilableValue(metric.RollupTimeWindow),
		Tags:                   StringSliceFromListValue(ctx, metric.Tags),
		Team:                   StringFromNilableValue(metric.Team),
		TeamId:                 StringFromNilableValue(metric.TeamId),
		Type:                   StringFromNilableValue(metric.Type),
		UnitTypes:              StringSliceFromListValue(ctx, metric.UnitTypes),
		WarehouseNative:        WarehouseNativeToAPIModel(ctx, metric.WarehouseNative),
	}
}

func MetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *resource_metric.MetricModel, res MetricAPIModel) {
	metric.CustomRollUpEnd = NilableFloatToFloatValue(res.CustomRollUpEnd)
	metric.CustomRollUpStart = NilableFloatToFloatValue(res.CustomRollUpStart)
	metric.Description = StringToNilableValue(res.Description)
	metric.Directionality = StringToNilableValue(res.Directionality)
	metric.DryRun = NilableBoolToBoolValue(res.DryRun)
	metric.FunnelCountDistinct = StringToNilableValue(res.FunnelCountDistinct)
	metric.FunnelEventList = FunnelEventsFromAPIModel(ctx, diags, res.FunnelEventList)
	metric.Id = StringToNilableValue(res.Id)
	metric.IsPermanent = NilableBoolToBoolValue(res.IsPermanent)
	metric.IsReadOnly = NilableBoolToBoolValue(res.IsReadOnly)
	metric.IsVerified = NilableBoolToBoolValue(res.IsVerified)
	metric.MetricComponentMetrics = MetricComponentMetricsFromAPIModel(ctx, diags, res.MetricComponentMetrics)
	metric.MetricEvents = MetricEventsFromAPIModel(ctx, diags, res.MetricEvents)
	metric.Name = StringToNilableValue(res.Name)
	metric.RollupTimeWindow = StringToNilableValue(res.RollupTimeWindow)
	metric.Tags = StringSliceToListValue(ctx, diags, res.Tags)
	metric.Team = StringToNilableValue(res.Team)
	metric.TeamId = StringToNilableValue(res.TeamId)
	metric.Type = StringToNilableValue(res.Type)
	metric.UnitTypes = StringSliceToListValue(ctx, diags, res.UnitTypes)
	metric.WarehouseNative = WarehouseNativeFromAPIModel(ctx, diags, res.WarehouseNative)
}

type WarehouseNativeAPIModel struct {
	Aggregation                         string                               `json:"aggregation,omitempty"`
	AllowNullRatioDenominator           *bool                                `json:"allowNullRatioDenominator,omitempty"`
	Cap                                 *float64                             `json:"cap,omitempty"`
	Criteria                            []CriteriaAPIModel                   `json:"criteria"`
	CupedAttributionWindow              *float64                             `json:"cupedAttributionWindow,omitempty"`
	CustomRollUpEnd                     *float64                             `json:"customRollUpEnd,omitempty"`
	CustomRollUpStart                   *float64                             `json:"customRollUpStart,omitempty"`
	DenominatorAggregation              string                               `json:"denominatorAggregation,omitempty"`
	DenominatorCriteria                 []CriteriaAPIModel                   `json:"denominatorCriteria"`
	DenominatorCustomRollupEnd          *float64                             `json:"denominatorCustomRollupEnd,omitempty"`
	DenominatorCustomRollupStart        *float64                             `json:"denominatorCustomRollupStart,omitempty"`
	DenominatorMetricSourceName         string                               `json:"denominatorMetricSourceName,omitempty"`
	DenominatorRollupTimeWindow         string                               `json:"denominatorRollupTimeWindow,omitempty"`
	DenominatorValueColumn              string                               `json:"denominatorValueColumn,omitempty"`
	FunnelCalculationWindow             *float64                             `json:"funnelCalculationWindow,omitempty"`
	FunnelCountDistinct                 string                               `json:"funnelCountDistinct,omitempty"`
	FunnelEvents                        []WarehouseNativeFunnelEventAPIModel `json:"funnelEvents,omitempty"`
	FunnelStartCriteria                 string                               `json:"funnelStartCriteria,omitempty"`
	MetricBakeDays                      *float64                             `json:"metricBakeDays,omitempty"`
	MetricDimensionColumns              []string                             `json:"metricDimensionColumns"`
	MetricSourceName                    string                               `json:"metricSourceName,omitempty"`
	NumeratorAggregation                string                               `json:"numeratorAggregation,omitempty"`
	OnlyIncludeUsersWithConversionEvent *bool                                `json:"onlyIncludeUsersWithConversionEvent,omitempty"`
	Percentile                          *float64                             `json:"percentile,omitempty"`
	RollupTimeWindow                    string                               `json:"rollupTimeWindow,omitempty"`
	ValueColumn                         string                               `json:"valueColumn,omitempty"`
	ValueThreshold                      *float64                             `json:"valueThreshold,omitempty"`
	WaitForCohortWindow                 *bool                                `json:"waitForCohortWindow,omitempty"`
	WinsorizationHigh                   *float64                             `json:"winsorizationHigh,omitempty"`
	WinsorizationLow                    *float64                             `json:"winsorizationLow,omitempty"`
}

func WarehouseNativeToAPIModel(ctx context.Context, warehouseNative resource_metric.WarehouseNativeValue) *WarehouseNativeAPIModel {
	if warehouseNative.IsNull() {
		return nil
	}
	return &WarehouseNativeAPIModel{
		Aggregation:                         StringFromNilableValue(warehouseNative.Aggregation),
		AllowNullRatioDenominator:           NilableBoolFromBoolValue(warehouseNative.AllowNullRatioDenominator),
		Cap:                                 NilableFloatFromFloatValue(warehouseNative.Cap),
		Criteria:                            CriteriasToAPIModel(ctx, warehouseNative.Criteria),
		CupedAttributionWindow:              NilableFloatFromFloatValue(warehouseNative.CupedAttributionWindow),
		CustomRollUpEnd:                     NilableFloatFromFloatValue(warehouseNative.CustomRollUpEnd),
		CustomRollUpStart:                   NilableFloatFromFloatValue(warehouseNative.CustomRollUpStart),
		DenominatorAggregation:              StringFromNilableValue(warehouseNative.DenominatorAggregation),
		DenominatorCriteria:                 CriteriasToAPIModel(ctx, warehouseNative.DenominatorCriteria),
		DenominatorCustomRollupEnd:          NilableFloatFromFloatValue(warehouseNative.DenominatorCustomRollupEnd),
		DenominatorCustomRollupStart:        NilableFloatFromFloatValue(warehouseNative.DenominatorCustomRollupStart),
		DenominatorMetricSourceName:         StringFromNilableValue(warehouseNative.DenominatorMetricSourceName),
		DenominatorRollupTimeWindow:         StringFromNilableValue(warehouseNative.DenominatorRollupTimeWindow),
		DenominatorValueColumn:              StringFromNilableValue(warehouseNative.DenominatorValueColumn),
		FunnelCalculationWindow:             NilableFloatFromFloatValue(warehouseNative.FunnelCalculationWindow),
		FunnelCountDistinct:                 StringFromNilableValue(warehouseNative.FunnelCountDistinct),
		FunnelEvents:                        WarehouseNativeFunnelEventsToAPIModel(ctx, warehouseNative.FunnelEvents),
		FunnelStartCriteria:                 StringFromNilableValue(warehouseNative.FunnelStartCriteria),
		MetricBakeDays:                      NilableFloatFromFloatValue(warehouseNative.MetricBakeDays),
		MetricDimensionColumns:              StringSliceFromListValue(ctx, warehouseNative.MetricDimensionColumns),
		MetricSourceName:                    StringFromNilableValue(warehouseNative.MetricSourceName),
		NumeratorAggregation:                StringFromNilableValue(warehouseNative.NumeratorAggregation),
		OnlyIncludeUsersWithConversionEvent: NilableBoolFromBoolValue(warehouseNative.OnlyIncludeUsersWithConversionEvent),
		Percentile:                          NilableFloatFromFloatValue(warehouseNative.Percentile),
		RollupTimeWindow:                    StringFromNilableValue(warehouseNative.RollupTimeWindow),
		ValueColumn:                         StringFromNilableValue(warehouseNative.ValueColumn),
		ValueThreshold:                      NilableFloatFromFloatValue(warehouseNative.ValueThreshold),
		WaitForCohortWindow:                 NilableBoolFromBoolValue(warehouseNative.WaitForCohortWindow),
		WinsorizationHigh:                   NilableFloatFromFloatValue(warehouseNative.WinsorizationHigh),
		WinsorizationLow:                    NilableFloatFromFloatValue(warehouseNative.WinsorizationLow),
	}
}

func WarehouseNativeFromAPIModel(ctx context.Context, diags diag.Diagnostics, warehouseNative *WarehouseNativeAPIModel) resource_metric.WarehouseNativeValue {
	if warehouseNative == nil {
		return resource_metric.NewWarehouseNativeValueNull()
	}

	var res resource_metric.WarehouseNativeValue
	res.Aggregation = StringToNilableValue(warehouseNative.Aggregation)
	res.AllowNullRatioDenominator = NilableBoolToBoolValue(warehouseNative.AllowNullRatioDenominator)
	res.Cap = NilableFloatToFloatValue(warehouseNative.Cap)
	res.Criteria = CriteriasFromAPIModel(ctx, diags, warehouseNative.Criteria)
	res.CupedAttributionWindow = NilableFloatToFloatValue(warehouseNative.CupedAttributionWindow)
	res.CustomRollUpEnd = NilableFloatToFloatValue(warehouseNative.CustomRollUpEnd)
	res.CustomRollUpStart = NilableFloatToFloatValue(warehouseNative.CustomRollUpStart)
	res.DenominatorAggregation = StringToNilableValue(warehouseNative.DenominatorAggregation)
	res.DenominatorCriteria = CriteriasFromAPIModel(ctx, diags, warehouseNative.DenominatorCriteria)
	res.DenominatorCustomRollupEnd = NilableFloatToFloatValue(warehouseNative.DenominatorCustomRollupEnd)
	res.DenominatorCustomRollupStart = NilableFloatToFloatValue(warehouseNative.DenominatorCustomRollupStart)
	res.DenominatorMetricSourceName = StringToNilableValue(warehouseNative.DenominatorMetricSourceName)
	res.DenominatorRollupTimeWindow = StringToNilableValue(warehouseNative.DenominatorRollupTimeWindow)
	res.DenominatorValueColumn = StringToNilableValue(warehouseNative.DenominatorValueColumn)
	res.FunnelCalculationWindow = NilableFloatToFloatValue(warehouseNative.FunnelCalculationWindow)
	res.FunnelCountDistinct = StringToNilableValue(warehouseNative.FunnelCountDistinct)
	res.FunnelEvents = WarehouseNativeFunnelEventsFromAPIModel(ctx, diags, warehouseNative.FunnelEvents)
	res.FunnelStartCriteria = StringToNilableValue(warehouseNative.FunnelStartCriteria)
	res.MetricBakeDays = NilableFloatToFloatValue(warehouseNative.MetricBakeDays)
	res.MetricDimensionColumns = StringSliceToListValue(ctx, diags, warehouseNative.MetricDimensionColumns)
	res.MetricSourceName = StringToNilableValue(warehouseNative.MetricSourceName)
	res.NumeratorAggregation = StringToNilableValue(warehouseNative.NumeratorAggregation)
	res.OnlyIncludeUsersWithConversionEvent = NilableBoolToBoolValue(warehouseNative.OnlyIncludeUsersWithConversionEvent)
	res.Percentile = NilableFloatToFloatValue(warehouseNative.Percentile)
	res.RollupTimeWindow = StringToNilableValue(warehouseNative.RollupTimeWindow)
	res.ValueColumn = StringToNilableValue(warehouseNative.ValueColumn)
	res.ValueThreshold = NilableFloatToFloatValue(warehouseNative.ValueThreshold)
	res.WaitForCohortWindow = NilableBoolToBoolValue(warehouseNative.WaitForCohortWindow)
	res.WinsorizationHigh = NilableFloatToFloatValue(warehouseNative.WinsorizationHigh)
	res.WinsorizationLow = NilableFloatToFloatValue(warehouseNative.WinsorizationLow)
	obj, d := res.ToObjectValue(ctx)
	diags = append(diags, d...)
	return resource_metric.NewWarehouseNativeValueMust(
		resource_metric.WarehouseNativeValue{}.AttributeTypes(ctx),
		obj.Attributes(),
	)
}

type MetricEventAPIModel struct {
	Criteria    []CriteriaAPIModel `json:"criteria"`
	MetadataKey string             `json:"metadataKey,omitempty"`
	Name        string             `json:"name"`
	Type        string             `json:"type,omitempty"`
}

func MetricEventToAPIModel(ctx context.Context, metricEvent *resource_metric.MetricEventsValue) MetricEventAPIModel {
	return MetricEventAPIModel{
		Criteria:    CriteriasToAPIModel(ctx, metricEvent.Criteria),
		MetadataKey: StringFromNilableValue(metricEvent.MetadataKey),
		Name:        StringFromNilableValue(metricEvent.Name),
		Type:        StringFromNilableValue(metricEvent.MetricEventsType),
	}
}

func MetricEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, metricEvents *resource_metric.MetricEventsValue, res MetricEventAPIModel) {
	metricEvents.Criteria = CriteriasFromAPIModel(ctx, diags, res.Criteria)
	metricEvents.MetadataKey = StringToNilableValue(res.MetadataKey)
	metricEvents.Name = StringToNilableValue(res.Name)
	metricEvents.MetricEventsType = StringToNilableValue(res.Type)
}

func MetricEventsToAPIModel(ctx context.Context, list basetypes.ListValue) []MetricEventAPIModel {
	var res []MetricEventAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricEventAPIModel, 0)
	} else {
		res = make([]MetricEventAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric.MetricEventsValue)
			if !ok {
				return nil
			}

			res[i] = MetricEventToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MetricEventsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricEventAPIModel) basetypes.ListValue {
	attrTypes := resource_metric.MetricEventsValue{}.AttributeTypes(ctx)
	metricEventsType := resource_metric.MetricEventsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricEventsType)
	} else {
		metricEvents := make([]attr.Value, len(list))
		for i, elem := range list {
			var metricEvent resource_metric.MetricEventsValue
			MetricEventFromAPIModel(ctx, diags, &metricEvent, elem)
			obj, d := metricEvent.ToObjectValue(ctx)
			metricEvents[i] = resource_metric.NewMetricEventsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(metricEventsType, metricEvents)
		diags = append(diags, d...)
		return v
	}
}

type CriteriaAPIModel struct {
	Column              string   `json:"column"`
	Condition           string   `json:"condition"`
	NullVacuousOverride *bool    `json:"nullVacuousOverride,omitempty"`
	Type                string   `json:"type"`
	Values              []string `json:"values"`
}

func CriteriaToAPIModel(ctx context.Context, criteria *resource_metric.CriteriaValue) CriteriaAPIModel {
	return CriteriaAPIModel{
		Column:              StringFromNilableValue(criteria.Column),
		Condition:           StringFromNilableValue(criteria.Condition),
		NullVacuousOverride: NilableBoolFromBoolValue(criteria.NullVacuousOverride),
		Type:                StringFromNilableValue(criteria.CriteriaType),
		Values:              StringSliceFromListValue(ctx, criteria.Values),
	}
}

func CriteriaFromAPIModel(ctx context.Context, diags diag.Diagnostics, criteria *resource_metric.CriteriaValue, res CriteriaAPIModel) {
	criteria.Column = StringToNilableValue(res.Column)
	criteria.Condition = StringToNilableValue(res.Condition)
	criteria.NullVacuousOverride = NilableBoolToBoolValue(res.NullVacuousOverride)
	criteria.CriteriaType = StringToNilableValue(res.Type)
	criteria.Values = StringSliceToListValue(ctx, diags, res.Values)
}

func CriteriasToAPIModel(ctx context.Context, list basetypes.ListValue) []CriteriaAPIModel {
	var res []CriteriaAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]CriteriaAPIModel, 0)
	} else {
		res = make([]CriteriaAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric.CriteriaValue)
			if !ok {
				return nil
			}

			res[i] = CriteriaToAPIModel(ctx, &obj)
		}
	}
	return res
}

func CriteriasFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []CriteriaAPIModel) basetypes.ListValue {
	attrTypes := resource_metric.CriteriaValue{}.AttributeTypes(ctx)
	criteriaType := resource_metric.CriteriaType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(criteriaType)
	} else {
		criterias := make([]attr.Value, len(list))
		for i, elem := range list {
			var criteria resource_metric.CriteriaValue
			CriteriaFromAPIModel(ctx, diags, &criteria, elem)
			obj, d := criteria.ToObjectValue(ctx)
			criterias[i] = resource_metric.NewCriteriaValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(criteriaType, criterias)
		diags = append(diags, d...)
		return v
	}
}

type MetricComponentMetricAPIModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func MetricComponentMetricToAPIModel(ctx context.Context, metricComponentMetrics *resource_metric.MetricComponentMetricsValue) MetricComponentMetricAPIModel {
	return MetricComponentMetricAPIModel{
		Name: StringFromNilableValue(metricComponentMetrics.Name),
		Type: StringFromNilableValue(metricComponentMetrics.MetricComponentMetricsType),
	}
}

func MetricComponentMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metricComponentMetrics *resource_metric.MetricComponentMetricsValue, res MetricComponentMetricAPIModel) {
	metricComponentMetrics.Name = StringToNilableValue(res.Name)
	metricComponentMetrics.MetricComponentMetricsType = StringToNilableValue(res.Type)
}

func MetricComponentMetricsToAPIModel(ctx context.Context, list basetypes.ListValue) []MetricComponentMetricAPIModel {
	var res []MetricComponentMetricAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricComponentMetricAPIModel, 0)
	} else {
		res = make([]MetricComponentMetricAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric.MetricComponentMetricsValue)
			if !ok {
				return nil
			}

			res[i] = MetricComponentMetricToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MetricComponentMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricComponentMetricAPIModel) basetypes.ListValue {
	attrTypes := resource_metric.FunnelEventListValue{}.AttributeTypes(ctx)
	metricComponentMetricsType := resource_metric.MetricComponentMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricComponentMetricsType)
	} else {
		metricComponentMetrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metricComponentMetric resource_metric.MetricComponentMetricsValue
			MetricComponentMetricFromAPIModel(ctx, diags, &metricComponentMetric, elem)
			obj, d := metricComponentMetric.ToObjectValue(ctx)
			metricComponentMetrics[i] = resource_metric.NewMetricComponentMetricsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(metricComponentMetricsType, metricComponentMetrics)
		diags = append(diags, d...)
		return v
	}
}

type FunnelEventAPIModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func FunnelEventToAPIModel(ctx context.Context, event *resource_metric.FunnelEventListValue) FunnelEventAPIModel {
	return FunnelEventAPIModel{
		Name: StringFromNilableValue(event.Name),
		Type: StringFromNilableValue(event.FunnelEventListType),
	}
}

func FunnelEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, event *resource_metric.FunnelEventListValue, res FunnelEventAPIModel) {
	event.Name = StringToNilableValue(res.Name)
	event.FunnelEventListType = StringToNilableValue(res.Type)
}

func FunnelEventsToAPIModel(ctx context.Context, list basetypes.ListValue) []FunnelEventAPIModel {
	var res []FunnelEventAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]FunnelEventAPIModel, 0)
	} else {
		res = make([]FunnelEventAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric.FunnelEventListValue)
			if !ok {
				return nil
			}

			res[i] = FunnelEventToAPIModel(ctx, &obj)
		}
	}
	return res
}

func FunnelEventsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []FunnelEventAPIModel) basetypes.ListValue {
	attrTypes := resource_metric.FunnelEventListValue{}.AttributeTypes(ctx)
	funnelEventsType := resource_metric.FunnelEventListType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(funnelEventsType)
	} else {
		events := make([]attr.Value, len(list))
		for i, elem := range list {
			var event resource_metric.FunnelEventListValue
			FunnelEventFromAPIModel(ctx, diags, &event, elem)
			obj, d := event.ToObjectValue(ctx)
			events[i] = resource_metric.NewFunnelEventListValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(funnelEventsType, events)
		diags = append(diags, d...)
		return v
	}
}

type WarehouseNativeFunnelEventAPIModel struct {
	Criteria               []CriteriaAPIModel `json:"criteria"`
	MetricSourceName       string             `json:"metricSourceName,omitempty"`
	Name                   string             `json:"name,omitempty"`
	SessionIdentifierField string             `json:"sessionIdentifierField,omitempty"`
}

func WarehouseNativeFunnelEventToAPIModel(ctx context.Context, funnelEvent *resource_metric.FunnelEventsValue) WarehouseNativeFunnelEventAPIModel {
	return WarehouseNativeFunnelEventAPIModel{
		Criteria:               CriteriasToAPIModel(ctx, funnelEvent.Criteria),
		MetricSourceName:       StringFromNilableValue(funnelEvent.MetricSourceName),
		Name:                   StringFromNilableValue(funnelEvent.Name),
		SessionIdentifierField: StringFromNilableValue(funnelEvent.SessionIdentifierField),
	}
}

func WarehouseNativeFunnelEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, event *resource_metric.FunnelEventsValue, res WarehouseNativeFunnelEventAPIModel) {
	event.Criteria = CriteriasFromAPIModel(ctx, diags, res.Criteria)
	event.MetricSourceName = StringToNilableValue(res.MetricSourceName)
	event.Name = StringToNilableValue(res.Name)
	event.SessionIdentifierField = StringToNilableValue(res.SessionIdentifierField)
}

func WarehouseNativeFunnelEventsToAPIModel(ctx context.Context, list basetypes.ListValue) []WarehouseNativeFunnelEventAPIModel {
	var res []WarehouseNativeFunnelEventAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]WarehouseNativeFunnelEventAPIModel, 0)
	} else {
		res = make([]WarehouseNativeFunnelEventAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_metric.FunnelEventsValue)
			if !ok {
				return nil
			}

			res[i] = WarehouseNativeFunnelEventToAPIModel(ctx, &obj)
		}
	}
	return res
}

func WarehouseNativeFunnelEventsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []WarehouseNativeFunnelEventAPIModel) basetypes.ListValue {
	attrTypes := resource_metric.FunnelEventsValue{}.AttributeTypes(ctx)
	warehouseNativeFunnelEventsType := resource_metric.FunnelEventsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(warehouseNativeFunnelEventsType)
	} else {
		events := make([]attr.Value, len(list))
		for i, elem := range list {
			var event resource_metric.FunnelEventsValue
			WarehouseNativeFunnelEventFromAPIModel(ctx, diags, &event, elem)
			obj, d := event.ToObjectValue(ctx)
			events[i] = resource_metric.NewFunnelEventsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(warehouseNativeFunnelEventsType, events)
		diags = append(diags, d...)
		return v
	}
}
