package resource_dynamic_config

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

// API data model for DynamicConfigModel
type DynamicConfigAPIModel struct {
	Id                string                 `json:"id,omitempty"`
	Name              string                 `json:"name,omitempty"`
	IdType            string                 `json:"idType,omitempty"`
	Description       string                 `json:"description"`
	IsEnabled         bool                   `json:"isEnabled"`
	IsTemplate        *bool                  `json:"isTemplate,omitempty"`
	Rules             []RuleAPIModel         `json:"rules"`
	Schema            string                 `json:"schema,omitempty"`
	SchemaJson5       string                 `json:"schemaJson5,omitempty"`
	DefaultValue      map[string]interface{} `json:"defaultValue"`
	DefaultValueJson5 string                 `json:"defaultValueJson5"`
	Tags              []string               `json:"tags,omitempty"`
	TargetApps        []string               `json:"targetApps,omitempty"`
	CreatorId         string                 `json:"creatorID,omitempty"`
	CreatorEmail      string                 `json:"creatorEmail,omitempty"`
	Team              string                 `json:"team,omitempty"`
}

// Every field is raw JSON so the request can leave out an attribute the
// configuration never mentioned. See utils.APIField.
type DynamicConfigAPIInputModel struct {
	Id                json.RawMessage `json:"id,omitempty"`
	Name              json.RawMessage `json:"name,omitempty"`
	IdType            json.RawMessage `json:"idType,omitempty"`
	Description       json.RawMessage `json:"description,omitempty"`
	IsEnabled         json.RawMessage `json:"isEnabled,omitempty"`
	IsTemplate        json.RawMessage `json:"isTemplate,omitempty"`
	Rules             json.RawMessage `json:"rules,omitempty"`
	Schema            json.RawMessage `json:"schema,omitempty"`
	SchemaJson5       json.RawMessage `json:"schemaJson5,omitempty"`
	DefaultValue      json.RawMessage `json:"defaultValue,omitempty"`
	DefaultValueJson5 json.RawMessage `json:"defaultValueJson5,omitempty"`
	Tags              json.RawMessage `json:"tags,omitempty"`
	TargetApps        json.RawMessage `json:"targetApps,omitempty"`
	CreatorId         json.RawMessage `json:"creatorID,omitempty"`
	CreatorEmail      json.RawMessage `json:"creatorEmail,omitempty"`
	Team              json.RawMessage `json:"team,omitempty"`
}

func DynamicConfigToAPIInputModel(ctx context.Context, dynamicConfig *DynamicConfigModel) DynamicConfigAPIInputModel {
	return DynamicConfigAPIInputModel{
		Id:                utils.StringAPIField(dynamicConfig.Id),
		Name:              utils.StringAPIField(dynamicConfig.Name),
		IdType:            utils.StringAPIField(dynamicConfig.IdType),
		Description:       utils.StringAPIField(dynamicConfig.Description),
		IsEnabled:         utils.BoolAPIField(dynamicConfig.IsEnabled),
		IsTemplate:        utils.BoolAPIField(dynamicConfig.IsTemplate),
		Rules:             utils.APIField(dynamicConfig.Rules, RulesToAPIModel(ctx, dynamicConfig.Rules)),
		Schema:            utils.StringAPIField(dynamicConfig.Schema),
		SchemaJson5:       utils.StringAPIField(dynamicConfig.SchemaJson5),
		DefaultValue:      utils.MapAPIField(ctx, dynamicConfig.DefaultValue),
		DefaultValueJson5: utils.StringAPIField(dynamicConfig.DefaultValueJson5),
		Tags:              utils.StringSliceAPIField(ctx, dynamicConfig.Tags),
		TargetApps:        utils.StringSliceAPIField(ctx, dynamicConfig.TargetApps),
		CreatorId:         utils.StringAPIField(dynamicConfig.CreatorId),
		CreatorEmail:      utils.StringAPIField(dynamicConfig.CreatorEmail),
		Team:              utils.StringAPIField(dynamicConfig.Team),
	}
}

