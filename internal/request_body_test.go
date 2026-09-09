package statsig

import (
	"context"
	"encoding/json"
	"math/big"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
				Rules:       gateRules(ctx, types.StringValue("rule_1")),
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
			body := requestBody(t, testCase.body)

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

// The rule reaches inside a nested attribute too. Terraform marks a nested
// Optional+Computed attribute unknown when the configuration does not mention
// it, exactly as it does a top-level one, so a rule id or a warehouse native
// column the config never wrote used to go out as a Go zero value and clear
// whatever the Statsig Console held.
func TestRequestBodyOmitsUnspecifiedNestedAttributes(t *testing.T) {
	ctx := context.Background()

	t.Run("a gate rule leaves out the id the config omits", func(t *testing.T) {
		body := requestBody(t, resource_gate.GateToAPIInputModel(ctx, &resource_gate.GateModel{
			Rules: gateRules(ctx, types.StringUnknown()),
		}))

		assert.JSONEq(t, `[{"name":"everyone","passPercentage":100,"conditions":[]}]`,
			string(body["rules"]),
			"an empty rule id renames the rule the Console API matches on")
	})

	t.Run("warehouse native leaves out the columns the config omits", func(t *testing.T) {
		body := requestBody(t, resource_metric.MetricToAPIInputModel(ctx, &resource_metric.MetricModel{
			WarehouseNative: warehouseNative(t, ctx, map[string]attr.Value{
				"metric_source_name": types.StringValue("shoppy_events"),
				"aggregation":        types.StringValue("count"),
			}),
		}))

		var warehouseNativeBody map[string]json.RawMessage
		require.NoError(t, json.Unmarshal(body["warehouseNative"], &warehouseNativeBody))

		assert.JSONEq(t, `"shoppy_events"`, string(warehouseNativeBody["metricSourceName"]))
		assert.JSONEq(t, `"count"`, string(warehouseNativeBody["aggregation"]))
		for _, field := range []string{"metricDimensionColumns", "criteria", "denominatorCriteria"} {
			assert.NotContains(t, warehouseNativeBody, field,
				"%s must be absent so the Console API keeps the current value", field)
		}
	})
}

// requestModels lists every type the provider marshals into a Console API
// request body. TestRequestModelFieldsCanAllBeOmitted proves the nested part of
// the list is complete rather than trusting it.
func requestModels() []any {
	return []any{
		resource_gate.GateAPIInputModel{},
		resource_gate.MonitoringMetricAPIInputModel{},
		resource_gate.RuleAPIInputModel{},
		resource_gate.ConditionAPIInputModel{},
		resource_segment.SegmentAPIInputModel{},
		resource_segment.SegmentRulesAPIInputModel{},
		resource_segment.RuleAPIInputModel{},
		resource_segment.ConditionAPIInputModel{},
		resource_keys.KeysAPIInputModel{},
		resource_experiment.ExperimentAPIInputModel{},
		resource_experiment.GroupAPIInputModel{},
		resource_experiment.ExperimentMetricAPIInputModel{},
		resource_experiment.LinkAPIInputModel{},
		resource_dynamic_config.DynamicConfigAPIInputModel{},
		resource_dynamic_config.RuleAPIInputModel{},
		resource_dynamic_config.ConditionAPIInputModel{},
		resource_metric.MetricAPIInputModel{},
		resource_metric.WarehouseNativeAPIInputModel{},
		resource_metric.MetricEventAPIInputModel{},
		resource_metric.CriteriaAPIInputModel{},
		resource_metric.MetricComponentMetricAPIInputModel{},
		resource_metric.FunnelEventAPIInputModel{},
		resource_metric.WarehouseNativeFunnelEventAPIInputModel{},
	}
}

// responseModels are the models a Console API response decodes into. Each one
// mirrors a request model field for field, in real Go types, which is what
// makes them the map of the request shape: a request model's json.RawMessage
// erases the nested types, so nothing can be reached through one.
//
// statsig_keys is left out on purpose. Its response carries display names where
// the request carries IDs, so the two models do not mirror each other. See
// resource_keys.KeysAPIOutputModel.
func responseModels() []any {
	return []any{
		resource_gate.GateAPIModel{},
		resource_segment.SegmentAPIModel{},
		resource_experiment.ExperimentAPIModel{},
		resource_dynamic_config.DynamicConfigAPIModel{},
		resource_metric.MetricAPIModel{},
	}
}

// The rule holds for attributes nobody has written a case for only while every
// request field can express "absent", which is what json.RawMessage plus
// omitempty buys. A field added back as a plain string, bool or slice would
// serialize its zero value again.
func TestRequestModelFieldsCanAllBeOmitted(t *testing.T) {
	registry := map[string]reflect.Type{}
	for _, model := range requestModels() {
		modelType := reflect.TypeOf(model)
		registry[typeKey(modelType)] = modelType
	}

	for _, model := range requestModels() {
		modelType := reflect.TypeOf(model)
		t.Run(typeKey(modelType), func(t *testing.T) {
			assertOmittable(t, modelType, typeKey(modelType))
		})
	}

	t.Run("every model a request body reaches is registered", func(t *testing.T) {
		for _, model := range responseModels() {
			for _, response := range structsReachableFrom(reflect.TypeOf(model)) {
				request, found := registry[requestKey(response)]
				if !assert.True(t, found,
					"%s is part of a request body, so it needs a %s whose fields go through utils.APIField",
					typeKey(response), requestName(response)) {
					continue
				}

				assert.Equal(t, jsonFieldNames(response), jsonFieldNames(request),
					"%s must carry the same fields as %s", typeKey(request), typeKey(response))
			}
		}
	})
}

var rawMessage = reflect.TypeOf(json.RawMessage{})

// assertOmittable fails unless every field the encoder can reach is a
// json.RawMessage tagged omitempty. It fails on a shape it cannot traverse
// rather than skipping it, so a field of some new kind is a failure here and
// not a field nobody is checking.
func assertOmittable(t *testing.T, modelType reflect.Type, path string) {
	t.Helper()

	if modelType == rawMessage {
		return
	}

	switch modelType.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Array, reflect.Map:
		assertOmittable(t, modelType.Elem(), path+"[]")
	case reflect.Struct:
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)
			fieldPath := path + "." + field.Name

			if !assert.Equal(t, rawMessage, field.Type,
				"%s must be json.RawMessage so an unspecified attribute can be left out", fieldPath) {
				continue
			}
			assert.True(t, strings.HasSuffix(field.Tag.Get("json"), ",omitempty"),
				"%s must be tagged omitempty so a nil value is left out", fieldPath)
		}
	default:
		t.Fatalf("%s is a %s, a shape this test cannot traverse: teach it the shape rather than leaving the field unchecked",
			path, modelType.Kind())
	}
}

