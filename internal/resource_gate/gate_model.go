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

// Every field is raw JSON so the request can leave out an attribute the
// configuration never mentioned. See utils.APIField.
type GateAPIInputModel struct {
	Id                 json.RawMessage `json:"id,omitempty"`   // (Name)
	Name               json.RawMessage `json:"name,omitempty"` // (Display name)
	IdType             json.RawMessage `json:"idType,omitempty"`
	Description        json.RawMessage `json:"description,omitempty"`
	IsEnabled          json.RawMessage `json:"isEnabled,omitempty"`
	IsTemplate         json.RawMessage `json:"isTemplate,omitempty"`
	MeasureMetricLifts json.RawMessage `json:"measureMetricLifts,omitempty"`
	MonitoringMetrics  json.RawMessage `json:"monitoringMetrics,omitempty"`
	Rules              json.RawMessage `json:"rules,omitempty"`
	Tags               json.RawMessage `json:"tags,omitempty"`
	Type               json.RawMessage `json:"type,omitempty"`
	TargetApps         json.RawMessage `json:"targetApps,omitempty"`
	CreatorId          json.RawMessage `json:"creatorID,omitempty"`
	CreatorEmail       json.RawMessage `json:"creatorEmail,omitempty"`
	Team               json.RawMessage `json:"team,omitempty"`
}

func GateToAPIInputModel(ctx context.Context, gate *GateModel) GateAPIInputModel {
	return GateAPIInputModel{
		Id:                 utils.StringAPIField(gate.Id),
		Name:               utils.StringAPIField(gate.Name),
		IdType:             utils.StringAPIField(gate.IdType),
		Description:        utils.StringAPIField(gate.Description),
		IsEnabled:          utils.BoolAPIField(gate.IsEnabled),
		IsTemplate:         utils.BoolAPIField(gate.IsTemplate),
		MeasureMetricLifts: utils.BoolAPIField(gate.MeasureMetricLifts),
		MonitoringMetrics:  utils.APIField(gate.MonitoringMetrics, MonitoringMetricsToAPIInputModel(ctx, gate.MonitoringMetrics)),
		Rules:              utils.APIField(gate.Rules, RulesToAPIInputModel(ctx, gate.Rules)),
		Tags:               utils.StringSliceAPIField(ctx, gate.Tags),
		Type:               utils.StringAPIField(gate.Type),
		TargetApps:         utils.StringSliceAPIField(ctx, gate.TargetApps),
		CreatorId:          utils.StringAPIField(gate.CreatorId),
		CreatorEmail:       utils.StringAPIField(gate.CreatorEmail),
		Team:               utils.StringAPIField(gate.Team),
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type MonitoringMetricAPIInputModel struct {
	Name json.RawMessage `json:"name,omitempty"`
	Type json.RawMessage `json:"type,omitempty"`
}

func MonitoringMetricToAPIInputModel(ctx context.Context, metric *MonitoringMetricsValue) MonitoringMetricAPIInputModel {
	return MonitoringMetricAPIInputModel{
		Name: utils.StringAPIField(metric.Name),
		Type: utils.StringAPIField(metric.MonitoringMetricsType),
	}
}

func MonitoringMetricFromAPIModel(ctx context.Context, diags diag.Diagnostics, metric *MonitoringMetricsValue, res MonitoringMetricAPIModel) {
	metric.Name = utils.StringToNilableValue(res.Name)
	metric.MonitoringMetricsType = utils.StringToNilableValue(res.Type)
}

func MonitoringMetricsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []MonitoringMetricAPIInputModel {
	var res []MonitoringMetricAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]MonitoringMetricAPIInputModel, 0)
	} else {
		res = make([]MonitoringMetricAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(MonitoringMetricsValue)
			if !ok {
				return nil
			}

			res[i] = MonitoringMetricToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type RuleAPIInputModel struct {
	Id             json.RawMessage `json:"id,omitempty"`
	BaseID         json.RawMessage `json:"baseID,omitempty"`
	Name           json.RawMessage `json:"name,omitempty"`
	PassPercentage json.RawMessage `json:"passPercentage,omitempty"`
	Conditions     json.RawMessage `json:"conditions,omitempty"`
	Environments   json.RawMessage `json:"environments,omitempty"`
}

func RuleToAPIInputModel(ctx context.Context, rule *RulesValue) RuleAPIInputModel {
	return RuleAPIInputModel{
		Id:             utils.StringAPIField(rule.Id),
		BaseID:         utils.StringAPIField(rule.BaseId),
		Name:           utils.StringAPIField(rule.Name),
		PassPercentage: utils.NumberAPIField(rule.PassPercentage),
		Conditions:     utils.APIField(rule.Conditions, ConditionsToAPIInputModel(ctx, rule.Conditions)),
		Environments:   utils.StringSliceAPIField(ctx, rule.Environments),
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

func RulesToAPIInputModel(ctx context.Context, list basetypes.ListValue) []RuleAPIInputModel {
	var res []RuleAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]RuleAPIInputModel, 0)
	} else {
		res = make([]RuleAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(RulesValue)
			if !ok {
				return nil
			}

			res[i] = RuleToAPIInputModel(ctx, &obj)
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

// Every field is raw JSON so the request can leave out a nested attribute the
// configuration never mentioned. See utils.APIField.
type ConditionAPIInputModel struct {
	TargetValue json.RawMessage `json:"targetValue,omitempty"`
	Operator    json.RawMessage `json:"operator,omitempty"`
	Field       json.RawMessage `json:"field,omitempty"`
	CustomID    json.RawMessage `json:"customID,omitempty"`
	Type        json.RawMessage `json:"type,omitempty"`
}

func ConditionToAPIInputModel(ctx context.Context, condition *ConditionsValue) ConditionAPIInputModel {
	return ConditionAPIInputModel{
		TargetValue: utils.StringSliceAPIField(ctx, condition.TargetValue),
		Operator:    utils.StringAPIField(condition.Operator),
		Field:       utils.StringAPIField(condition.Field),
		CustomID:    utils.StringAPIField(condition.CustomId),
		Type:        utils.StringAPIField(condition.ConditionsType),
	}
}

func ConditionFromAPIModel(ctx context.Context, diags diag.Diagnostics, condition *ConditionsValue, res ConditionAPIModel) {
	condition.TargetValue = utils.StringSliceToListValue(ctx, diags, res.TargetValue)
	condition.Operator = utils.StringToNilableValue(res.Operator)
	condition.Field = utils.StringToNilableValue(res.Field)
	condition.CustomId = utils.StringToNilableValue(res.CustomID)
	condition.ConditionsType = utils.StringToNilableValue(res.Type)
}

func ConditionsToAPIInputModel(ctx context.Context, list basetypes.ListValue) []ConditionAPIInputModel {
	var res []ConditionAPIInputModel
	if list.IsNull() || list.IsUnknown() {
		res = make([]ConditionAPIInputModel, 0)
	} else {
		res = make([]ConditionAPIInputModel, len(list.Elements()))
		for i, elem := range list.Elements() {
			obj, ok := elem.(ConditionsValue)
			if !ok {
				return nil
			}

			res[i] = ConditionToAPIInputModel(ctx, &obj)
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
