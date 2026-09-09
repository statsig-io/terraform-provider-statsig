package resource_metric

import (
	"context"
	"encoding/json"

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

// Every field is raw JSON so the request can leave out an attribute the
// configuration never mentioned. See utils.APIField.
type MetricAPIInputModel struct {
	CustomRollUpEnd        json.RawMessage `json:"customRollUpEnd,omitempty"`
	CustomRollUpStart      json.RawMessage `json:"customRollUpStart,omitempty"`
	Description            json.RawMessage `json:"description,omitempty"`
	Directionality         json.RawMessage `json:"directionality,omitempty"`
	DryRun                 json.RawMessage `json:"dryRun,omitempty"`
	FunnelCountDistinct    json.RawMessage `json:"funnelCountDistinct,omitempty"`
	FunnelEventList        json.RawMessage `json:"funnelEventList,omitempty"`
	Id                     json.RawMessage `json:"id,omitempty"`
	IsPermanent            json.RawMessage `json:"isPermanent,omitempty"`
	IsReadOnly             json.RawMessage `json:"isReadOnly,omitempty"`
	IsVerified             json.RawMessage `json:"isVerified,omitempty"`
	MetricComponentMetrics json.RawMessage `json:"metricComponentMetrics,omitempty"`
	MetricEvents           json.RawMessage `json:"metricEvents,omitempty"`
	Name                   json.RawMessage `json:"name,omitempty"`
	RollupTimeWindow       json.RawMessage `json:"rollupTimeWindow,omitempty"`
	Tags                   json.RawMessage `json:"tags,omitempty"`
	Team                   json.RawMessage `json:"team,omitempty"`
	TeamId                 json.RawMessage `json:"teamID,omitempty"`
	Type                   json.RawMessage `json:"type,omitempty"`
	UnitTypes              json.RawMessage `json:"unitTypes,omitempty"`
	WarehouseNative        json.RawMessage `json:"warehouseNative,omitempty"`
}

func MetricToAPIInputModel(ctx context.Context, metric *MetricModel) MetricAPIInputModel {
	return MetricAPIInputModel{
		CustomRollUpEnd:        utils.FloatAPIField(metric.CustomRollUpEnd),
		CustomRollUpStart:      utils.FloatAPIField(metric.CustomRollUpStart),
		Description:            utils.StringAPIField(metric.Description),
		Directionality:         utils.StringAPIField(metric.Directionality),
		DryRun:                 utils.BoolAPIField(metric.DryRun),
		FunnelCountDistinct:    utils.StringAPIField(metric.FunnelCountDistinct),
		FunnelEventList:        utils.APIField(metric.FunnelEventList, FunnelEventsToAPIInputModel(ctx, metric.FunnelEventList)),
		Id:                     utils.StringAPIField(metric.Id),
		IsPermanent:            utils.BoolAPIField(metric.IsPermanent),
		IsReadOnly:             utils.BoolAPIField(metric.IsReadOnly),
		IsVerified:             utils.BoolAPIField(metric.IsVerified),
		MetricComponentMetrics: utils.APIField(metric.MetricComponentMetrics, MetricComponentMetricsToAPIInputModel(ctx, metric.MetricComponentMetrics)),
		MetricEvents:           utils.APIField(metric.MetricEvents, MetricEventsToAPIInputModel(ctx, metric.MetricEvents)),
		Name:                   utils.StringAPIField(metric.Name),
		RollupTimeWindow:       utils.StringAPIField(metric.RollupTimeWindow),
		Tags:                   utils.StringSliceAPIField(ctx, metric.Tags),
		Team:                   utils.StringAPIField(metric.Team),
		TeamId:                 utils.StringAPIField(metric.TeamId),
		Type:                   utils.StringAPIField(metric.Type),
		UnitTypes:              utils.StringSliceAPIField(ctx, metric.UnitTypes),
		WarehouseNative:        utils.APIField(metric.WarehouseNative, WarehouseNativeToAPIInputModel(ctx, metric.WarehouseNative)),
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type WarehouseNativeAPIInputModel struct {
	Aggregation                         json.RawMessage `json:"aggregation,omitempty"`
	AllowNullRatioDenominator           json.RawMessage `json:"allowNullRatioDenominator,omitempty"`
	Cap                                 json.RawMessage `json:"cap,omitempty"`
	Criteria                            json.RawMessage `json:"criteria,omitempty"`
	CupedAttributionWindow              json.RawMessage `json:"cupedAttributionWindow,omitempty"`
	CustomRollUpEnd                     json.RawMessage `json:"customRollUpEnd,omitempty"`
	CustomRollUpStart                   json.RawMessage `json:"customRollUpStart,omitempty"`
	DenominatorAggregation              json.RawMessage `json:"denominatorAggregation,omitempty"`
	DenominatorCriteria                 json.RawMessage `json:"denominatorCriteria,omitempty"`
	DenominatorCustomRollupEnd          json.RawMessage `json:"denominatorCustomRollupEnd,omitempty"`
	DenominatorCustomRollupStart        json.RawMessage `json:"denominatorCustomRollupStart,omitempty"`
	DenominatorMetricSourceName         json.RawMessage `json:"denominatorMetricSourceName,omitempty"`
	DenominatorRollupTimeWindow         json.RawMessage `json:"denominatorRollupTimeWindow,omitempty"`
	DenominatorValueColumn              json.RawMessage `json:"denominatorValueColumn,omitempty"`
	FunnelCalculationWindow             json.RawMessage `json:"funnelCalculationWindow,omitempty"`
	FunnelCountDistinct                 json.RawMessage `json:"funnelCountDistinct,omitempty"`
	FunnelEvents                        json.RawMessage `json:"funnelEvents,omitempty"`
	FunnelStartCriteria                 json.RawMessage `json:"funnelStartCriteria,omitempty"`
	MetricBakeDays                      json.RawMessage `json:"metricBakeDays,omitempty"`
	MetricDimensionColumns              json.RawMessage `json:"metricDimensionColumns,omitempty"`
	MetricSourceName                    json.RawMessage `json:"metricSourceName,omitempty"`
	NumeratorAggregation                json.RawMessage `json:"numeratorAggregation,omitempty"`
	OnlyIncludeUsersWithConversionEvent json.RawMessage `json:"onlyIncludeUsersWithConversionEvent,omitempty"`
	Percentile                          json.RawMessage `json:"percentile,omitempty"`
	RollupTimeWindow                    json.RawMessage `json:"rollupTimeWindow,omitempty"`
	ValueColumn                         json.RawMessage `json:"valueColumn,omitempty"`
	ValueThreshold                      json.RawMessage `json:"valueThreshold,omitempty"`
	WaitForCohortWindow                 json.RawMessage `json:"waitForCohortWindow,omitempty"`
	WinsorizationHigh                   json.RawMessage `json:"winsorizationHigh,omitempty"`
	WinsorizationLow                    json.RawMessage `json:"winsorizationLow,omitempty"`
}

func WarehouseNativeToAPIInputModel(ctx context.Context, warehouseNative WarehouseNativeValue) *WarehouseNativeAPIInputModel {
	if warehouseNative.IsNull() {
		return nil
	}
	return &WarehouseNativeAPIInputModel{
		Aggregation:                         utils.StringAPIField(warehouseNative.Aggregation),
		AllowNullRatioDenominator:           utils.BoolAPIField(warehouseNative.AllowNullRatioDenominator),
		Cap:                                 utils.FloatAPIField(warehouseNative.Cap),
		Criteria:                            utils.APIField(warehouseNative.Criteria, CriteriasToAPIInputModel(ctx, warehouseNative.Criteria)),
		CupedAttributionWindow:              utils.FloatAPIField(warehouseNative.CupedAttributionWindow),
		CustomRollUpEnd:                     utils.FloatAPIField(warehouseNative.CustomRollUpEnd),
		CustomRollUpStart:                   utils.FloatAPIField(warehouseNative.CustomRollUpStart),
		DenominatorAggregation:              utils.StringAPIField(warehouseNative.DenominatorAggregation),
		DenominatorCriteria:                 utils.APIField(warehouseNative.DenominatorCriteria, CriteriasToAPIInputModel(ctx, warehouseNative.DenominatorCriteria)),
		DenominatorCustomRollupEnd:          utils.FloatAPIField(warehouseNative.DenominatorCustomRollupEnd),
		DenominatorCustomRollupStart:        utils.FloatAPIField(warehouseNative.DenominatorCustomRollupStart),
		DenominatorMetricSourceName:         utils.StringAPIField(warehouseNative.DenominatorMetricSourceName),
		DenominatorRollupTimeWindow:         utils.StringAPIField(warehouseNative.DenominatorRollupTimeWindow),
		DenominatorValueColumn:              utils.StringAPIField(warehouseNative.DenominatorValueColumn),
		FunnelCalculationWindow:             utils.FloatAPIField(warehouseNative.FunnelCalculationWindow),
		FunnelCountDistinct:                 utils.StringAPIField(warehouseNative.FunnelCountDistinct),
		FunnelEvents:                        utils.APIField(warehouseNative.FunnelEvents, WarehouseNativeFunnelEventsToAPIInputModel(ctx, warehouseNative.FunnelEvents)),
		FunnelStartCriteria:                 utils.StringAPIField(warehouseNative.FunnelStartCriteria),
		MetricBakeDays:                      utils.FloatAPIField(warehouseNative.MetricBakeDays),
		MetricDimensionColumns:              utils.StringSliceAPIField(ctx, warehouseNative.MetricDimensionColumns),
		MetricSourceName:                    utils.StringAPIField(warehouseNative.MetricSourceName),
		NumeratorAggregation:                utils.StringAPIField(warehouseNative.NumeratorAggregation),
		OnlyIncludeUsersWithConversionEvent: utils.BoolAPIField(warehouseNative.OnlyIncludeUsersWithConversionEvent),
		Percentile:                          utils.FloatAPIField(warehouseNative.Percentile),
		RollupTimeWindow:                    utils.StringAPIField(warehouseNative.RollupTimeWindow),
		ValueColumn:                         utils.StringAPIField(warehouseNative.ValueColumn),
		ValueThreshold:                      utils.FloatAPIField(warehouseNative.ValueThreshold),
		WaitForCohortWindow:                 utils.BoolAPIField(warehouseNative.WaitForCohortWindow),
		WinsorizationHigh:                   utils.FloatAPIField(warehouseNative.WinsorizationHigh),
		WinsorizationLow:                    utils.FloatAPIField(warehouseNative.WinsorizationLow),
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type MetricEventAPIInputModel struct {
	Criteria    json.RawMessage `json:"criteria,omitempty"`
	MetadataKey json.RawMessage `json:"metadataKey,omitempty"`
	Name        json.RawMessage `json:"name,omitempty"`
	Type        json.RawMessage `json:"type,omitempty"`
}

func MetricEventToAPIInputModel(ctx context.Context, metricEvent *MetricEventsValue) MetricEventAPIInputModel {
	return MetricEventAPIInputModel{
		Criteria:    utils.APIField(metricEvent.Criteria, CriteriasToAPIInputModel(ctx, metricEvent.Criteria)),
		MetadataKey: utils.StringAPIField(metricEvent.MetadataKey),
		Name:        utils.StringAPIField(metricEvent.Name),
		Type:        utils.StringAPIField(metricEvent.MetricEventsType),
	}
}

func MetricEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, metricEvents *MetricEventsValue, res MetricEventAPIModel) {
	metricEvents.Criteria = CriteriasFromAPIModel(ctx, diags, res.Criteria)
	metricEvents.MetadataKey = utils.StringToNilableValue(res.MetadataKey)
	metricEvents.Name = utils.StringToNilableValue(res.Name)
	metricEvents.MetricEventsType = utils.StringToNilableValue(res.Type)
}

func MetricEventsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []MetricEventAPIInputModel {
	var res []MetricEventAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricEventAPIInputModel, 0)
	} else {
		res = make([]MetricEventAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(MetricEventsValue)
			if !ok {
				return nil
			}

			res[i] = MetricEventToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type CriteriaAPIInputModel struct {
	Column              json.RawMessage `json:"column,omitempty"`
	Condition           json.RawMessage `json:"condition,omitempty"`
	NullVacuousOverride json.RawMessage `json:"nullVacuousOverride,omitempty"`
	Type                json.RawMessage `json:"type,omitempty"`
	Values              json.RawMessage `json:"values,omitempty"`
}

func CriteriaToAPIInputModel(ctx context.Context, criteria *CriteriaValue) CriteriaAPIInputModel {
	return CriteriaAPIInputModel{
		Column:              utils.StringAPIField(criteria.Column),
		Condition:           utils.StringAPIField(criteria.Condition),
		NullVacuousOverride: utils.BoolAPIField(criteria.NullVacuousOverride),
		Type:                utils.StringAPIField(criteria.CriteriaType),
		Values:              utils.StringSliceAPIField(ctx, criteria.Values),
	}
}

func CriteriaFromAPIModel(ctx context.Context, diags diag.Diagnostics, criteria *CriteriaValue, res CriteriaAPIModel) {
	criteria.Column = utils.StringToNilableValue(res.Column)
	criteria.Condition = utils.StringToNilableValue(res.Condition)
	criteria.NullVacuousOverride = utils.NilableBoolToBoolValue(res.NullVacuousOverride)
	criteria.CriteriaType = utils.StringToNilableValue(res.Type)
	criteria.Values = utils.StringSliceToListValue(ctx, diags, res.Values)
}

func CriteriasToAPIInputModel(ctx context.Context, list basetypes.ListValue) []CriteriaAPIInputModel {
	var res []CriteriaAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]CriteriaAPIInputModel, 0)
	} else {
		res = make([]CriteriaAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(CriteriaValue)
			if !ok {
				return nil
			}

			res[i] = CriteriaToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type MetricComponentMetricAPIInputModel struct {
	Name json.RawMessage `json:"name,omitempty"`
	Type json.RawMessage `json:"type,omitempty"`
}

func MetricComponentMetricToAPIInputModel(ctx context.Context, metricComponentMetrics *MetricComponentMetricsValue) MetricComponentMetricAPIInputModel {
	return MetricComponentMetricAPIInputModel{
		Name: utils.StringAPIField(metricComponentMetrics.Name),
		Type: utils.StringAPIField(metricComponentMetrics.MetricComponentMetricsType),
	}
}

func MetricComponentMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metricComponentMetrics *MetricComponentMetricsValue, res MetricComponentMetricAPIModel) {
	metricComponentMetrics.Name = utils.StringToNilableValue(res.Name)
	metricComponentMetrics.MetricComponentMetricsType = utils.StringToNilableValue(res.Type)
}

func MetricComponentMetricsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []MetricComponentMetricAPIInputModel {
	var res []MetricComponentMetricAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricComponentMetricAPIInputModel, 0)
	} else {
		res = make([]MetricComponentMetricAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(MetricComponentMetricsValue)
			if !ok {
				return nil
			}

			res[i] = MetricComponentMetricToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type FunnelEventAPIInputModel struct {
	Name json.RawMessage `json:"name,omitempty"`
	Type json.RawMessage `json:"type,omitempty"`
}

func FunnelEventToAPIInputModel(ctx context.Context, event *FunnelEventListValue) FunnelEventAPIInputModel {
	return FunnelEventAPIInputModel{
		Name: utils.StringAPIField(event.Name),
		Type: utils.StringAPIField(event.FunnelEventListType),
	}
}

func FunnelEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, event *FunnelEventListValue, res FunnelEventAPIModel) {
	event.Name = utils.StringToNilableValue(res.Name)
	event.FunnelEventListType = utils.StringToNilableValue(res.Type)
}

func FunnelEventsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []FunnelEventAPIInputModel {
	var res []FunnelEventAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]FunnelEventAPIInputModel, 0)
	} else {
		res = make([]FunnelEventAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(FunnelEventListValue)
			if !ok {
				return nil
			}

			res[i] = FunnelEventToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type WarehouseNativeFunnelEventAPIInputModel struct {
	Criteria               json.RawMessage `json:"criteria,omitempty"`
	MetricSourceName       json.RawMessage `json:"metricSourceName,omitempty"`
	Name                   json.RawMessage `json:"name,omitempty"`
	SessionIdentifierField json.RawMessage `json:"sessionIdentifierField,omitempty"`
}

func WarehouseNativeFunnelEventToAPIInputModel(ctx context.Context, funnelEvent *FunnelEventsValue) WarehouseNativeFunnelEventAPIInputModel {
	return WarehouseNativeFunnelEventAPIInputModel{
		Criteria:               utils.APIField(funnelEvent.Criteria, CriteriasToAPIInputModel(ctx, funnelEvent.Criteria)),
		MetricSourceName:       utils.StringAPIField(funnelEvent.MetricSourceName),
		Name:                   utils.StringAPIField(funnelEvent.Name),
		SessionIdentifierField: utils.StringAPIField(funnelEvent.SessionIdentifierField),
	}
}

func WarehouseNativeFunnelEventFromAPIModel(ctx context.Context, diags diag.Diagnostics, event *FunnelEventsValue, res WarehouseNativeFunnelEventAPIModel) {
	event.Criteria = CriteriasFromAPIModel(ctx, diags, res.Criteria)
	event.MetricSourceName = utils.StringToNilableValue(res.MetricSourceName)
	event.Name = utils.StringToNilableValue(res.Name)
	event.SessionIdentifierField = utils.StringToNilableValue(res.SessionIdentifierField)
}

func WarehouseNativeFunnelEventsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []WarehouseNativeFunnelEventAPIInputModel {
	var res []WarehouseNativeFunnelEventAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]WarehouseNativeFunnelEventAPIInputModel, 0)
	} else {
		res = make([]WarehouseNativeFunnelEventAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(FunnelEventsValue)
			if !ok {
				return nil
			}

			res[i] = WarehouseNativeFunnelEventToAPIInputModel(ctx, &obj)
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
