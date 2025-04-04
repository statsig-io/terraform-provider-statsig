package resource_experiment

import (
	"context"

	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for ExperimentModel
type ExperimentAPIModel struct {
	Allocation                     float64                    `json:"allocation"`
	AllocationDuration             *int64                     `json:"allocationDuration,omitempty"`
	AnalysisEndTime                string                     `json:"analysisEndTime,omitempty"`
	AnalyticsType                  string                     `json:"analyticsType,omitempty"`
	AssignmentSourceExperimentName string                     `json:"assignmentSourceExperimentName,omitempty"`
	AssignmentSourceName           string                     `json:"assignmentSourceName,omitempty"`
	BenjaminiHochbergPerMetric     *bool                      `json:"benjaminiHochbergPerMetric,omitempty"`
	BenjaminiHochbergPerVariant    *bool                      `json:"benjaminiHochbergPerVariant,omitempty"`
	BenjaminiPrimaryMetricsOnly    *bool                      `json:"benjaminiPrimaryMetricsOnly,omitempty"`
	BonferroniCorrection           bool                       `json:"bonferroniCorrection"`
	BonferroniCorrectionPerMetric  *bool                      `json:"bonferroniCorrectionPerMetric,omitempty"`
	CohortWaitUntilEndToInclude    *bool                      `json:"cohortWaitUntilEndToInclude,omitempty"`
	CohortedAnalysisDuration       *int64                     `json:"cohortedAnalysisDuration,omitempty"`
	CohortedMetricsMatureAfterEnd  *bool                      `json:"cohortedMetricsMatureAfterEnd,omitempty"`
	ControlGroupId                 string                     `json:"controlGroupID,omitempty"`
	CreatorEmail                   string                     `json:"creatorEmail,omitempty"`
	CreatorId                      string                     `json:"creatorID,omitempty"`
	DefaultConfidenceInterval      string                     `json:"defaultConfidenceInterval,omitempty"`
	Description                    string                     `json:"description"`
	Duration                       *int64                     `json:"duration,omitempty"`
	FixedAnalysisDuration          *int64                     `json:"fixedAnalysisDuration,omitempty"`
	Groups                         []GroupAPIModel            `json:"groups"`
	Hypothesis                     string                     `json:"hypothesis"`
	Id                             string                     `json:"id,omitempty"`
	IdType                         string                     `json:"idType"`
	IsAnalysisOnly                 *bool                      `json:"isAnalysisOnly,omitempty"`
	LaunchedGroupId                string                     `json:"launchedGroupID,omitempty"`
	LayerId                        string                     `json:"layerID,omitempty"`
	Links                          []LinkAPIModel             `json:"links"`
	Name                           string                     `json:"name"`
	PrimaryMetricTags              []string                   `json:"primaryMetricTags"`
	PrimaryMetrics                 []ExperimentMetricAPIModel `json:"primaryMetrics"`
	ScheduledReloadHour            *int64                     `json:"scheduledReloadHour,omitempty"`
	ScheduledReloadType            string                     `json:"scheduledReloadType,omitempty"`
	SecondaryIdtype                string                     `json:"secondaryIDType,omitempty"`
	SecondaryMetricTags            []string                   `json:"secondaryMetricTags"`
	SecondaryMetrics               []ExperimentMetricAPIModel `json:"secondaryMetrics"`
	SequentialTesting              *bool                      `json:"sequentialTesting,omitempty"`
	Status                         string                     `json:"status,omitempty"`
	Tags                           []string                   `json:"tags,omitempty"`
	TargetApps                     []string                   `json:"targetApps,omitempty"`
	TargetExposures                *int64                     `json:"targetExposures,omitempty"`
	TargetingGateId                string                     `json:"targetingGateID,omitempty"`
	Team                           string                     `json:"team,omitempty"`
}

func ExperimentToAPIModel(ctx context.Context, experiment *ExperimentModel) ExperimentAPIModel {
	return ExperimentAPIModel{
		Allocation:                     utils.FloatFromFloatValue(experiment.Allocation),
		AllocationDuration:             utils.NilableInt64FromInt64Value(experiment.AllocationDuration),
		AnalysisEndTime:                utils.StringFromNilableValue(experiment.AnalysisEndTime),
		AnalyticsType:                  utils.StringFromNilableValue(experiment.AnalyticsType),
		AssignmentSourceExperimentName: utils.StringFromNilableValue(experiment.AssignmentSourceExperimentName),
		AssignmentSourceName:           utils.StringFromNilableValue(experiment.AssignmentSourceName),
		BenjaminiHochbergPerMetric:     utils.NilableBoolFromBoolValue(experiment.BenjaminiHochbergPerMetric),
		BenjaminiHochbergPerVariant:    utils.NilableBoolFromBoolValue(experiment.BenjaminiHochbergPerVariant),
		BenjaminiPrimaryMetricsOnly:    utils.NilableBoolFromBoolValue(experiment.BenjaminiPrimaryMetricsOnly),
		BonferroniCorrection:           utils.BoolFromBoolValue(experiment.BonferroniCorrection),
		BonferroniCorrectionPerMetric:  utils.NilableBoolFromBoolValue(experiment.BonferroniCorrectionPerMetric),
		CohortWaitUntilEndToInclude:    utils.NilableBoolFromBoolValue(experiment.CohortWaitUntilEndToInclude),
		CohortedAnalysisDuration:       utils.NilableInt64FromInt64Value(experiment.CohortedAnalysisDuration),
		CohortedMetricsMatureAfterEnd:  utils.NilableBoolFromBoolValue(experiment.CohortedMetricsMatureAfterEnd),
		ControlGroupId:                 utils.StringFromNilableValue(experiment.ControlGroupId),
		CreatorEmail:                   utils.StringFromNilableValue(experiment.CreatorEmail),
		CreatorId:                      utils.StringFromNilableValue(experiment.CreatorId),
		DefaultConfidenceInterval:      utils.StringFromNilableValue(experiment.DefaultConfidenceInterval),
		Description:                    utils.StringFromNilableValue(experiment.Description),
		Duration:                       utils.NilableInt64FromInt64Value(experiment.Duration),
		FixedAnalysisDuration:          utils.NilableInt64FromInt64Value(experiment.FixedAnalysisDuration),
		Groups:                         GroupsToAPIModel(ctx, experiment.Groups),
		Hypothesis:                     utils.StringFromNilableValue(experiment.Hypothesis),
		Id:                             utils.StringFromNilableValue(experiment.Id),
		IdType:                         utils.StringFromNilableValue(experiment.IdType),
		IsAnalysisOnly:                 utils.NilableBoolFromBoolValue(experiment.IsAnalysisOnly),
		LaunchedGroupId:                utils.StringFromNilableValue(experiment.LaunchedGroupId),
		LayerId:                        utils.StringFromNilableValue(experiment.LayerId),
		Links:                          LinksToAPIModel(ctx, experiment.Links),
		Name:                           utils.StringFromNilableValue(experiment.Name),
		PrimaryMetricTags:              utils.StringSliceFromListValue(ctx, experiment.PrimaryMetricTags),
		PrimaryMetrics:                 MetricsToAPIModel(ctx, experiment.PrimaryMetrics),
		ScheduledReloadHour:            utils.NilableInt64FromInt64Value(experiment.ScheduledReloadHour),
		ScheduledReloadType:            utils.StringFromNilableValue(experiment.ScheduledReloadType),
		SecondaryIdtype:                utils.StringFromNilableValue(experiment.SecondaryIdtype),
		SecondaryMetricTags:            utils.StringSliceFromListValue(ctx, experiment.SecondaryMetricTags),
		SecondaryMetrics:               MetricsToAPIModel(ctx, experiment.SecondaryMetrics),
		SequentialTesting:              utils.NilableBoolFromBoolValue(experiment.SequentialTesting),
		Status:                         utils.StringFromNilableValue(experiment.Status),
		Tags:                           utils.StringSliceFromListValue(ctx, experiment.Tags),
		TargetApps:                     utils.StringSliceFromListValue(ctx, experiment.TargetApps),
		TargetExposures:                utils.NilableInt64FromInt64Value(experiment.TargetExposures),
		TargetingGateId:                utils.StringFromNilableValue(experiment.TargetingGateId),
		Team:                           utils.StringFromNilableValue(experiment.Team),
	}
}

func ExperimentFromAPIModel(ctx context.Context, diags diag.Diagnostics, experiment *ExperimentModel, res ExperimentAPIModel) {
	experiment.Allocation = utils.FloatToFloatValue(res.Allocation)
	experiment.AllocationDuration = utils.NilableInt64ToInt64Value(res.AllocationDuration)
	experiment.AnalysisEndTime = utils.StringToNilableValue(res.AnalysisEndTime)
	experiment.AnalyticsType = utils.StringToNilableValue(res.AnalyticsType)
	experiment.AssignmentSourceExperimentName = utils.StringToNilableValue(res.AssignmentSourceExperimentName)
	experiment.AssignmentSourceName = utils.StringToNilableValue(res.AssignmentSourceName)
	experiment.BenjaminiHochbergPerMetric = utils.NilableBoolToBoolValue(res.BenjaminiHochbergPerMetric)
	experiment.BenjaminiHochbergPerVariant = utils.NilableBoolToBoolValue(res.BenjaminiHochbergPerVariant)
	experiment.BenjaminiPrimaryMetricsOnly = utils.NilableBoolToBoolValue(res.BenjaminiPrimaryMetricsOnly)
	experiment.BonferroniCorrection = utils.BoolToBoolValue(res.BonferroniCorrection)
	experiment.BonferroniCorrectionPerMetric = utils.NilableBoolToBoolValue(res.BonferroniCorrectionPerMetric)
	experiment.CohortWaitUntilEndToInclude = utils.NilableBoolToBoolValue(res.CohortWaitUntilEndToInclude)
	experiment.CohortedAnalysisDuration = utils.NilableInt64ToInt64Value(res.CohortedAnalysisDuration)
	experiment.CohortedMetricsMatureAfterEnd = utils.NilableBoolToBoolValue(res.CohortedMetricsMatureAfterEnd)
	experiment.ControlGroupId = utils.StringToNilableValue(res.ControlGroupId)
	experiment.CreatorEmail = utils.StringToNilableValue(res.CreatorEmail)
	experiment.CreatorId = utils.StringToNilableValue(res.CreatorId)
	experiment.DefaultConfidenceInterval = utils.StringToNilableValue(res.DefaultConfidenceInterval)
	experiment.Description = utils.StringToNilableValue(res.Description)
	experiment.Duration = utils.NilableInt64ToInt64Value(res.Duration)
	experiment.FixedAnalysisDuration = utils.NilableInt64ToInt64Value(res.FixedAnalysisDuration)
	experiment.Groups = GroupsFromAPIModel(ctx, diags, res.Groups)
	experiment.Hypothesis = utils.StringToNilableValue(res.Hypothesis)
	experiment.Id = utils.StringToNilableValue(res.Id)
	experiment.IdType = utils.StringToNilableValue(res.IdType)
	experiment.IsAnalysisOnly = utils.NilableBoolToBoolValue(res.IsAnalysisOnly)
	experiment.LaunchedGroupId = utils.StringToNilableValue(res.LaunchedGroupId)
	experiment.LayerId = utils.StringToNilableValue(res.LayerId)
	experiment.Links = LinksFromAPIModel(ctx, diags, res.Links)
	experiment.Name = utils.StringToNilableValue(res.Name)
	experiment.PrimaryMetricTags = utils.StringSliceToListValue(ctx, diags, res.PrimaryMetricTags)
	experiment.PrimaryMetrics = PrimaryMetricsFromAPIModel(ctx, diags, res.PrimaryMetrics)
	experiment.ScheduledReloadHour = utils.NilableInt64ToInt64Value(res.ScheduledReloadHour)
	experiment.ScheduledReloadType = utils.StringToNilableValue(res.ScheduledReloadType)
	experiment.SecondaryIdtype = utils.StringToNilableValue(res.SecondaryIdtype)
	experiment.SecondaryMetricTags = utils.StringSliceToListValue(ctx, diags, res.SecondaryMetricTags)
	experiment.SecondaryMetrics = SecondaryMetricsFromAPIModel(ctx, diags, res.SecondaryMetrics)
	experiment.SequentialTesting = utils.NilableBoolToBoolValue(res.SequentialTesting)
	experiment.Status = utils.StringToNilableValue(res.Status)
	experiment.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	experiment.TargetApps = utils.StringSliceToListValue(ctx, diags, res.TargetApps)
	experiment.TargetExposures = utils.NilableInt64ToInt64Value(res.TargetExposures)
	experiment.TargetingGateId = utils.StringToNilableValue(res.TargetingGateId)
	experiment.Team = utils.StringToNilableValue(res.Team)
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

func GroupToAPIModel(ctx context.Context, group *GroupsValue) GroupAPIModel {
	return GroupAPIModel{
		Name:            utils.StringFromNilableValue(group.Name),
		Id:              utils.StringFromNilableValue(group.Id),
		Size:            utils.FloatFromFloatValue(group.Size),
		ParameterValues: utils.MapFromMapValue(ctx, group.ParameterValues),
		Disabled:        utils.BoolFromBoolValue(group.Disabled),
		Description:     utils.StringFromNilableValue(group.Description),
		ForeignGroupId:  utils.StringFromNilableValue(group.ForeignGroupId),
	}
}

func GroupFromAPIModel(ctx context.Context, diags diag.Diagnostics, group *GroupsValue, res GroupAPIModel) {
	group.Name = utils.StringToNilableValue(res.Name)
	group.Id = utils.StringToNilableValue(res.Id)
	group.Size = utils.FloatToFloatValue(res.Size)
	group.ParameterValues = utils.MapToMapValue(ctx, diags, res.ParameterValues)
	group.Disabled = utils.BoolToBoolValue(res.Disabled)
	group.Description = utils.StringToNilableValue(res.Description)
	group.ForeignGroupId = utils.StringToNilableValue(res.ForeignGroupId)
}

func GroupsToAPIModel(ctx context.Context, list basetypes.ListValue) []GroupAPIModel {
	var res []GroupAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]GroupAPIModel, 0)
	} else {
		res = make([]GroupAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(GroupsValue)
			if !ok {
				return nil
			}

			res[i] = GroupToAPIModel(ctx, &obj)
		}
	}
	return res
}

func GroupsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []GroupAPIModel) basetypes.ListValue {
	attrTypes := GroupsValue{}.AttributeTypes(ctx)
	groupsType := GroupsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(groupsType)
	} else {
		groups := make([]attr.Value, len(list))
		for i, elem := range list {
			var group GroupsValue
			GroupFromAPIModel(ctx, diags, &group, elem)
			obj, d := group.ToObjectValue(ctx)
			groups[i] = NewGroupsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(groupsType, groups)
		diags = append(diags, d...)
		return v
	}
}

type ExperimentMetricAPIModel struct {
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	Direction         string  `json:"direction,omitempty"`
	HypothesizedValue float64 `json:"hypothesizedValue,omitempty"`
}

func PrimaryMetricToAPIModel(ctx context.Context, metric *PrimaryMetricsValue) ExperimentMetricAPIModel {
	return ExperimentMetricAPIModel{
		Name:              utils.StringFromNilableValue(metric.Name),
		Type:              utils.StringFromNilableValue(metric.PrimaryMetricsType),
		Direction:         utils.StringFromNilableValue(metric.Direction),
		HypothesizedValue: utils.FloatFromFloatValue(metric.HypothesizedValue),
	}
}

func SecondaryMetricToAPIModel(ctx context.Context, metric *SecondaryMetricsValue) ExperimentMetricAPIModel {
	return ExperimentMetricAPIModel{
		Name:              utils.StringFromNilableValue(metric.Name),
		Type:              utils.StringFromNilableValue(metric.SecondaryMetricsType),
		Direction:         utils.StringFromNilableValue(metric.Direction),
		HypothesizedValue: utils.FloatFromFloatValue(metric.HypothesizedValue),
	}
}

func PrimaryMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *PrimaryMetricsValue, res ExperimentMetricAPIModel) {
	metric.Name = utils.StringToNilableValue(res.Name)
	metric.PrimaryMetricsType = utils.StringToNilableValue(res.Type)
	metric.Direction = utils.StringToNilableValue(res.Direction)
	metric.HypothesizedValue = utils.FloatToFloatValue(res.HypothesizedValue)
}

func SecondaryMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *SecondaryMetricsValue, res ExperimentMetricAPIModel) {
	metric.Name = utils.StringToNilableValue(res.Name)
	metric.SecondaryMetricsType = utils.StringToNilableValue(res.Type)
	metric.Direction = utils.StringToNilableValue(res.Direction)
	metric.HypothesizedValue = utils.FloatToFloatValue(res.HypothesizedValue)
}

func MetricsToAPIModel(ctx context.Context, list basetypes.ListValue) []ExperimentMetricAPIModel {
	var res []ExperimentMetricAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]ExperimentMetricAPIModel, 0)
	} else {
		res = make([]ExperimentMetricAPIModel, len(list.Elements()))
		list.ElementsAs(ctx, &res, false)

		for i, elem := range list.Elements() {
			obj, ok := elem.(PrimaryMetricsValue)
			if ok {
				res[i] = PrimaryMetricToAPIModel(ctx, &obj)
			} else {
				obj, ok := elem.(SecondaryMetricsValue)
				if ok {
					res[i] = SecondaryMetricToAPIModel(ctx, &obj)
				}
			}
		}
	}
	return res
}

func PrimaryMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []ExperimentMetricAPIModel) basetypes.ListValue {
	attrTypes := PrimaryMetricsValue{}.AttributeTypes(ctx)
	metricsType := PrimaryMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricsType)
	} else {
		metrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metric PrimaryMetricsValue
			PrimaryMetricFromAPIModel(ctx, diags, &metric, elem)
			obj, d := metric.ToObjectValue(ctx)
			metrics[i] = NewPrimaryMetricsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(metricsType, metrics)
		diags = append(diags, d...)
		return v
	}
}

func SecondaryMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []ExperimentMetricAPIModel) basetypes.ListValue {
	attrTypes := SecondaryMetricsValue{}.AttributeTypes(ctx)
	metricsType := SecondaryMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(metricsType)
	} else {
		metrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metric SecondaryMetricsValue
			SecondaryMetricFromAPIModel(ctx, diags, &metric, elem)
			obj, d := metric.ToObjectValue(ctx)
			metrics[i] = NewSecondaryMetricsValueMust(attrTypes, obj.Attributes())
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

func LinkToAPIModel(ctx context.Context, link *LinksValue) LinkAPIModel {
	return LinkAPIModel{
		Url:   utils.StringFromNilableValue(link.Url),
		Title: utils.StringFromNilableValue(link.Title),
	}
}

func LinkFromAPIModel(ctx context.Context, diags diag.Diagnostics, link *LinksValue, res LinkAPIModel) {
	link.Url = utils.StringToNilableValue(res.Url)
	link.Title = utils.StringToNilableValue(res.Title)
}

func LinksToAPIModel(ctx context.Context, list basetypes.ListValue) []LinkAPIModel {
	var res []LinkAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]LinkAPIModel, 0)
	} else {
		res = make([]LinkAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(LinksValue)
			if !ok {
				return nil
			}

			res[i] = LinkToAPIModel(ctx, &obj)
		}
	}
	return res
}

func LinksFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []LinkAPIModel) basetypes.ListValue {
	attrTypes := LinksValue{}.AttributeTypes(ctx)
	linksType := LinksType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(linksType)
	} else {
		links := make([]attr.Value, len(list))
		for i, elem := range list {
			var link LinksValue
			LinkFromAPIModel(ctx, diags, &link, elem)
			obj, d := link.ToObjectValue(ctx)
			links[i] = NewLinksValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(linksType, links)
		diags = append(diags, d...)
		return v
	}
}
