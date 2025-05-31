package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_gate"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// API data model for GateModel (NOTE: see if we can get Terraform to also codegen this from OpenAPI)
type GateAPIModel struct {
	Id                 string                     `json:"id,omitempty"`   // (Name)
	Name               string                     `json:"name,omitempty"` // (Display name)
	IdType             string                     `json:"idType,omitempty"`
	Description        string                     `json:"description"`
	IsEnabled          bool                       `json:"isEnabled"`
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

func GateToAPIModel(ctx context.Context, gate *resource_gate.GateModel) GateAPIModel {
	return GateAPIModel{
		Id:                 gate.Id.ValueString(),
		Name:               gate.Name.ValueString(),
		IdType:             gate.IdType.ValueString(),
		Description:        gate.Description.ValueString(),
		IsEnabled:          BoolFromBoolValue(gate.IsEnabled),
		MeasureMetricLifts: NilableBoolFromBoolValue(gate.MeasureMetricLifts),
		MonitoringMetrics:  MonitoringMetricsToAPIModel(ctx, gate.MonitoringMetrics),
		Rules:              RulesToAPIModel(ctx, gate.Rules),
		Tags:               StringSliceFromListValue(ctx, gate.Tags),
		Type:               gate.Type.ValueString(),
		TargetApps:         StringSliceFromListValue(ctx, gate.TargetApps),
		CreatorId:          gate.CreatorId.ValueString(),
		CreatorEmail:       gate.CreatorEmail.ValueString(),
		Team:               gate.Team.ValueString(),
	}
}

func GateFromAPIModel(ctx context.Context, diags diag.Diagnostics, gate *resource_gate.GateModel, res GateAPIModel) {
	gate.Id = StringToNilableValue(res.Id)
	gate.Name = StringToNilableValue(res.Name)
	gate.IdType = StringToNilableValue(res.IdType)
	gate.Description = StringToNilableValue(res.Description)
	gate.IsEnabled = BoolToBoolValue(res.IsEnabled)
	gate.MeasureMetricLifts = NilableBoolToBoolValue(res.MeasureMetricLifts)
	gate.MonitoringMetrics = MonitoringMetricsFromAPIModel(ctx, diags, res.MonitoringMetrics)
	gate.Rules = RulesFromAPIModel(ctx, diags, res.Rules)
	gate.Tags = StringSliceToListValue(ctx, diags, res.Tags)
	gate.Type = StringToNilableValue(res.Type)
	gate.TargetApps = StringSliceToListValue(ctx, diags, res.TargetApps)
	gate.CreatorId = StringToNilableValue(res.CreatorId)
	gate.CreatorEmail = StringToNilableValue(res.CreatorEmail)
	gate.Team = StringToNilableValue(res.Team)
}

type MonitoringMetricAPIModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func MonitoringMetricToAPIModel(ctx context.Context, metric *resource_gate.MonitoringMetricsValue) MonitoringMetricAPIModel {
	return MonitoringMetricAPIModel{
		Name: StringFromNilableValue(metric.Name),
		Type: StringFromNilableValue(metric.MonitoringMetricsType),
	}
}

func MonitoringMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *resource_gate.MonitoringMetricsValue, res MonitoringMetricAPIModel) {
	metric.Name = StringToNilableValue(res.Name)
	metric.MonitoringMetricsType = StringToNilableValue(res.Type)
}

func MonitoringMetricsToAPIModel(ctx context.Context, list basetypes.ListValue) []MonitoringMetricAPIModel {
	var res []MonitoringMetricAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MonitoringMetricAPIModel, 0)
	} else {
		res = make([]MonitoringMetricAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_gate.MonitoringMetricsValue)
			if !ok {
				return nil
			}

			res[i] = MonitoringMetricToAPIModel(ctx, &obj)
		}
	}
	return res
}

func MonitoringMetricsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []MonitoringMetricAPIModel) basetypes.ListValue {
	attrTypes := resource_gate.MonitoringMetricsValue{}.AttributeTypes(ctx)
	monitoringMetricsType := resource_gate.MonitoringMetricsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil || len(list) == 0 {
		return types.ListNull(monitoringMetricsType)
	} else {
		metrics := make([]attr.Value, len(list))
		for i, elem := range list {
			var metric resource_gate.MonitoringMetricsValue
			MonitoringMetricFromAPIModel(ctx, diags, &metric, elem)
			obj, d := metric.ToObjectValue(ctx)
			metrics[i] = resource_gate.NewMonitoringMetricsValueMust(attrTypes, obj.Attributes())
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

func RuleToAPIModel(ctx context.Context, rule *resource_gate.RulesValue) RuleAPIModel {
	return RuleAPIModel{
		Id:             StringFromNilableValue(rule.Id),
		BaseID:         StringFromNilableValue(rule.BaseId),
		Name:           StringFromNilableValue(rule.Name),
		PassPercentage: IntFromNumberValue(rule.PassPercentage),
		Conditions:     ConditionsToAPIModel(ctx, rule.Conditions),
		Environments:   StringSliceFromListValue(ctx, rule.Environments),
	}
}

func RuleFromAPIModel(ctx context.Context, diags diag.Diagnostics, rule *resource_gate.RulesValue, res RuleAPIModel) {
	rule.Id = StringToNilableValue(res.Id)
	rule.BaseId = StringToNilableValue(res.BaseID)
	rule.Name = StringToNilableValue(res.Name)
	rule.PassPercentage = IntToNumberValue(res.PassPercentage)
	rule.Conditions = ConditionsFromAPIModel(ctx, diags, res.Conditions)
	rule.Environments = StringSliceToListValue(ctx, diags, res.Environments)
}

func RulesToAPIModel(ctx context.Context, list basetypes.ListValue) []RuleAPIModel {
	var res []RuleAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]RuleAPIModel, 0)
	} else {
		res = make([]RuleAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_gate.RulesValue)
			if !ok {
				return nil
			}

			res[i] = RuleToAPIModel(ctx, &obj)
		}
	}
	return res
}

func RulesFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []RuleAPIModel) basetypes.ListValue {
	attrTypes := resource_gate.RulesValue{}.AttributeTypes(ctx)
	rulesType := resource_gate.RulesType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(rulesType)
	} else {
		rules := make([]attr.Value, len(list))
		for i, elem := range list {
			var rule resource_gate.RulesValue
			RuleFromAPIModel(ctx, diags, &rule, elem)
			obj, d := rule.ToObjectValue(ctx)
			rules[i] = resource_gate.NewRulesValueMust(attrTypes, obj.Attributes())
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

func ConditionToAPIModel(ctx context.Context, condition *resource_gate.ConditionsValue) ConditionAPIModel {
	return ConditionAPIModel{
		TargetValue: StringSliceFromListValue(ctx, condition.TargetValue),
		Operator:    StringFromNilableValue(condition.Operator),
		Field:       StringFromNilableValue(condition.Field),
		CustomID:    StringFromNilableValue(condition.CustomId),
		Type:        StringFromNilableValue(condition.ConditionsType),
	}
}

func ConditionFromAPIModel(ctx context.Context, diags diag.Diagnostics, condition *resource_gate.ConditionsValue, res ConditionAPIModel) {
	condition.TargetValue = StringSliceToListValue(ctx, diags, res.TargetValue)
	condition.Operator = StringToNilableValue(res.Operator)
	condition.Field = StringToNilableValue(res.Field)
	condition.CustomId = StringToNilableValue(res.CustomID)
	condition.ConditionsType = StringToNilableValue(res.Type)
}

func ConditionsToAPIModel(ctx context.Context, list basetypes.ListValue) []ConditionAPIModel {
	var res []ConditionAPIModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]ConditionAPIModel, 0)
	} else {
		res = make([]ConditionAPIModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(resource_gate.ConditionsValue)
			if !ok {
				return nil
			}

			res[i] = ConditionToAPIModel(ctx, &obj)
		}
	}
	return res
}

func ConditionsFromAPIModel(ctx context.Context, diags diag.Diagnostics, list []ConditionAPIModel) basetypes.ListValue {
	attrTypes := resource_gate.ConditionsValue{}.AttributeTypes(ctx)
	conditionsType := resource_gate.ConditionsType{
		ObjectType: types.ObjectType{
			AttrTypes: attrTypes,
		},
	}
	if list == nil {
		return types.ListNull(conditionsType)
	} else {
		conditions := make([]attr.Value, len(list))
		for i, elem := range list {
			var condition resource_gate.ConditionsValue
			ConditionFromAPIModel(ctx, diags, &condition, elem)
			obj, d := condition.ToObjectValue(ctx)
			conditions[i] = resource_gate.NewConditionsValueMust(attrTypes, obj.Attributes())
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
