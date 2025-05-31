package resource_metric

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
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

func MetricToAPIModel(ctx context.Context, metric *MetricModel) MetricAPIModel {
	return MetricAPIModel{
		CustomRollUpEnd:        utils.NilableFloatFromFloatValue(metric.CustomRollUpEnd),
		CustomRollUpStart:      utils.NilableFloatFromFloatValue(metric.CustomRollUpStart),
		Description:            utils.StringFromNilableValue(metric.Description),
		Directionality:         utils.StringFromNilableValue(metric.Directionality),
		DryRun:                 utils.NilableBoolFromBoolValue(metric.DryRun),
		FunnelCountDistinct:    utils.StringFromNilableValue(metric.FunnelCountDistinct),
		FunnelEventList:        FunnelEventsToAPIModel(ctx, metric.FunnelEventList),
		Id:                     utils.StringFromNilableValue(metric.Id),
		IsPermanent:            utils.NilableBoolFromBoolValue(metric.IsPermanent),
		IsReadOnly:             utils.NilableBoolFromBoolValue(metric.IsReadOnly),
		IsVerified:             utils.NilableBoolFromBoolValue(metric.IsVerified),
		MetricComponentMetrics: MetricComponentMetricsToAPIModel(ctx, metric.MetricComponentMetrics),
		MetricEvents:           MetricEventsToAPIModel(ctx, metric.MetricEvents),
		Name:                   utils.StringFromNilableValue(metric.Name),
		RollupTimeWindow:       utils.StringFromNilableValue(metric.RollupTimeWindow),
		Tags:                   utils.StringSliceFromListValue(ctx, metric.Tags),
		Team:                   utils.StringFromNilableValue(metric.Team),
		TeamId:                 utils.StringFromNilableValue(metric.TeamId),
		Type:                   utils.StringFromNilableValue(metric.Type),
		UnitTypes:              utils.StringSliceFromListValue(ctx, metric.UnitTypes),
		WarehouseNative:        WarehouseNativeToAPIModel(ctx, metric.WarehouseNative),
	}
}

func MetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *MetricModel, res MetricAPIModel) {
	metric.CustomRollUpEnd = utils.NilableFloatToFloatValue(res.CustomRollUpEnd)
	metric.CustomRollUpStart = utils.NilableFloatToFloatValue(res.CustomRollUpStart)
	metric.Description = utils.StringToNilableValue(res.Description)
	metric.Directionality = utils.StringToNilableValue(res.Directionality)
	metric.DryRun = utils.NilableBoolToBoolValue(res.DryRun)
	metric.FunnelCountDistinct = utils.StringToNilableValue(res.FunnelCountDistinct)
	metric.FunnelEventList = FunnelEventsFromAPIModel(ctx, diags, res.FunnelEventList)
	metric.Id = utils.StringToNilableValue(res.Id)
	metric.IsPermanent = utils.NilableBoolToBoolValue(res.IsPermanent)
	metric.IsReadOnly = utils.NilableBoolToBoolValue(res.IsReadOnly)
	metric.IsVerified = utils.NilableBoolToBoolValue(res.IsVerified)
	metric.MetricComponentMetrics = MetricComponentMetricsFromAPIModel(ctx, diags, res.MetricComponentMetrics)
	metric.MetricEvents = MetricEventsFromAPIModel(ctx, diags, res.MetricEvents)
	metric.Name = utils.StringToNilableValue(res.Name)
	metric.RollupTimeWindow = utils.StringToNilableValue(res.RollupTimeWindow)
	metric.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	metric.Team = utils.StringToNilableValue(res.Team)
	metric.TeamId = utils.StringToNilableValue(res.TeamId)
	metric.Type = utils.StringToNilableValue(res.Type)
	metric.UnitTypes = utils.StringSliceToListValue(ctx, diags, res.UnitTypes)
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

func WarehouseNativeToAPIModel(ctx context.Context, warehouseNative WarehouseNativeValue) *WarehouseNativeAPIModel {
	if warehouseNative.IsNull() {
		return nil
	}
	return &WarehouseNativeAPIModel{
		Aggregation:                         utils.StringFromNilableValue(warehouseNative.Aggregation),
		AllowNullRatioDenominator:           utils.NilableBoolFromBoolValue(warehouseNative.AllowNullRatioDenominator),
		Cap:                                 utils.NilableFloatFromFloatValue(warehouseNative.Cap),
		Criteria:                            CriteriasToAPIModel(ctx, warehouseNative.Criteria),
		CupedAttributionWindow:              utils.NilableFloatFromFloatValue(warehouseNative.CupedAttributionWindow),
		CustomRollUpEnd:                     utils.NilableFloatFromFloatValue(warehouseNative.CustomRollUpEnd),
		CustomRollUpStart:                   utils.NilableFloatFromFloatValue(warehouseNative.CustomRollUpStart),
		DenominatorAggregation:              utils.StringFromNilableValue(warehouseNative.DenominatorAggregation),
		DenominatorCriteria:                 CriteriasToAPIModel(ctx, warehouseNative.DenominatorCriteria),
		DenominatorCustomRollupEnd:          utils.NilableFloatFromFloatValue(warehouseNative.DenominatorCustomRollupEnd),
		DenominatorCustomRollupStart:        utils.NilableFloatFromFloatValue(warehouseNative.DenominatorCustomRollupStart),
		DenominatorMetricSourceName:         utils.StringFromNilableValue(warehouseNative.DenominatorMetricSourceName),
		DenominatorRollupTimeWindow:         utils.StringFromNilableValue(warehouseNative.DenominatorRollupTimeWindow),
		DenominatorValueColumn:              utils.StringFromNilableValue(warehouseNative.DenominatorValueColumn),
		FunnelCalculationWindow:             utils.NilableFloatFromFloatValue(warehouseNative.FunnelCalculationWindow),
		FunnelCountDistinct:                 utils.StringFromNilableValue(warehouseNative.FunnelCountDistinct),
		FunnelEvents:                        WarehouseNativeFunnelEventsToAPIModel(ctx, warehouseNative.FunnelEvents),
		FunnelStartCriteria:                 utils.StringFromNilableValue(warehouseNative.FunnelStartCriteria),
		MetricBakeDays:                      utils.NilableFloatFromFloatValue(warehouseNative.MetricBakeDays),
		MetricDimensionColumns:              utils.StringSliceFromListValue(ctx, warehouseNative.MetricDimensionColumns),
		MetricSourceName:                    utils.StringFromNilableValue(warehouseNative.MetricSourceName),
		NumeratorAggregation:                utils.StringFromNilableValue(warehouseNative.NumeratorAggregation),
		OnlyIncludeUsersWithConversionEvent: utils.NilableBoolFromBoolValue(warehouseNative.OnlyIncludeUsersWithConversionEvent),
		Percentile:                          utils.NilableFloatFromFloatValue(warehouseNative.Percentile),
		RollupTimeWindow:                    utils.StringFromNilableValue(warehouseNative.RollupTimeWindow),
		ValueColumn:                         utils.StringFromNilableValue(warehouseNative.ValueColumn),
		ValueThreshold:                      utils.NilableFloatFromFloatValue(warehouseNative.ValueThreshold),
		WaitForCohortWindow:                 utils.NilableBoolFromBoolValue(warehouseNative.WaitForCohortWindow),
		WinsorizationHigh:                   utils.NilableFloatFromFloatValue(warehouseNative.WinsorizationHigh),
		WinsorizationLow:                    utils.NilableFloatFromFloatValue(warehouseNative.WinsorizationLow),
	}
}

func WarehouseNativeFromAPIModel(ctx context.Context, diags diag.Diagnostics, warehouseNative *WarehouseNativeAPIModel) WarehouseNativeValue {
	if warehouseNative == nil {
		return NewWarehouseNativeValueNull()
	}

	var res WarehouseNativeValue
	res.Aggregation = utils.StringToNilableValue(warehouseNative.Aggregation)
	res.AllowNullRatioDenominator = utils.NilableBoolToBoolValue(warehouseNative.AllowNullRatioDenominator)
	res.Cap = utils.NilableFloatToFloatValue(warehouseNative.Cap)
	res.Criteria = CriteriasFromAPIModel(ctx, diags, warehouseNative.Criteria)
	res.CupedAttributionWindow = utils.NilableFloatToFloatValue(warehouseNative.CupedAttributionWindow)
	res.CustomRollUpEnd = utils.NilableFloatToFloatValue(warehouseNative.CustomRollUpEnd)
	res.CustomRollUpStart = utils.NilableFloatToFloatValue(warehouseNative.CustomRollUpStart)
	res.DenominatorAggregation = utils.StringToNilableValue(warehouseNative.DenominatorAggregation)
	res.DenominatorCriteria = CriteriasFromAPIModel(ctx, diags, warehouseNative.DenominatorCriteria)
	res.DenominatorCustomRollupEnd = utils.NilableFloatToFloatValue(warehouseNative.DenominatorCustomRollupEnd)
	res.DenominatorCustomRollupStart = utils.NilableFloatToFloatValue(warehouseNative.DenominatorCustomRollupStart)
	res.DenominatorMetricSourceName = utils.StringToNilableValue(warehouseNative.DenominatorMetricSourceName)
	res.DenominatorRollupTimeWindow = utils.StringToNilableValue(warehouseNative.DenominatorRollupTimeWindow)
	res.DenominatorValueColumn = utils.StringToNilableValue(warehouseNative.DenominatorValueColumn)
	res.FunnelCalculationWindow = utils.NilableFloatToFloatValue(warehouseNative.FunnelCalculationWindow)
	res.FunnelCountDistinct = utils.StringToNilableValue(warehouseNative.FunnelCountDistinct)
	res.FunnelEvents = WarehouseNativeFunnelEventsFromAPIModel(ctx, diags, warehouseNative.FunnelEvents)
	res.FunnelStartCriteria = utils.StringToNilableValue(warehouseNative.FunnelStartCriteria)
	res.MetricBakeDays = utils.NilableFloatToFloatValue(warehouseNative.MetricBakeDays)
	res.MetricDimensionColumns = utils.StringSliceToListValue(ctx, diags, warehouseNative.MetricDimensionColumns)
	res.MetricSourceName = utils.StringToNilableValue(warehouseNative.MetricSourceName)
	res.NumeratorAggregation = utils.StringToNilableValue(warehouseNative.NumeratorAggregation)
	res.OnlyIncludeUsersWithConversionEvent = utils.NilableBoolToBoolValue(warehouseNative.OnlyIncludeUsersWithConversionEvent)
	res.Percentile = utils.NilableFloatToFloatValue(warehouseNative.Percentile)
	res.RollupTimeWindow = utils.StringToNilableValue(warehouseNative.RollupTimeWindow)
	res.ValueColumn = utils.StringToNilableValue(warehouseNative.ValueColumn)
	res.ValueThreshold = utils.NilableFloatToFloatValue(warehouseNative.ValueThreshold)
	res.WaitForCohortWindow = utils.NilableBoolToBoolValue(warehouseNative.WaitForCohortWindow)
	res.WinsorizationHigh = utils.NilableFloatToFloatValue(warehouseNative.WinsorizationHigh)
	res.WinsorizationLow = utils.NilableFloatToFloatValue(warehouseNative.WinsorizationLow)
	obj, d := res.ToObjectValue(ctx)
	diags = append(diags, d...)
	return NewWarehouseNativeValueMust(
		WarehouseNativeValue{}.AttributeTypes(ctx),
		obj.Attributes(),
	)
}

type MetricEventAPIModel struct {
	Criteria    []CriteriaAPIModel `json:"criteria"`
	MetadataKey string             `json:"metadataKey,omitempty"`
	Name        string             `json:"name"`
	Type        string             `json:"type,omitempty"`
}

func MetricEventToAPIModel(ctx context.Context, metricEvent *MetricEventsValue) MetricEventAPIModel {
	return MetricEventAPIModel{
		Criteria:    CriteriasToAPIModel(ctx, metricEvent.Criteria),
		MetadataKey: utils.StringFromNilableValue(metricEvent.MetadataKey),
		Name:        utils.StringFromNilableValue(metricEvent.Name),
		Type:        utils.StringFromNilableValue(metricEvent.MetricEventsType),
	}
}

func MetricEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, metricEvents *MetricEventsValue, res MetricEventAPIModel) {
	metricEvents.Criteria = CriteriasFromAPIModel(ctx, diags, res.Criteria)
	metricEvents.MetadataKey = utils.StringToNilableValue(res.MetadataKey)
	metricEvents.Name = utils.StringToNilableValue(res.Name)
	metricEvents.MetricEventsType = utils.StringToNilableValue(res.Type)
}

