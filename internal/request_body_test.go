package statsig

import (
	"context"
	"encoding/json"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/statsig-io/terraform-provider-statsig/internal/resource_dynamic_config"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_experiment"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_gate"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_keys"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_metric"
	"github.com/statsig-io/terraform-provider-statsig/internal/resource_segment"
)

// One rule covers six resources, so the cases live together rather than being
// repeated in six packages: an attribute the configuration never mentioned is
// left out of the request, an attribute set to an empty value is sent and
// clears the field, and a configured value is sent.
//
// The three states are told apart by the encoded body, not by the Go value, so
// every assertion is on the bytes. Before this rule the update path sent the Go
// zero value for an unmentioned attribute, and the Console API reads a present
// field as an instruction: an unrelated edit could clear a description, disable
// an enabled gate, wipe its rules, or revoke a key's scopes.
func TestRequestBodyOmitsUnspecifiedAttributes(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name string
		body any
		// An empty want means the field must not appear in the body at all.
		want map[string]string
	}{
		{
			name: "gate leaves out the attributes the config omits",
			body: resource_gate.GateToAPIInputModel(ctx, &resource_gate.GateModel{
				Id:          types.StringValue("my_gate"),
				Name:        types.StringUnknown(),
				Description: types.StringUnknown(),
				IsEnabled:   types.BoolUnknown(),
				Rules:       types.ListUnknown(resource_gate.RulesValue{}.Type(ctx)),
			}),
			want: map[string]string{
				"id": `"my_gate"`, "name": "", "description": "", "isEnabled": "", "rules": "",
			},
		},
		{
			name: "gate sends explicitly empty attributes, which clear them",
			body: resource_gate.GateToAPIInputModel(ctx, &resource_gate.GateModel{
				Name:        types.StringValue(""),
				Description: types.StringValue(""),
				IsEnabled:   types.BoolValue(false),
				Rules:       types.ListValueMust(resource_gate.RulesValue{}.Type(ctx), []attr.Value{}),
			}),
			want: map[string]string{
				"name": `""`, "description": `""`, "isEnabled": "false", "rules": "[]",
			},
		},
		{
			name: "gate sends configured attributes",
			body: resource_gate.GateToAPIInputModel(ctx, &resource_gate.GateModel{
				Name:        types.StringValue("My Gate"),
				Description: types.StringValue("who sees the new checkout"),
				IsEnabled:   types.BoolValue(true),
				Rules:       gateRules(ctx),
			}),
			want: map[string]string{
				"name":        `"My Gate"`,
				"description": `"who sees the new checkout"`,
				"isEnabled":   "true",
				"rules":       `[{"id":"rule_1","name":"everyone","passPercentage":100,"conditions":[]}]`,
			},
		},
		{
			name: "keys leaves out the attributes the config omits",
			body: resource_keys.KeyToAPIInputModel(ctx, &resource_keys.KeysModel{
				Type:         types.StringValue("SERVER"),
				Scopes:       types.ListUnknown(types.StringType),
				Environments: types.ListUnknown(types.StringType),
			}),
			want: map[string]string{"type": `"SERVER"`, "scopes": "", "environments": ""},
		},
		{
			name: "keys sends explicitly empty attributes, which revoke them",
			body: resource_keys.KeyToAPIInputModel(ctx, &resource_keys.KeysModel{
				Scopes:       types.ListValueMust(types.StringType, []attr.Value{}),
				Environments: types.ListValueMust(types.StringType, []attr.Value{}),
			}),
			want: map[string]string{"scopes": "[]", "environments": "[]"},
		},
		{
			name: "keys sends configured attributes",
			body: resource_keys.KeyToAPIInputModel(ctx, &resource_keys.KeysModel{
				Scopes:       stringList("gates:read"),
				Environments: stringList("production"),
			}),
			want: map[string]string{"scopes": `["gates:read"]`, "environments": `["production"]`},
		},
		{
			name: "segment leaves out the description the config omits",
			body: resource_segment.SegmentToAPIInputModel(ctx, &resource_segment.SegmentModel{
				Id:          types.StringValue("my_segment"),
				Description: types.StringUnknown(),
			}),
			want: map[string]string{"id": `"my_segment"`, "description": ""},
		},
		{
			name: "segment sends an explicitly empty description, which clears it",
			body: resource_segment.SegmentToAPIInputModel(ctx, &resource_segment.SegmentModel{
				Description: types.StringValue(""),
			}),
			want: map[string]string{"description": `""`},
		},
		{
			name: "segment sends a configured description",
			body: resource_segment.SegmentToAPIInputModel(ctx, &resource_segment.SegmentModel{
				Description: types.StringValue("beta testers"),
			}),
			want: map[string]string{"description": `"beta testers"`},
		},
		{
			name: "experiment leaves out the description the config omits",
			body: resource_experiment.ExperimentToAPIInputModel(ctx, &resource_experiment.ExperimentModel{
				Id:          types.StringValue("my_experiment"),
				Description: types.StringUnknown(),
			}),
			want: map[string]string{"id": `"my_experiment"`, "description": ""},
		},
		{
			name: "experiment sends an explicitly empty description, which clears it",
			body: resource_experiment.ExperimentToAPIInputModel(ctx, &resource_experiment.ExperimentModel{
				Description: types.StringValue(""),
			}),
			want: map[string]string{"description": `""`},
		},
		{
			name: "experiment sends a configured description",
			body: resource_experiment.ExperimentToAPIInputModel(ctx, &resource_experiment.ExperimentModel{
				Description: types.StringValue("checkout redesign"),
			}),
			want: map[string]string{"description": `"checkout redesign"`},
		},
		{
			name: "dynamic config leaves out the is_enabled the config omits",
			body: resource_dynamic_config.DynamicConfigToAPIInputModel(ctx, &resource_dynamic_config.DynamicConfigModel{
				Id:        types.StringValue("my_config"),
				IsEnabled: types.BoolUnknown(),
			}),
			want: map[string]string{"id": `"my_config"`, "isEnabled": ""},
		},
		{
			name: "dynamic config sends an explicit false, which disables it",
			body: resource_dynamic_config.DynamicConfigToAPIInputModel(ctx, &resource_dynamic_config.DynamicConfigModel{
				IsEnabled: types.BoolValue(false),
			}),
			want: map[string]string{"isEnabled": "false"},
		},
		{
			name: "dynamic config sends an explicit true",
			body: resource_dynamic_config.DynamicConfigToAPIInputModel(ctx, &resource_dynamic_config.DynamicConfigModel{
				IsEnabled: types.BoolValue(true),
			}),
			want: map[string]string{"isEnabled": "true"},
		},
		{
			name: "metric leaves out the tags the config omits",
			body: resource_metric.MetricToAPIInputModel(ctx, &resource_metric.MetricModel{
				Id:   types.StringValue("my_metric"),
				Tags: types.ListUnknown(types.StringType),
			}),
			want: map[string]string{"id": `"my_metric"`, "tags": ""},
		},
		{
			name: "metric sends explicitly empty tags, which clear them",
			body: resource_metric.MetricToAPIInputModel(ctx, &resource_metric.MetricModel{
				Tags: types.ListValueMust(types.StringType, []attr.Value{}),
			}),
			want: map[string]string{"tags": "[]"},
		},
		{
			name: "metric sends configured tags",
			body: resource_metric.MetricToAPIInputModel(ctx, &resource_metric.MetricModel{
				Tags: stringList("checkout"),
			}),
			want: map[string]string{"tags": `["checkout"]`},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			encoded, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			var body map[string]json.RawMessage
			require.NoError(t, json.Unmarshal(encoded, &body))

			for field, want := range testCase.want {
				raw, present := body[field]
				if want == "" {
					assert.False(t, present,
						"%s must be absent so the API leaves the current value alone", field)
					continue
				}

				require.True(t, present, "%s must be sent", field)
				assert.JSONEq(t, want, string(raw))
			}
		})
	}
}

