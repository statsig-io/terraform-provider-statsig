package resource_gate

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
)

// API data model for GateModel (NOTE: see if we can get Terraform to also codegen this from OpenAPI)
type GateAPIModel struct {
	Id                 string                     `json:"id,omitempty"`   // (Name)
	Name               string                     `json:"name,omitempty"` // (Display name)
	IdType             string                     `json:"idType,omitempty"`
	Description        string                     `json:"description"`
	IsEnabled          bool                       `json:"isEnabled"`
	IsTemplate         *bool                      `json:"isTemplate,omitempty"`
	MeasureMetricLifts *bool                      `json:"measureMetricLifts,omitempty"`
	MonitoringMetrics  []MonitoringMetricAPIModel `json:"monitoringMetrics,omitempty"`
	Rules              []RuleAPIModel             `json:"rules"`
	Tags               []string                   `json:"tags,omitempty"`
	Type               string                     `json:"type,omitempty"`
	TargetApps         []string                   `json:"targetApps,omitempty"`
	CreatorId          string                     `json:"creatorID,omitempty"`
	CreatorEmail       string                     `json:"creatorEmail,omitempty"`
	Team               string                     `json:"team,omitempty"`
}

func GateToAPIModel(ctx context.Context, gate *GateModel) GateAPIModel {
	return GateAPIModel{
		Id:                 gate.Id.ValueString(),
		Name:               gate.Name.ValueString(),
		IdType:             gate.IdType.ValueString(),
		Description:        gate.Description.ValueString(),
		IsEnabled:          utils.BoolFromBoolValue(gate.IsEnabled),
		IsTemplate:         utils.NilableBoolFromBoolValue(gate.IsTemplate),
		MeasureMetricLifts: utils.NilableBoolFromBoolValue(gate.MeasureMetricLifts),
		MonitoringMetrics:  MonitoringMetricsToAPIModel(ctx, gate.MonitoringMetrics),
		Rules:              RulesToAPIModel(ctx, gate.Rules),
		Tags:               utils.StringSliceFromListValue(ctx, gate.Tags),
		Type:               gate.Type.ValueString(),
		TargetApps:         utils.StringSliceFromListValue(ctx, gate.TargetApps),
		CreatorId:          gate.CreatorId.ValueString(),
		CreatorEmail:       gate.CreatorEmail.ValueString(),
		Team:               gate.Team.ValueString(),
	}
}

func GateFromAPIModel(ctx context.Context, diags diag.Diagnostics, gate *GateModel, res GateAPIModel) {
	gate.Id = utils.StringToNilableValue(res.Id)
	gate.Name = utils.StringToNilableValue(res.Name)
	gate.IdType = utils.StringToNilableValue(res.IdType)
	gate.Description = utils.StringToNilableValue(res.Description)
	gate.IsEnabled = utils.BoolToBoolValue(res.IsEnabled)
	gate.IsTemplate = utils.NilableBoolToBoolValue(res.IsTemplate)
	gate.MeasureMetricLifts = utils.NilableBoolToBoolValue(res.MeasureMetricLifts)
	gate.MonitoringMetrics = MonitoringMetricsFromAPIModel(ctx, diags, res.MonitoringMetrics)
	gate.Rules = RulesFromAPIModel(ctx, diags, res.Rules)
	gate.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	gate.Type = utils.StringToNilableValue(res.Type)
	gate.TargetApps = utils.StringSliceToListValue(ctx, diags, res.TargetApps)
	gate.CreatorId = utils.StringToNilableValue(res.CreatorId)
	gate.CreatorEmail = utils.StringToNilableValue(res.CreatorEmail)
	gate.Team = utils.StringToNilableValue(res.Team)
}

type MonitoringMetricAPIModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func MonitoringMetricToAPIModel(ctx context.Context, metric *MonitoringMetricsValue) MonitoringMetricAPIModel {
	return MonitoringMetricAPIModel{
		Name: utils.StringFromNilableValue(metric.Name),
		Type: utils.StringFromNilableValue(metric.MonitoringMetricsType),
	}
}

func MonitoringMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *MonitoringMetricsValue, res MonitoringMetricAPIModel) {
	metric.Name = utils.StringToNilableValue(res.Name)
	metric.MonitoringMetricsType = utils.StringToNilableValue(res.Type)
}