func MetricEventsToAPIModel(ctx context.Context, list basetypes.ListValue) []MetricEventAPIModel {
	var res []MetricEventAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricEventAPIModel, 0)
	} else {
		res = make([]MetricEventAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(MetricEventsValue)
			if !ok {
				return nil
			}

			res[i] = MetricEventToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MetricEventsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricEventAPIModel) basetypes.ListValue {
	attrTypes := MetricEventsValue{}.AttributeTypes(ctx)
	metricEventsType := MetricEventsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricEventsType)
	} else {
		metricEvents := make([]attr.Value, len(list))
		for i, elem := range list {
			var metricEvent MetricEventsValue
			MetricEventFromAPIModel(ctx, diags, &metricEvent, elem)
			obj, d := metricEvent.ToObjectValue(ctx)
			metricEvents[i] = NewMetricEventsValueMust(attrTypes, obj.Attributes())
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

func CriteriaToAPIModel(ctx context.Context, criteria *CriteriaValue) CriteriaAPIModel {
	return CriteriaAPIModel{
		Column:              utils.StringFromNilableValue(criteria.Column),
		Condition:           utils.StringFromNilableValue(criteria.Condition),
		NullVacuousOverride: utils.NilableBoolFromBoolValue(criteria.NullVacuousOverride),
		Type:                utils.StringFromNilableValue(criteria.CriteriaType),
		Values:              utils.StringSliceFromListValue(ctx, criteria.Values),
	}
}

func CriteriaFromAPIModel(ctx context.Context, diags diag.Diagnostics, criteria *CriteriaValue, res CriteriaAPIModel) {
	criteria.Column = utils.StringToNilableValue(res.Column)
	criteria.Condition = utils.StringToNilableValue(res.Condition)
	criteria.NullVacuousOverride = utils.NilableBoolToBoolValue(res.NullVacuousOverride)
	criteria.CriteriaType = utils.StringToNilableValue(res.Type)
	criteria.Values = utils.StringSliceToListValue(ctx, diags, res.Values)
}

func CriteriasToAPIModel(ctx context.Context, list basetypes.ListValue) []CriteriaAPIModel {
	var res []CriteriaAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]CriteriaAPIModel, 0)
	} else {
		res = make([]CriteriaAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(CriteriaValue)
			if !ok {
				return nil
			}

			res[i] = CriteriaToAPIModel(ctx, &obj)
		}
	}
	return res
}

func CriteriasFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []CriteriaAPIModel) basetypes.ListValue {
	attrTypes := CriteriaValue{}.AttributeTypes(ctx)
	criteriaType := CriteriaType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(criteriaType)
	} else {
		criterias := make([]attr.Value, len(list))
		for i, elem := range list {
			var criteria CriteriaValue
			CriteriaFromAPIModel(ctx, diags, &criteria, elem)
			obj, d := criteria.ToObjectValue(ctx)
			criterias[i] = NewCriteriaValueMust(attrTypes, obj.Attributes())
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

func MetricComponentMetricToAPIModel(ctx context.Context, metricComponentMetrics *MetricComponentMetricsValue) MetricComponentMetricAPIModel {
	return MetricComponentMetricAPIModel{
		Name: utils.StringFromNilableValue(metricComponentMetrics.Name),
		Type: utils.StringFromNilableValue(metricComponentMetrics.MetricComponentMetricsType),
	}
}

func MetricComponentMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metricComponentMetrics *MetricComponentMetricsValue, res MetricComponentMetricAPIModel) {
	metricComponentMetrics.Name = utils.StringToNilableValue(res.Name)
	metricComponentMetrics.MetricComponentMetricsType = utils.StringToNilableValue(res.Type)
}

func MetricComponentMetricsToAPIModel(ctx context.Context, list basetypes.ListValue) []MetricComponentMetricAPIModel {
	var res []MetricComponentMetricAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricComponentMetricAPIModel, 0)
	} else {
		res = make([]MetricComponentMetricAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(MetricComponentMetricsValue)
			if !ok {
				return nil
			}

			res[i] = MetricComponentMetricToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MetricComponentMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricComponentMetricAPIModel) basetypes.ListValue {
	attrTypes := FunnelEventListValue{}.AttributeTypes(ctx)
	metricComponentMetricsType := MetricComponentMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricComponentMetricsType)
	} else {
		metricComponentMetrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metricComponentMetric MetricComponentMetricsValue
			MetricComponentMetricFromAPIModel(ctx, diags, &metricComponentMetric, elem)
			obj, d := metricComponentMetric.ToObjectValue(ctx)
			metricComponentMetrics[i] = NewMetricComponentMetricsValueMust(attrTypes, obj.Attributes())
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

func FunnelEventToAPIModel(ctx context.Context, event *FunnelEventListValue) FunnelEventAPIModel {
	return FunnelEventAPIModel{
		Name: utils.StringFromNilableValue(event.Name),
		Type: utils.StringFromNilableValue(event.FunnelEventListType),
	}
}

func FunnelEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, event *FunnelEventListValue, res FunnelEventAPIModel) {
	event.Name = utils.StringToNilableValue(res.Name)
	event.FunnelEventListType = utils.StringToNilableValue(res.Type)
}

