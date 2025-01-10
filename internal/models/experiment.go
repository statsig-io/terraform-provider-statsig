package models

import (
	"context"

	"terraform-provider-statsig/internal/resource_experiment"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for ExperimentModel
type ExperimentAPIModel struct {
	Allocation                     float64          `json:"allocation"`
	AllocationDuration             *int64           `json:"allocationDuration,omitempty"`
	AnalysisEndTime                string           `json:"analysisEndTime,omitempty"`
	AnalyticsType                  string           `json:"analyticsType,omitempty"`
	AssignmentSourceExperimentName string           `json:"assignmentSourceExperimentName,omitempty"`
	AssignmentSourceName           string           `json:"assignmentSourceName,omitempty"`
	BenjaminiHochbergPerMetric     *bool            `json:"benjaminiHochbergPerMetric,omitempty"`
	BenjaminiHochbergPerVariant    *bool            `json:"benjaminiHochbergPerVariant,omitempty"`
	BenjaminiPrimaryMetricsOnly    *bool            `json:"benjaminiPrimaryMetricsOnly,omitempty"`
	BonferroniCorrection           bool             `json:"bonferroniCorrection"`
	BonferroniCorrectionPerMetric  *bool            `json:"bonferroniCorrectionPerMetric,omitempty"`
	CohortWaitUntilEndToInclude    *bool            `json:"cohortWaitUntilEndToInclude,omitempty"`
	CohortedAnalysisDuration       *int64           `json:"cohortedAnalysisDuration,omitempty"`
	CohortedMetricsMatureAfterEnd  *bool            `json:"cohortedMetricsMatureAfterEnd,omitempty"`
	ControlGroupId                 string           `json:"controlGroupID,omitempty"`
	CreatorEmail                   string           `json:"creatorEmail,omitempty"`
	CreatorId                      string           `json:"creatorID,omitempty"`
	DefaultConfidenceInterval      string           `json:"defaultConfidenceInterval,omitempty"`
	Description                    string           `json:"description"`
	Duration                       *int64           `json:"duration,omitempty"`
	FixedAnalysisDuration          *int64           `json:"fixedAnalysisDuration,omitempty"`
	Groups                         []GroupAPIModel  `json:"groups"`
	Hypothesis                     string           `json:"hypothesis"`
	Id                             string           `json:"id,omitempty"`
	IdType                         string           `json:"idType"`
	IsAnalysisOnly                 *bool            `json:"isAnalysisOnly,omitempty"`
	LaunchedGroupId                string           `json:"launchedGroupID,omitempty"`
	LayerId                        string           `json:"layerID,omitempty"`
	Links                          []LinkAPIModel   `json:"links"`
	Name                           string           `json:"name"`
	PrimaryMetricTags              []string         `json:"primaryMetricTags"`
	PrimaryMetrics                 []MetricAPIModel `json:"primaryMetrics"`
	ScheduledReloadHour            *int64           `json:"scheduledReloadHour,omitempty"`
	ScheduledReloadType            string           `json:"scheduledReloadType,omitempty"`
	SecondaryIdtype                string           `json:"secondaryIDType,omitempty"`
	SecondaryMetricTags            []string         `json:"secondaryMetricTags"`
	SecondaryMetrics               []MetricAPIModel `json:"secondaryMetrics"`
	SequentialTesting              *bool            `json:"sequentialTesting,omitempty"`
	Status                         string           `json:"status,omitempty"`
	Tags                           []string         `json:"tags,omitempty"`
	TargetApps                     []string         `json:"targetApps,omitempty"`
	TargetExposures                *int64           `json:"targetExposures,omitempty"`
	TargetingGateId                string           `json:"targetingGateID,omitempty"`
	Team                           string           `json:"team,omitempty"`
}

func ExperimentToAPIModel(ctx context.Context, experiment *resource_experiment.ExperimentModel) ExperimentAPIModel {
	return ExperimentAPIModel{
		Allocation:                     FloatFromFloatValue(experiment.Allocation),
		AllocationDuration:             NilableInt64FromInt64Value(experiment.AllocationDuration),
		AnalysisEndTime:                StringFromNilableValue(experiment.AnalysisEndTime),
		AnalyticsType:                  StringFromNilableValue(experiment.AnalyticsType),
		AssignmentSourceExperimentName: StringFromNilableValue(experiment.AssignmentSourceExperimentName),
		AssignmentSourceName:           StringFromNilableValue(experiment.AssignmentSourceName),
		BenjaminiHochbergPerMetric:     NilableBoolFromBoolValue(experiment.BenjaminiHochbergPerMetric),
		BenjaminiHochbergPerVariant:    NilableBoolFromBoolValue(experiment.BenjaminiHochbergPerVariant),
		BenjaminiPrimaryMetricsOnly:    NilableBoolFromBoolValue(experiment.BenjaminiPrimaryMetricsOnly),
		BonferroniCorrection:           BoolFromBoolValue(experiment.BonferroniCorrection),
		BonferroniCorrectionPerMetric:  NilableBoolFromBoolValue(experiment.BonferroniCorrectionPerMetric),
		CohortWaitUntilEndToInclude:    NilableBoolFromBoolValue(experiment.CohortWaitUntilEndToInclude),
		CohortedAnalysisDuration:       NilableInt64FromInt64Value(experiment.CohortedAnalysisDuration),
		CohortedMetricsMatureAfterEnd:  NilableBoolFromBoolValue(experiment.CohortedMetricsMatureAfterEnd),
		ControlGroupId:                 StringFromNilableValue(experiment.ControlGroupId),
		CreatorEmail:                   StringFromNilableValue(experiment.CreatorEmail),
		CreatorId:                      StringFromNilableValue(experiment.CreatorId),
		DefaultConfidenceInterval:      StringFromNilableValue(experiment.DefaultConfidenceInterval),
		Description:                    StringFromNilableValue(experiment.Description),
		Duration:                       NilableInt64FromInt64Value(experiment.Duration),
		FixedAnalysisDuration:          NilableInt64FromInt64Value(experiment.FixedAnalysisDuration),
		Groups:                         GroupsToAPIModel(ctx, experiment.Groups),
		Hypothesis:                     StringFromNilableValue(experiment.Hypothesis),
		Id:                             StringFromNilableValue(experiment.Id),
		IdType:                         StringFromNilableValue(experiment.IdType),
		IsAnalysisOnly:                 NilableBoolFromBoolValue(experiment.IsAnalysisOnly),
		LaunchedGroupId:                StringFromNilableValue(experiment.LaunchedGroupId),
		LayerId:                        StringFromNilableValue(experiment.LayerId),
		Links:                          LinksToAPIModel(ctx, experiment.Links),
		Name:                           StringFromNilableValue(experiment.Name),
		PrimaryMetricTags:              StringSliceFromListValue(ctx, experiment.PrimaryMetricTags),
		PrimaryMetrics:                 MetricsToAPIModel(ctx, experiment.PrimaryMetrics),
		ScheduledReloadHour:            NilableInt64FromInt64Value(experiment.ScheduledReloadHour),
		ScheduledReloadType:            StringFromNilableValue(experiment.ScheduledReloadType),
		SecondaryIdtype:                StringFromNilableValue(experiment.SecondaryIdtype),
		SecondaryMetricTags:            StringSliceFromListValue(ctx, experiment.SecondaryMetricTags),
		SecondaryMetrics:               MetricsToAPIModel(ctx, experiment.SecondaryMetrics),
		SequentialTesting:              NilableBoolFromBoolValue(experiment.SequentialTesting),
		Status:                         StringFromNilableValue(experiment.Status),
		Tags:                           StringSliceFromListValue(ctx, experiment.Tags),
		TargetApps:                     StringSliceFromListValue(ctx, experiment.TargetApps),
		TargetExposures:                NilableInt64FromInt64Value(experiment.TargetExposures),
		TargetingGateId:                StringFromNilableValue(experiment.TargetingGateId),
		Team:                           StringFromNilableValue(experiment.Team),
	}
}

func ExperimentFromAPIModel(ctx context.Context, diags diag.Diagnostics, experiment *resource_experiment.ExperimentModel, res ExperimentAPIModel) {
	experiment.Allocation = FloatToFloatValue(res.Allocation)
	experiment.AllocationDuration = NilableInt64ToInt64Value(res.AllocationDuration)
	experiment.AnalysisEndTime = StringToNilableValue(res.AnalysisEndTime)
	experiment.AnalyticsType = StringToNilableValue(res.AnalyticsType)
	experiment.AssignmentSourceExperimentName = StringToNilableValue(res.AssignmentSourceExperimentName)
	experiment.AssignmentSourceName = StringToNilableValue(res.AssignmentSourceName)
	experiment.BenjaminiHochbergPerMetric = NilableBoolToBoolValue(res.BenjaminiHochbergPerMetric)
	experiment.BenjaminiHochbergPerVariant = NilableBoolToBoolValue(res.BenjaminiHochbergPerVariant)
	experiment.BenjaminiPrimaryMetricsOnly = NilableBoolToBoolValue(res.BenjaminiPrimaryMetricsOnly)
	experiment.BonferroniCorrection = BoolToBoolValue(res.BonferroniCorrection)
	experiment.BonferroniCorrectionPerMetric = NilableBoolToBoolValue(res.BonferroniCorrectionPerMetric)
	experiment.CohortWaitUntilEndToInclude = NilableBoolToBoolValue(res.CohortWaitUntilEndToInclude)
	experiment.CohortedAnalysisDuration = NilableInt64ToInt64Value(res.CohortedAnalysisDuration)
	experiment.CohortedMetricsMatureAfterEnd = NilableBoolToBoolValue(res.CohortedMetricsMatureAfterEnd)
	experiment.ControlGroupId = StringToNilableValue(res.ControlGroupId)
	experiment.CreatorEmail = StringToNilableValue(res.CreatorEmail)
	experiment.CreatorId = StringToNilableValue(res.CreatorId)
	experiment.DefaultConfidenceInterval = StringToNilableValue(res.DefaultConfidenceInterval)
	experiment.Description = StringToNilableValue(res.Description)
	experiment.Duration = NilableInt64ToInt64Value(res.Duration)
	experiment.FixedAnalysisDuration = NilableInt64ToInt64Value(res.FixedAnalysisDuration)
	experiment.Groups = GroupsFromAPIModel(ctx, diags, res.Groups)
	experiment.Hypothesis = StringToNilableValue(res.Hypothesis)
	experiment.Id = StringToNilableValue(res.Id)
	experiment.IdType = StringToNilableValue(res.IdType)
	experiment.IsAnalysisOnly = NilableBoolToBoolValue(res.IsAnalysisOnly)
	experiment.LaunchedGroupId = StringToNilableValue(res.LaunchedGroupId)
	experiment.LayerId = StringToNilableValue(res.LayerId)
	experiment.Links = LinksFromAPIModel(ctx, diags, res.Links)
	experiment.Name = StringToNilableValue(res.Name)
	experiment.PrimaryMetricTags = StringSliceToListValue(ctx, diags, res.PrimaryMetricTags)
	experiment.PrimaryMetrics = PrimaryMetricsFromAPIModel(ctx, diags, res.PrimaryMetrics)
	experiment.ScheduledReloadHour = NilableInt64ToInt64Value(res.ScheduledReloadHour)
	experiment.ScheduledReloadType = StringToNilableValue(res.ScheduledReloadType)
	experiment.SecondaryIdtype = StringToNilableValue(res.SecondaryIdtype)
	experiment.SecondaryMetricTags = StringSliceToListValue(ctx, diags, res.SecondaryMetricTags)
	experiment.SecondaryMetrics = SecondaryMetricsFromAPIModel(ctx, diags, res.SecondaryMetrics)
	experiment.SequentialTesting = NilableBoolToBoolValue(res.SequentialTesting)
	experiment.Status = StringToNilableValue(res.Status)
	experiment.Tags = StringSliceToListValue(ctx, diags, res.Tags)
	experiment.TargetApps = StringSliceToListValue(ctx, diags, res.TargetApps)
	experiment.TargetExposures = NilableInt64ToInt64Value(res.TargetExposures)
	experiment.TargetingGateId = StringToNilableValue(res.TargetingGateId)
	experiment.Team = StringToNilableValue(res.Team)
}

type GroupAPIModel struct {
	Name            string                 `json:"name"`
	Id              string                 `json:"id,omitempty"`
	Size            float64                `json:"size"`
	ParameterValues map[string]interface{} `json:"parameterValues"`
	Disabled        bool                   `json:"disabled,omitempty"`
	Description     string                 `json:"description,omitempty"`
	ForeignGroupId  string                 `json:"foreignGroupID,omitempty"`
}

func GroupToAPIModel(ctx context.Context, group *resource_experiment.GroupsValue) GroupAPIModel {
	return GroupAPIModel{
		Name:            StringFromNilableValue(group.Name),
		Id:              StringFromNilableValue(group.Id),
		Size:            FloatFromFloatValue(group.Size),
		ParameterValues: MapFromMapValue(ctx, group.ParameterValues),
		Disabled:        BoolFromBoolValue(group.Disabled),
		Description:     StringFromNilableValue(group.Description),
		ForeignGroupId:  StringFromNilableValue(group.ForeignGroupId),
	}
}

func GroupFromAPIModel(ctx context.Context, diags diag.Diagnostics, group *resource_experiment.GroupsValue, res GroupAPIModel) {
	group.Name = StringToNilableValue(res.Name)
	group.Id = StringToNilableValue(res.Id)
	group.Size = FloatToFloatValue(res.Size)
	group.ParameterValues = MapToMapValue(ctx, diags, res.ParameterValues)
	group.Disabled = BoolToBoolValue(res.Disabled)
	group.Description = StringToNilableValue(res.Description)
	group.ForeignGroupId = StringToNilableValue(res.ForeignGroupId)
}

func GroupsToAPIModel(ctx context.Context, list basetypes.ListValue) []GroupAPIModel {
	var res []GroupAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]GroupAPIModel, 0)
	} else {
		res = make([]GroupAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_experiment.GroupsValue)
			if !ok {
				return nil
			}

			res[i] = GroupToAPIModel(ctx, &obj)
		}
	}
	return res
}

func GroupsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []GroupAPIModel) basetypes.ListValue {
	attrTypes := resource_experiment.GroupsValue{}.AttributeTypes(ctx)
	groupsType := resource_experiment.GroupsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(groupsType)
	} else {
		groups := make([]attr.Value, len(list))
		for i, elem := range list {
			var group resource_experiment.GroupsValue
			GroupFromAPIModel(ctx, diags, &group, elem)
			obj, d := group.ToObjectValue(ctx)
			groups[i] = resource_experiment.NewGroupsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(groupsType, groups)
		diags = append(diags, d...)
		return v
	}
}

type MetricAPIModel struct {
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	Direction         string  `json:"direction,omitempty"`
	HypothesizedValue float64 `json:"hypothesizedValue,omitempty"`
}

func PrimaryMetricToAPIModel(ctx context.Context, metric *resource_experiment.PrimaryMetricsValue) MetricAPIModel {
	return MetricAPIModel{
		Name:              StringFromNilableValue(metric.Name),
		Type:              StringFromNilableValue(metric.PrimaryMetricsType),
		Direction:         StringFromNilableValue(metric.Direction),
		HypothesizedValue: FloatFromFloatValue(metric.HypothesizedValue),
	}
}

func SecondaryMetricToAPIModel(ctx context.Context, metric *resource_experiment.SecondaryMetricsValue) MetricAPIModel {
	return MetricAPIModel{
		Name:              StringFromNilableValue(metric.Name),
		Type:              StringFromNilableValue(metric.SecondaryMetricsType),
		Direction:         StringFromNilableValue(metric.Direction),
		HypothesizedValue: FloatFromFloatValue(metric.HypothesizedValue),
	}
}

func PrimaryMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *resource_experiment.PrimaryMetricsValue, res MetricAPIModel) {
	metric.Name = StringToNilableValue(res.Name)
	metric.PrimaryMetricsType = StringToNilableValue(res.Type)
	metric.Direction = StringToNilableValue(res.Direction)
	metric.HypothesizedValue = FloatToFloatValue(res.HypothesizedValue)
}

func SecondaryMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *resource_experiment.SecondaryMetricsValue, res MetricAPIModel) {
	metric.Name = StringToNilableValue(res.Name)
	metric.SecondaryMetricsType = StringToNilableValue(res.Type)
	metric.Direction = StringToNilableValue(res.Direction)
	metric.HypothesizedValue = FloatToFloatValue(res.HypothesizedValue)
}

func MetricsToAPIModel(ctx context.Context, list basetypes.ListValue) []MetricAPIModel {
	var res []MetricAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MetricAPIModel, 0)
	} else {
		res = make([]MetricAPIModel, len(list.Elements()))
		list.ElementsAs(ctx, &res, false)

		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_experiment.PrimaryMetricsValue)
			if ok {
				res[i] = PrimaryMetricToAPIModel(ctx, &obj)
			} else {
				obj, ok := elem.(resource_experiment.SecondaryMetricsValue)
				if ok {
					res[i] = SecondaryMetricToAPIModel(ctx, &obj)
				}
			}
		}
	}
	return res
}

func PrimaryMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricAPIModel) basetypes.ListValue {
	attrTypes := resource_experiment.PrimaryMetricsValue{}.AttributeTypes(ctx)
	metricsType := resource_experiment.PrimaryMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricsType)
	} else {
		metrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metric resource_experiment.PrimaryMetricsValue
			PrimaryMetricFromAPIModel(ctx, diags, &metric, elem)
			obj, d := metric.ToObjectValue(ctx)
			metrics[i] = resource_experiment.NewPrimaryMetricsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(metricsType, metrics)
		diags = append(diags, d...)
		return v
	}
}

func SecondaryMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MetricAPIModel) basetypes.ListValue {
	attrTypes := resource_experiment.SecondaryMetricsValue{}.AttributeTypes(ctx)
	metricsType := resource_experiment.SecondaryMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricsType)
	} else {
		metrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metric resource_experiment.SecondaryMetricsValue
			SecondaryMetricFromAPIModel(ctx, diags, &metric, elem)
			obj, d := metric.ToObjectValue(ctx)
			metrics[i] = resource_experiment.NewSecondaryMetricsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(metricsType, metrics)
		diags = append(diags, d...)
		return v
	}
}

type LinkAPIModel struct {
	Url   string `json:"url"`
	Title string `json:"title,omitempty"`
}

