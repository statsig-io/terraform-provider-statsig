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

func SegmentToAPIModel(ctx context.Context, segment *SegmentModel) SegmentAPIModel {
	return SegmentAPIModel{
		Id:                segment.Id.ValueString(),
		Name:              segment.Name.ValueString(),
		IdType:            segment.IdType.ValueString(),
		Description:       segment.Description.ValueString(),
		IsEnabled:         segment.IsEnabled.ValueBool(),
		Rules:             RulesToAPIModel(ctx, segment.Rules),
		Type:              segment.Type.ValueString(),
		CreatorId:         segment.CreatorId.ValueString(),
		CreatorName:       segment.CreatorName.ValueString(),
		CreatorEmail:      segment.CreatorEmail.ValueString(),
		CreatedTime:       segment.CreatedTime.ValueFloat64(),
		LastModifierId:    segment.LastModifierId.ValueString(),
		LastModifierName:  segment.LastModifierName.ValueString(),
		LastModifierEmail: segment.LastModifierEmail.ValueString(),
		LastModifiedTime:  segment.LastModifiedTime.ValueFloat64(),
		HoldoutIds:        utils.StringSliceFromListValue(ctx, segment.HoldoutIds),
		Tags:              utils.StringSliceFromListValue(ctx, segment.Tags),
		TargetApps:        utils.StringSliceFromListValue(ctx, segment.TargetApps),
		Team:              segment.Team.ValueString(),
		TeamId:            segment.TeamId.ValueString(),
		Version:           segment.Version.ValueFloat64(),
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

type SegmentRulesAPIModel struct {
	Rules []RuleAPIModel `json:"rules"`
}

func SegmentToRulesAPIModel(ctx context.Context, segment *SegmentModel) SegmentRulesAPIModel {
	return SegmentRulesAPIModel{
		Rules: RulesToAPIModel(ctx, segment.Rules),
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