func FunnelEventsToAPIModel(ctx context.Context, list basetypes.ListValue) []FunnelEventAPIModel {
	var res []FunnelEventAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]FunnelEventAPIModel, 0)
	} else {
		res = make([]FunnelEventAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(FunnelEventListValue)
			if !ok {
				return nil
			}

			res[i] = FunnelEventToAPIModel(ctx, &obj)
		}
	}
	return res
}

func FunnelEventsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []FunnelEventAPIModel) basetypes.ListValue {
	attrTypes := FunnelEventListValue{}.AttributeTypes(ctx)
	funnelEventsType := FunnelEventListType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(funnelEventsType)
	} else {
		events := make([]attr.Value, len(list))
		for i, elem := range list {
			var event FunnelEventListValue
			FunnelEventFromAPIModel(ctx, diags, &event, elem)
			obj, d := event.ToObjectValue(ctx)
			events[i] = NewFunnelEventListValueMust(attrTypes, obj.Attributes())
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

func WarehouseNativeFunnelEventToAPIModel(ctx context.Context, funnelEvent *FunnelEventsValue) WarehouseNativeFunnelEventAPIModel {
	return WarehouseNativeFunnelEventAPIModel{
		Criteria:               CriteriasToAPIModel(ctx, funnelEvent.Criteria),
		MetricSourceName:       utils.StringFromNilableValue(funnelEvent.MetricSourceName),
		Name:                   utils.StringFromNilableValue(funnelEvent.Name),
		SessionIdentifierField: utils.StringFromNilableValue(funnelEvent.SessionIdentifierField),
	}
}

func WarehouseNativeFunnelEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, event *FunnelEventsValue, res WarehouseNativeFunnelEventAPIModel) {
	event.Criteria = CriteriasFromAPIModel(ctx, diags, res.Criteria)
	event.MetricSourceName = utils.StringToNilableValue(res.MetricSourceName)
	event.Name = utils.StringToNilableValue(res.Name)
	event.SessionIdentifierField = utils.StringToNilableValue(res.SessionIdentifierField)
}

func WarehouseNativeFunnelEventsToAPIModel(ctx context.Context, list basetypes.ListValue) []WarehouseNativeFunnelEventAPIModel {
	var res []WarehouseNativeFunnelEventAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]WarehouseNativeFunnelEventAPIModel, 0)
	} else {
		res = make([]WarehouseNativeFunnelEventAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(FunnelEventsValue)
			if !ok {
				return nil
			}

			res[i] = WarehouseNativeFunnelEventToAPIModel(ctx, &obj)
		}
	}
	return res
}

func WarehouseNativeFunnelEventsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []WarehouseNativeFunnelEventAPIModel) basetypes.ListValue {
	attrTypes := FunnelEventsValue{}.AttributeTypes(ctx)
	warehouseNativeFunnelEventsType := FunnelEventsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(warehouseNativeFunnelEventsType)
	} else {
		events := make([]attr.Value, len(list))
		for i, elem := range list {
			var event FunnelEventsValue
			WarehouseNativeFunnelEventFromAPIModel(ctx, diags, &event, elem)
			obj, d := event.ToObjectValue(ctx)
			events[i] = NewFunnelEventsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(warehouseNativeFunnelEventsType, events)
		diags = append(diags, d...)
		return v
	}
}
