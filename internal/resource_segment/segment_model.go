package resource_segment

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

type SegmentAPIModel struct {
	Id                string         `json:"id,omitempty"`   // (Name)
	Name              string         `json:"name,omitempty"` // (Display name)
	IdType            string         `json:"idType,omitempty"`
	Description       string         `json:"description"`
	IsEnabled         bool           `json:"isEnabled"`
	Rules             []RuleAPIModel `json:"rules"`
	Type              string         `json:"type,omitempty"`
	CreatorId         string         `json:"creatorID,omitempty"`
	CreatorName       string         `json:"creatorName,omitempty"`
	CreatorEmail      string         `json:"creatorEmail,omitempty"`
	CreatedTime       float64        `json:"createdTime"`
	LastModifierId    string         `json:"lastModifiedID,omitempty"`
	LastModifierName  string         `json:"lastModifiedName,omitempty"`
	LastModifierEmail string         `json:"lastModifiedEmail,omitempty"`
	LastModifiedTime  float64        `json:"lastModifiedTime,omitempty"`
	HoldoutIds        []string       `json:"holdoutIDs"`
	Tags              []string       `json:"tags,omitempty"`
	TargetApps        []string       `json:"targetApps,omitempty"`
	Team              string         `json:"team,omitempty"`
	TeamId            string         `json:"teamID,omitempty"`
	Version           float64        `json:"version,omitempty"`
}

// Every field is raw JSON so the request can leave out an attribute the
// configuration never mentioned. See utils.APIField.
type SegmentAPIInputModel struct {
	Id                json.RawMessage `json:"id,omitempty"`   // (Name)
	Name              json.RawMessage `json:"name,omitempty"` // (Display name)
	IdType            json.RawMessage `json:"idType,omitempty"`
	Description       json.RawMessage `json:"description,omitempty"`
	IsEnabled         json.RawMessage `json:"isEnabled,omitempty"`
	Rules             json.RawMessage `json:"rules,omitempty"`
	Type              json.RawMessage `json:"type,omitempty"`
	CreatorId         json.RawMessage `json:"creatorID,omitempty"`
	CreatorName       json.RawMessage `json:"creatorName,omitempty"`
	CreatorEmail      json.RawMessage `json:"creatorEmail,omitempty"`
	CreatedTime       json.RawMessage `json:"createdTime,omitempty"`
	LastModifierId    json.RawMessage `json:"lastModifiedID,omitempty"`
	LastModifierName  json.RawMessage `json:"lastModifiedName,omitempty"`
	LastModifierEmail json.RawMessage `json:"lastModifiedEmail,omitempty"`
	LastModifiedTime  json.RawMessage `json:"lastModifiedTime,omitempty"`
	HoldoutIds        json.RawMessage `json:"holdoutIDs,omitempty"`
	Tags              json.RawMessage `json:"tags,omitempty"`
	TargetApps        json.RawMessage `json:"targetApps,omitempty"`
	Team              json.RawMessage `json:"team,omitempty"`
	TeamId            json.RawMessage `json:"teamID,omitempty"`
	Version           json.RawMessage `json:"version,omitempty"`
}

func SegmentToAPIInputModel(ctx context.Context, segment *SegmentModel) SegmentAPIInputModel {
	return SegmentAPIInputModel{
		Id:                utils.StringAPIField(segment.Id),
		Name:              utils.StringAPIField(segment.Name),
		IdType:            utils.StringAPIField(segment.IdType),
		Description:       utils.StringAPIField(segment.Description),
		IsEnabled:         utils.BoolAPIField(segment.IsEnabled),
		Rules:             utils.APIField(segment.Rules, RulesToAPIInputModel(ctx, segment.Rules)),
		Type:              utils.StringAPIField(segment.Type),
		CreatorId:         utils.StringAPIField(segment.CreatorId),
		CreatorName:       utils.StringAPIField(segment.CreatorName),
		CreatorEmail:      utils.StringAPIField(segment.CreatorEmail),
		CreatedTime:       utils.FloatAPIField(segment.CreatedTime),
		LastModifierId:    utils.StringAPIField(segment.LastModifierId),
		LastModifierName:  utils.StringAPIField(segment.LastModifierName),
		LastModifierEmail: utils.StringAPIField(segment.LastModifierEmail),
		LastModifiedTime:  utils.FloatAPIField(segment.LastModifiedTime),
		HoldoutIds:        utils.StringSliceAPIField(ctx, segment.HoldoutIds),
		Tags:              utils.StringSliceAPIField(ctx, segment.Tags),
		TargetApps:        utils.StringSliceAPIField(ctx, segment.TargetApps),
		Team:              utils.StringAPIField(segment.Team),
		TeamId:            utils.StringAPIField(segment.TeamId),
		Version:           utils.FloatAPIField(segment.Version),
	}
}

func SegmentFromAPIModel(ctx context.Context, diags diag.Diagnostics, segment *SegmentModel, res SegmentAPIModel) {
	segment.Id = utils.StringToNilableValue(res.Id)
	segment.Name = utils.StringToNilableValue(res.Name)
	segment.IdType = utils.StringToNilableValue(res.IdType)
	segment.Description = utils.StringToNilableValue(res.Description)
	segment.IsEnabled = utils.BoolToBoolValue(res.IsEnabled)
	segment.Rules = RulesFromAPIModel(ctx, diags, res.Rules)
	segment.Type = utils.StringToNilableValue(res.Type)
	segment.CreatorId = utils.StringToNilableValue(res.CreatorId)
	segment.CreatorName = utils.StringToNilableValue(res.CreatorName)
	segment.CreatorEmail = utils.StringToNilableValue(res.CreatorEmail)
	segment.CreatedTime = utils.FloatToFloatValue(res.CreatedTime)
	segment.LastModifierId = utils.StringToNilableValue(res.LastModifierId)
	segment.LastModifierName = utils.StringToNilableValue(res.LastModifierName)
	segment.LastModifierEmail = utils.StringToNilableValue(res.LastModifierEmail)
	segment.LastModifiedTime = utils.FloatToFloatValue(res.LastModifiedTime)
	segment.HoldoutIds = utils.StringSliceToListValue(ctx, diags, res.HoldoutIds)
	segment.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	segment.TargetApps = utils.StringSliceToListValue(ctx, diags, res.TargetApps)
	segment.Team = utils.StringToNilableValue(res.Team)
	segment.TeamId = utils.StringToNilableValue(res.TeamId)
	segment.Version = utils.FloatToFloatValue(res.Version)
}

type SegmentRulesAPIInputModel struct {
	Rules json.RawMessage `json:"rules,omitempty"`
}

func SegmentToRulesAPIInputModel(ctx context.Context, segment *SegmentModel) SegmentRulesAPIInputModel {
	return SegmentRulesAPIInputModel{
		Rules: utils.APIField(segment.Rules, RulesToAPIInputModel(ctx, segment.Rules)),
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