func MonitoringMetricsToAPIModel(ctx context.Context, list basetypes.ListValue) []MonitoringMetricAPIModel {
	var res []MonitoringMetricAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MonitoringMetricAPIModel, 0)
	} else {
		res = make([]MonitoringMetricAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(MonitoringMetricsValue)
			if !ok {
				return nil
			}

			res[i] = MonitoringMetricToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MonitoringMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MonitoringMetricAPIModel) basetypes.ListValue {
	attrTypes := MonitoringMetricsValue{}.AttributeTypes(ctx)
	monitoringMetricsType := MonitoringMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil || len(list) == 0 {
		return types.ListNull(monitoringMetricsType)
	} else {
		metrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metric MonitoringMetricsValue
			MonitoringMetricFromAPIModel(ctx, diags, &metric, elem)
			obj, d := metric.ToObjectValue(ctx)
			metrics[i] = NewMonitoringMetricsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(monitoringMetricsType, metrics)
		diags = append(diags, d...)
		return v
	}
}

type RuleAPIModel struct {
	Id             string              `json:"id"`
	BaseID         string              `json:"baseID,omitempty"`
	Name           string              `json:"name"`
	PassPercentage int                 `json:"passPercentage"`
	Conditions     []ConditionAPIModel `json:"conditions"`
	Environments   []string            `json:"environments,omitempty"`
}

func RuleToAPIModel(ctx context.Context, rule *RulesValue) RuleAPIModel {
	return RuleAPIModel{
		Id:             utils.StringFromNilableValue(rule.Id),
		BaseID:         utils.StringFromNilableValue(rule.BaseId),
		Name:           utils.StringFromNilableValue(rule.Name),
		PassPercentage: utils.IntFromNumberValue(rule.PassPercentage),
		Conditions:     ConditionsToAPIModel(ctx, rule.Conditions),
		Environments:   utils.StringSliceFromListValue(ctx, rule.Environments),
	}
}

func RuleFromAPIModel(ctx context.Context, diags diag.Diagnostics, rule *RulesValue, res RuleAPIModel) {
	rule.Id = utils.StringToNilableValue(res.Id)
	rule.BaseId = utils.StringToNilableValue(res.BaseID)
	rule.Name = utils.StringToNilableValue(res.Name)
	rule.PassPercentage = utils.IntToNumberValue(res.PassPercentage)
	rule.Conditions = ConditionsFromAPIModel(ctx, diags, res.Conditions)
	rule.Environments = utils.StringSliceToListValue(ctx, diags, res.Environments)
}

func RulesToAPIModel(ctx context.Context, list basetypes.ListValue) []RuleAPIModel {
	var res []RuleAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]RuleAPIModel, 0)
	} else {
		res = make([]RuleAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(RulesValue)
			if !ok {
				return nil
			}

			res[i] = RuleToAPIModel(ctx, &obj)
		}
	}
	return res
}

func RulesFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []RuleAPIModel) basetypes.ListValue {
	attrTypes := RulesValue{}.AttributeTypes(ctx)
	rulesType := RulesType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(rulesType)
	} else {
		rules := make([]attr.Value, len(list))
		for i, elem := range list {
			var rule RulesValue
			RuleFromAPIModel(ctx, diags, &rule, elem)
			obj, d := rule.ToObjectValue(ctx)
			rules[i] = NewRulesValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(rulesType, rules)
		diags = append(diags, d...)
		return v
	}
}

type ConditionAPIModel struct {
	TargetValue TargetValue `json:"targetValue"`
	Operator    string      `json:"operator"`
	Field       string      `json:"field,omitempty"`
	CustomID    string      `json:"customID,omitempty"`
	Type        string      `json:"type"`
}

func ConditionToAPIModel(ctx context.Context, condition *ConditionsValue) ConditionAPIModel {
	return ConditionAPIModel{
		TargetValue: utils.StringSliceFromListValue(ctx, condition.TargetValue),
		Operator:    utils.StringFromNilableValue(condition.Operator),
		Field:       utils.StringFromNilableValue(condition.Field),
		CustomID:    utils.StringFromNilableValue(condition.CustomId),
		Type:        utils.StringFromNilableValue(condition.ConditionsType),
	}
}

func ConditionFromAPIModel(ctx context.Context, diags diag.Diagnostics, condition *ConditionsValue, res ConditionAPIModel) {
	condition.TargetValue = utils.StringSliceToListValue(ctx, diags, res.TargetValue)
	condition.Operator = utils.StringToNilableValue(res.Operator)
	condition.Field = utils.StringToNilableValue(res.Field)
	condition.CustomId = utils.StringToNilableValue(res.CustomID)
	condition.ConditionsType = utils.StringToNilableValue(res.Type)
}

func ConditionsToAPIModel(ctx context.Context, list basetypes.ListValue) []ConditionAPIModel {
	var res []ConditionAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]ConditionAPIModel, 0)
	} else {
		res = make([]ConditionAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(ConditionsValue)
			if !ok {
				return nil
			}

			res[i] = ConditionToAPIModel(ctx, &obj)
		}
	}
	return res
}

func ConditionsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []ConditionAPIModel) basetypes.ListValue {
	attrTypes := ConditionsValue{}.AttributeTypes(ctx)
	conditionsType := ConditionsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(conditionsType)
	} else {
		conditions := make([]attr.Value, len(list))
		for i, elem := range list {
			var condition ConditionsValue
			ConditionFromAPIModel(ctx, diags, &condition, elem)
			obj, d := condition.ToObjectValue(ctx)
			conditions[i] = NewConditionsValueMust(attrTypes, obj.Attributes())
			diags = append(diags, d...)
		}
		v, d := types.ListValue(conditionsType, conditions)
		diags = append(diags, d...)
		return v
	}
}

type TargetValue []string

// Custom unmarshal function to parse all possible types of target value as a string slice
func (t *TargetValue) UnmarshalJSON(data []byte) error {
	// Try to unmarshal target value as a string slice
	var stringSliceTargetValue []string
	if err := json.Unmarshal(data, &stringSliceTargetValue); err == nil {
		*t = stringSliceTargetValue
		return nil
	}

	// Try to unmarshal target value as single string
	var stringTargetValue string
	if err := json.Unmarshal(data, &stringTargetValue); err == nil {
		*t = []string{stringTargetValue}
		return nil
	}

	// Try to unmarshal target value as a float slice
	var numberSliceTargetValue []float64
	if err := json.Unmarshal(data, &numberSliceTargetValue); err == nil {
		for _, element := range numberSliceTargetValue {
			*t = append(*t, strconv.FormatFloat(element, 'f', -1, 64))
		}
		return nil
	}

	// Try to unmarshal target value as a single float
	var numberTargetValue float64
	if err := json.Unmarshal(data, &numberTargetValue); err == nil {
		*t = []string{strconv.FormatFloat(numberTargetValue, 'f', -1, 64)}
		return nil
	}

	return fmt.Errorf("Cannot unmarshal targetValue: %s", string(data))
}
