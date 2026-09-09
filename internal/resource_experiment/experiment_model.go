package resource_experiment

import (
	"context"
	"encoding/json"

	"github.com/statsig-io/terraform-provider-statsig/internal/utils"

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

// Every field is raw JSON so the request can leave out an attribute the
// configuration never mentioned. See utils.APIField.
type ExperimentAPIInputModel struct {
	Allocation                     json.RawMessage `json:"allocation,omitempty"`
	AllocationDuration             json.RawMessage `json:"allocationDuration,omitempty"`
	AnalysisEndTime                json.RawMessage `json:"analysisEndTime,omitempty"`
	AnalyticsType                  json.RawMessage `json:"analyticsType,omitempty"`
	AssignmentSourceExperimentName json.RawMessage `json:"assignmentSourceExperimentName,omitempty"`
	AssignmentSourceName           json.RawMessage `json:"assignmentSourceName,omitempty"`
	BenjaminiHochbergPerMetric     json.RawMessage `json:"benjaminiHochbergPerMetric,omitempty"`
	BenjaminiHochbergPerVariant    json.RawMessage `json:"benjaminiHochbergPerVariant,omitempty"`
	BenjaminiPrimaryMetricsOnly    json.RawMessage `json:"benjaminiPrimaryMetricsOnly,omitempty"`
	BonferroniCorrection           json.RawMessage `json:"bonferroniCorrection,omitempty"`
	BonferroniCorrectionPerMetric  json.RawMessage `json:"bonferroniCorrectionPerMetric,omitempty"`
	CohortWaitUntilEndToInclude    json.RawMessage `json:"cohortWaitUntilEndToInclude,omitempty"`
	CohortedAnalysisDuration       json.RawMessage `json:"cohortedAnalysisDuration,omitempty"`
	CohortedMetricsMatureAfterEnd  json.RawMessage `json:"cohortedMetricsMatureAfterEnd,omitempty"`
	ControlGroupId                 json.RawMessage `json:"controlGroupID,omitempty"`
	CreatorEmail                   json.RawMessage `json:"creatorEmail,omitempty"`
	CreatorId                      json.RawMessage `json:"creatorID,omitempty"`
	DefaultConfidenceInterval      json.RawMessage `json:"defaultConfidenceInterval,omitempty"`
	Description                    json.RawMessage `json:"description,omitempty"`
	Duration                       json.RawMessage `json:"duration,omitempty"`
	FixedAnalysisDuration          json.RawMessage `json:"fixedAnalysisDuration,omitempty"`
	Groups                         json.RawMessage `json:"groups,omitempty"`
	Hypothesis                     json.RawMessage `json:"hypothesis,omitempty"`
	Id                             json.RawMessage `json:"id,omitempty"`
	IdType                         json.RawMessage `json:"idType,omitempty"`
	IsAnalysisOnly                 json.RawMessage `json:"isAnalysisOnly,omitempty"`
	LaunchedGroupId                json.RawMessage `json:"launchedGroupID,omitempty"`
	LayerId                        json.RawMessage `json:"layerID,omitempty"`
	Links                          json.RawMessage `json:"links,omitempty"`
	Name                           json.RawMessage `json:"name,omitempty"`
	PrimaryMetricTags              json.RawMessage `json:"primaryMetricTags,omitempty"`
	PrimaryMetrics                 json.RawMessage `json:"primaryMetrics,omitempty"`
	ScheduledReloadHour            json.RawMessage `json:"scheduledReloadHour,omitempty"`
	ScheduledReloadType            json.RawMessage `json:"scheduledReloadType,omitempty"`
	SecondaryIdtype                json.RawMessage `json:"secondaryIDType,omitempty"`
	SecondaryMetricTags            json.RawMessage `json:"secondaryMetricTags,omitempty"`
	SecondaryMetrics               json.RawMessage `json:"secondaryMetrics,omitempty"`
	SequentialTesting              json.RawMessage `json:"sequentialTesting,omitempty"`
	Status                         json.RawMessage `json:"status,omitempty"`
	Tags                           json.RawMessage `json:"tags,omitempty"`
	TargetApps                     json.RawMessage `json:"targetApps,omitempty"`
	TargetExposures                json.RawMessage `json:"targetExposures,omitempty"`
	TargetingGateId                json.RawMessage `json:"targetingGateID,omitempty"`
	Team                           json.RawMessage `json:"team,omitempty"`
}

func ExperimentToAPIInputModel(ctx context.Context, experiment *ExperimentModel) ExperimentAPIInputModel {
	return ExperimentAPIInputModel{
		Allocation:                     utils.FloatAPIField(experiment.Allocation),
		AllocationDuration:             utils.Int64APIField(experiment.AllocationDuration),
		AnalysisEndTime:                utils.StringAPIField(experiment.AnalysisEndTime),
		AnalyticsType:                  utils.StringAPIField(experiment.AnalyticsType),
		AssignmentSourceExperimentName: utils.StringAPIField(experiment.AssignmentSourceExperimentName),
		AssignmentSourceName:           utils.StringAPIField(experiment.AssignmentSourceName),
		BenjaminiHochbergPerMetric:     utils.BoolAPIField(experiment.BenjaminiHochbergPerMetric),
		BenjaminiHochbergPerVariant:    utils.BoolAPIField(experiment.BenjaminiHochbergPerVariant),
		BenjaminiPrimaryMetricsOnly:    utils.BoolAPIField(experiment.BenjaminiPrimaryMetricsOnly),
		BonferroniCorrection:           utils.BoolAPIField(experiment.BonferroniCorrection),
		BonferroniCorrectionPerMetric:  utils.BoolAPIField(experiment.BonferroniCorrectionPerMetric),
		CohortWaitUntilEndToInclude:    utils.BoolAPIField(experiment.CohortWaitUntilEndToInclude),
		CohortedAnalysisDuration:       utils.Int64APIField(experiment.CohortedAnalysisDuration),
		CohortedMetricsMatureAfterEnd:  utils.BoolAPIField(experiment.CohortedMetricsMatureAfterEnd),
		ControlGroupId:                 utils.StringAPIField(experiment.ControlGroupId),
		CreatorEmail:                   utils.StringAPIField(experiment.CreatorEmail),
		CreatorId:                      utils.StringAPIField(experiment.CreatorId),
		DefaultConfidenceInterval:      utils.StringAPIField(experiment.DefaultConfidenceInterval),
		Description:                    utils.StringAPIField(experiment.Description),
		Duration:                       utils.Int64APIField(experiment.Duration),
		FixedAnalysisDuration:          utils.Int64APIField(experiment.FixedAnalysisDuration),
		Groups:                         utils.APIField(experiment.Groups, GroupsToAPIInputModel(ctx, experiment.Groups)),
		Hypothesis:                     utils.StringAPIField(experiment.Hypothesis),
		Id:                             utils.StringAPIField(experiment.Id),
		IdType:                         utils.StringAPIField(experiment.IdType),
		IsAnalysisOnly:                 utils.BoolAPIField(experiment.IsAnalysisOnly),
		LaunchedGroupId:                utils.StringAPIField(experiment.LaunchedGroupId),
		LayerId:                        utils.StringAPIField(experiment.LayerId),
		Links:                          utils.APIField(experiment.Links, LinksToAPIInputModel(ctx, experiment.Links)),
		Name:                           utils.StringAPIField(experiment.Name),
		PrimaryMetricTags:              utils.StringSliceAPIField(ctx, experiment.PrimaryMetricTags),
		PrimaryMetrics:                 utils.APIField(experiment.PrimaryMetrics, MetricsToAPIInputModel(ctx, experiment.PrimaryMetrics)),
		ScheduledReloadHour:            utils.Int64APIField(experiment.ScheduledReloadHour),
		ScheduledReloadType:            utils.StringAPIField(experiment.ScheduledReloadType),
		SecondaryIdtype:                utils.StringAPIField(experiment.SecondaryIdtype),
		SecondaryMetricTags:            utils.StringSliceAPIField(ctx, experiment.SecondaryMetricTags),
		SecondaryMetrics:               utils.APIField(experiment.SecondaryMetrics, MetricsToAPIInputModel(ctx, experiment.SecondaryMetrics)),
		SequentialTesting:              utils.BoolAPIField(experiment.SequentialTesting),
		Status:                         utils.StringAPIField(experiment.Status),
		Tags:                           utils.StringSliceAPIField(ctx, experiment.Tags),
		TargetApps:                     utils.StringSliceAPIField(ctx, experiment.TargetApps),
		TargetExposures:                utils.Int64APIField(experiment.TargetExposures),
		TargetingGateId:                utils.StringAPIField(experiment.TargetingGateId),
		Team:                           utils.StringAPIField(experiment.Team),
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type GroupAPIInputModel struct {
	Name            json.RawMessage `json:"name,omitempty"`
	Id              json.RawMessage `json:"id,omitempty"`
	Size            json.RawMessage `json:"size,omitempty"`
	ParameterValues json.RawMessage `json:"parameterValues,omitempty"`
	Disabled        json.RawMessage `json:"disabled,omitempty"`
	Description     json.RawMessage `json:"description,omitempty"`
	ForeignGroupId  json.RawMessage `json:"foreignGroupID,omitempty"`
}

func GroupToAPIInputModel(ctx context.Context, group *GroupsValue) GroupAPIInputModel {
	return GroupAPIInputModel{
		Name:            utils.StringAPIField(group.Name),
		Id:              utils.StringAPIField(group.Id),
		Size:            utils.FloatAPIField(group.Size),
		ParameterValues: utils.MapAPIField(ctx, group.ParameterValues),
		Disabled:        utils.BoolAPIField(group.Disabled),
		Description:     utils.StringAPIField(group.Description),
		ForeignGroupId:  utils.StringAPIField(group.ForeignGroupId),
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

func GroupsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []GroupAPIInputModel {
	var res []GroupAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]GroupAPIInputModel, 0)
	} else {
		res = make([]GroupAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(GroupsValue)
			if !ok {
				return nil
			}

			res[i] = GroupToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type ExperimentMetricAPIInputModel struct {
	Name              json.RawMessage `json:"name,omitempty"`
	Type              json.RawMessage `json:"type,omitempty"`
	Direction         json.RawMessage `json:"direction,omitempty"`
	HypothesizedValue json.RawMessage `json:"hypothesizedValue,omitempty"`
}

func PrimaryMetricToAPIInputModel(ctx context.Context, metric *PrimaryMetricsValue) ExperimentMetricAPIInputModel {
	return ExperimentMetricAPIInputModel{
		Name:              utils.StringAPIField(metric.Name),
		Type:              utils.StringAPIField(metric.PrimaryMetricsType),
		Direction:         utils.StringAPIField(metric.Direction),
		HypothesizedValue: utils.FloatAPIField(metric.HypothesizedValue),
	}
}

func SecondaryMetricToAPIInputModel(ctx context.Context, metric *SecondaryMetricsValue) ExperimentMetricAPIInputModel {
	return ExperimentMetricAPIInputModel{
		Name:              utils.StringAPIField(metric.Name),
		Type:              utils.StringAPIField(metric.SecondaryMetricsType),
		Direction:         utils.StringAPIField(metric.Direction),
		HypothesizedValue: utils.FloatAPIField(metric.HypothesizedValue),
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

func MetricsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []ExperimentMetricAPIInputModel {
	var res []ExperimentMetricAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]ExperimentMetricAPIInputModel, 0)
	} else {
		res = make([]ExperimentMetricAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(PrimaryMetricsValue)
			if ok {
				res[i] = PrimaryMetricToAPIInputModel(ctx, &obj)
			} else {
				obj, ok := elem.(SecondaryMetricsValue)
				if ok {
					res[i] = SecondaryMetricToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type LinkAPIInputModel struct {
	Url   json.RawMessage `json:"url,omitempty"`
	Title json.RawMessage `json:"title,omitempty"`
}

func LinkToAPIInputModel(ctx context.Context, link *LinksValue) LinkAPIInputModel {
	return LinkAPIInputModel{
		Url:   utils.StringAPIField(link.Url),
		Title: utils.StringAPIField(link.Title),
	}
}

func LinkFromAPIModel(ctx context.Context, diags diag.Diagnostics, link *LinksValue, res LinkAPIModel) {
	link.Url = utils.StringToNilableValue(res.Url)
	link.Title = utils.StringToNilableValue(res.Title)
}

func LinksToAPIInputModel(ctx context.Context, list basetypes.ListValue) []LinkAPIInputModel {
	var res []LinkAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]LinkAPIInputModel, 0)
	} else {
		res = make([]LinkAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(LinksValue)
			if !ok {
				return nil
			}

			res[i] = LinkToAPIInputModel(ctx, &obj)
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