func LinkToAPIModel(ctx context.Context, link *resource_experiment.LinksValue) LinkAPIModel {
	return LinkAPIModel{
		Url:   StringFromNilableValue(link.Url),
		Title: StringFromNilableValue(link.Title),
	}
}

func LinkFromAPIModel(ctx context.Context, diags diag.Diagnostics, link *resource_experiment.LinksValue, res LinkAPIModel) {
	link.Url = StringToNilableValue(res.Url)
	link.Title = StringToNilableValue(res.Title)
}

func LinksToAPIModel(ctx context.Context, list basetypes.ListValue) []LinkAPIModel {
	var res []LinkAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]LinkAPIModel, 0)
	} else {
		res = make([]LinkAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_experiment.LinksValue)
			if !ok {
				return nil
			}

			res[i] = LinkToAPIModel(ctx, &obj)
		}
	}
	return res
}

func LinksFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []LinkAPIModel) basetypes.ListValue {
	attrTypes := resource_experiment.LinksValue{}.AttributeTypes(ctx)
	linksType := resource_experiment.LinksType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(linksType)
	} else {
		links := make([]attr.Value, len(list))
		for i, elem := range list {
			var link resource_experiment.LinksValue
			LinkFromAPIModel(ctx, diags, &link, elem)
			obj, d := link.ToObjectValue(ctx)
			links[i] = resource_experiment.NewLinksValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(linksType, links)
		diags = append(diags, d...)
		return v
	}
}