func DynamicConfigFromAPIModel(ctx context.Context, diags diag.Diagnostics, dynamicConfig *DynamicConfigModel, res DynamicConfigAPIModel) {
	dynamicConfig.Id = utils.StringToNilableValue(res.Id)
	dynamicConfig.Name = utils.StringToNilableValue(res.Name)
	dynamicConfig.IdType = utils.StringToNilableValue(res.IdType)
	dynamicConfig.Description = utils.StringToNilableValue(res.Description)
	dynamicConfig.IsEnabled = utils.BoolToBoolValue(res.IsEnabled)
	dynamicConfig.IsTemplate = utils.NilableBoolToBoolValue(res.IsTemplate)
	dynamicConfig.Rules = RulesFromAPIModel(ctx, diags, res.Rules)
	dynamicConfig.Schema = utils.StringToStringValue(res.Schema)
	dynamicConfig.SchemaJson5 = utils.StringToStringValue(res.SchemaJson5)
	dynamicConfig.DefaultValue = utils.MapToMapValue(ctx, diags, res.DefaultValue)
	dynamicConfig.DefaultValueJson5 = utils.StringToStringValue(res.DefaultValueJson5)
	dynamicConfig.Tags = utils.StringSliceToListValue(ctx, diags, res.Tags)
	dynamicConfig.TargetApps = utils.StringSliceToListValue(ctx, diags, res.TargetApps)
	dynamicConfig.CreatorId = utils.StringToNilableValue(res.CreatorId)
	dynamicConfig.CreatorEmail = utils.StringToNilableValue(res.CreatorEmail)
	dynamicConfig.Team = utils.StringToNilableValue(res.Team)
}

type RuleAPIModel struct {
	Id               string                 `json:"id,omitempty"`
	BaseId           string                 `json:"baseID,omitempty"`
	Name             string                 `json:"name"`
	PassPercentage   int                    `json:"passPercentage"`
	Conditions       []ConditionAPIModel    `json:"conditions"`
	Environments     []string               `json:"environments"`
	ReturnValue      map[string]interface{} `json:"returnValue"`
	ReturnValueJson5 string                 `json:"returnValueJson5"`
}

func RuleToAPIModel(ctx context.Context, rule *RulesValue) RuleAPIModel {
	return RuleAPIModel{
		Id:               utils.StringFromNilableValue(rule.Id),
		BaseId:           utils.StringFromNilableValue(rule.BaseId),
		Name:             utils.StringFromNilableValue(rule.Name),
		PassPercentage:   utils.IntFromNumberValue(rule.PassPercentage),
		Conditions:       ConditionsToAPIModel(ctx, rule.Conditions),
		Environments:     utils.StringSliceFromListValue(ctx, rule.Environments),
		ReturnValue:      utils.MapFromMapValue(ctx, rule.ReturnValue),
		ReturnValueJson5: utils.StringFromNilableValue(rule.ReturnValueJson5),
	}
}

func RuleFromAPIModel(ctx context.Context, diags diag.Diagnostics, rule *RulesValue, res RuleAPIModel) {
	rule.Id = utils.StringToNilableValue(res.Id)
	rule.BaseId = utils.StringToNilableValue(res.BaseId)
	rule.Name = utils.StringToNilableValue(res.Name)
	rule.PassPercentage = utils.IntToNumberValue(res.PassPercentage)
	rule.Conditions = ConditionsFromAPIModel(ctx, diags, res.Conditions)
	rule.Environments = utils.StringSliceToListValue(ctx, diags, res.Environments)
	rule.ReturnValue = utils.MapToMapValue(ctx, diags, res.ReturnValue)
	rule.ReturnValueJson5 = utils.StringToStringValue(res.ReturnValueJson5)
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

	return fmt.Errorf("cannot unmarshal targetValue: %s", string(data))
}