// The rule holds for attributes nobody has written a case for only while every
// request field can express "absent", which is what json.RawMessage plus
// omitempty buys. A field added back as a plain string, bool or slice would
// serialize its zero value again.
func TestRequestModelFieldsCanAllBeOmitted(t *testing.T) {
	models := []any{
		resource_gate.GateAPIInputModel{},
		resource_segment.SegmentAPIInputModel{},
		resource_segment.SegmentRulesAPIInputModel{},
		resource_keys.KeysAPIInputModel{},
		resource_experiment.ExperimentAPIInputModel{},
		resource_dynamic_config.DynamicConfigAPIInputModel{},
		resource_metric.MetricAPIInputModel{},
	}

	rawMessage := reflect.TypeOf(json.RawMessage{})

	for _, model := range models {
		modelType := reflect.TypeOf(model)
		t.Run(modelType.Name(), func(t *testing.T) {
			for i := 0; i < modelType.NumField(); i++ {
				field := modelType.Field(i)
				assert.Equal(t, rawMessage, field.Type,
					"%s must be json.RawMessage so an unspecified attribute can be left out", field.Name)
				assert.True(t, strings.HasSuffix(field.Tag.Get("json"), ",omitempty"),
					"%s must be tagged omitempty so a nil value is left out", field.Name)
			}
		})
	}
}

func stringList(values ...string) types.List {
	elements := make([]attr.Value, len(values))
	for i, value := range values {
		elements[i] = types.StringValue(value)
	}
	return types.ListValueMust(types.StringType, elements)
}

// gateRules builds the smallest list a gate can carry, so the configured-value
// case has something to send.
func gateRules(ctx context.Context) types.List {
	attributeTypes := resource_gate.RulesValue{}.AttributeTypes(ctx)
	rule := resource_gate.NewRulesValueMust(attributeTypes, map[string]attr.Value{
		"base_id":         types.StringNull(),
		"conditions":      types.ListValueMust(resource_gate.ConditionsValue{}.Type(ctx), []attr.Value{}),
		"environments":    types.ListNull(types.StringType),
		"id":              types.StringValue("rule_1"),
		"name":            types.StringValue("everyone"),
		"pass_percentage": types.NumberValue(big.NewFloat(100)),
		"return_value":    types.ObjectNull(resource_gate.ReturnValueValue{}.AttributeTypes(ctx)),
	})

	return types.ListValueMust(resource_gate.RulesValue{}.Type(ctx), []attr.Value{rule})
}