// structsReachableFrom returns the root and every struct type the encoder can
// reach from it, following pointers, slices, arrays and maps.
func structsReachableFrom(root reflect.Type) []reflect.Type {
	var found []reflect.Type
	seen := map[reflect.Type]bool{}

	var walk func(reflect.Type)
	walk = func(modelType reflect.Type) {
		switch modelType.Kind() {
		case reflect.Pointer, reflect.Slice, reflect.Array, reflect.Map:
			walk(modelType.Elem())
		case reflect.Struct:
			if seen[modelType] {
				return
			}
			seen[modelType] = true
			found = append(found, modelType)
			for i := 0; i < modelType.NumField(); i++ {
				walk(modelType.Field(i).Type)
			}
		}
	}

	walk(root)
	return found
}

func typeKey(modelType reflect.Type) string {
	return path.Base(modelType.PkgPath()) + "." + modelType.Name()
}

func requestName(response reflect.Type) string {
	return strings.TrimSuffix(response.Name(), "APIModel") + "APIInputModel"
}

func requestKey(response reflect.Type) string {
	return path.Base(response.PkgPath()) + "." + requestName(response)
}

func jsonFieldNames(modelType reflect.Type) []string {
	names := make([]string, 0, modelType.NumField())
	for i := 0; i < modelType.NumField(); i++ {
		names = append(names, strings.Split(modelType.Field(i).Tag.Get("json"), ",")[0])
	}
	sort.Strings(names)
	return names
}

func requestBody(t *testing.T, model any) map[string]json.RawMessage {
	t.Helper()

	encoded, err := json.Marshal(model)
	require.NoError(t, err)

	var body map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(encoded, &body))
	return body
}

func stringList(values ...string) types.List {
	elements := make([]attr.Value, len(values))
	for i, value := range values {
		elements[i] = types.StringValue(value)
	}
	return types.ListValueMust(types.StringType, elements)
}

// gateRules builds the smallest list a gate can carry. The caller supplies the
// rule id so a case can set it or leave it unspecified.
func gateRules(ctx context.Context, id types.String) types.List {
	attributeTypes := resource_gate.RulesValue{}.AttributeTypes(ctx)
	rule := resource_gate.NewRulesValueMust(attributeTypes, map[string]attr.Value{
		"base_id":         types.StringNull(),
		"conditions":      types.ListValueMust(resource_gate.ConditionsValue{}.Type(ctx), []attr.Value{}),
		"environments":    types.ListNull(types.StringType),
		"id":              id,
		"name":            types.StringValue("everyone"),
		"pass_percentage": types.NumberValue(big.NewFloat(100)),
		"return_value":    types.ObjectNull(resource_gate.ReturnValueValue{}.AttributeTypes(ctx)),
	})

	return types.ListValueMust(resource_gate.RulesValue{}.Type(ctx), []attr.Value{rule})
}

// warehouseNative builds a warehouse native block where only the named
// attributes are set. Everything else is unknown, the state Terraform leaves an
// Optional+Computed attribute the configuration never mentions.
func warehouseNative(t *testing.T, ctx context.Context, set map[string]attr.Value) resource_metric.WarehouseNativeValue {
	t.Helper()

	attributeTypes := resource_metric.WarehouseNativeValue{}.AttributeTypes(ctx)
	attributes := make(map[string]attr.Value, len(attributeTypes))
	for name, attributeType := range attributeTypes {
		value, ok := set[name]
		if !ok {
			var err error
			value, err = attributeType.ValueFromTerraform(ctx,
				tftypes.NewValue(attributeType.TerraformType(ctx), tftypes.UnknownValue))
			require.NoError(t, err)
		}
		attributes[name] = value
	}

	return resource_metric.NewWarehouseNativeValueMust(attributeTypes, attributes)
}
