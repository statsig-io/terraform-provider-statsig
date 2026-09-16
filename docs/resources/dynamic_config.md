---
page_title: "statsig_dynamic_config Resource - terraform-provider-statsig"
subcategory: ""
description: |-
  
---

# statsig_dynamic_config (Resource)

This resource allows you to configure Dynamic Configs for your Statsig project.

To learn more about Dynamic Configs, see [Dynamic Configs Documentation](https://docs.statsig.com/dynamic-config)

To learn more about the API powering this resource, see [Dynamic Configs API Documentation](https://docs.statsig.com/console-api/all-endpoints-generated#post-/console/v1/dynamic_configs)

## Example Usage

### Basic

```terraform
resource "statsig_dynamic_config" "basic" {
  id          = "my_dynamic_config"
  name        = "my_dynamic_config"
  description = "A short description of what this Dynamic Config is used for."
  is_enabled  = true
  id_type     = "userID"
  rules       = []
  default_value       = jsonencode({ my_field = "My Value" })
  default_value_json5 = jsonencode({ my_field = "My Value" })
}
```

### With Nested Objects

`default_value` and `return_value` accept any valid JSON via `jsonencode()`, including deeply nested objects.

```terraform
locals {
  premium_config = {
    ui = {
      theme = {
        primary   = "#6200EE"
        secondary = "#03DAC5"
      }
      layout = "grid"
    }
    limits = {
      max_projects = 50
      max_members  = 200
    }
    features = ["analytics", "export", "api_access"]
  }

  free_config = {
    ui = {
      theme = {
        primary   = "#333333"
        secondary = "#666666"
      }
      layout = "list"
    }
    limits = {
      max_projects = 3
      max_members  = 5
    }
    features = ["analytics"]
  }
}

resource "statsig_dynamic_config" "app_config" {
  id          = "app-config"
  name        = "App Config"
  description = "Application configuration per user tier."
  is_enabled  = true
  id_type     = "userID"
  rules = [
    {
      name            = "Premium Users"
      pass_percentage = 100
      conditions = [
        {
          type         = "custom_field"
          field        = "plan"
          target_value = ["premium"]
          operator     = "any"
        }
      ]
      return_value       = jsonencode(local.premium_config)
      return_value_json5 = jsonencode(local.premium_config)
    },
  ]
  default_value       = jsonencode(local.free_config)
  default_value_json5 = jsonencode(local.free_config)
}
```

### Full Example

```terraform
terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_dynamic_config" "full" {
  id          = "my_dynamic_config"
  name        = "my_dynamic_config"
  description = "A short description of what this Dynamic Config is used for."
  is_enabled  = true
  is_template = true
  id_type     = "userID"
  rules = [
    {
      name            = "All Conditions"
      pass_percentage = 10
      environments    = ["production"]
      conditions = [
        {
          type         = "public"
          target_value = []
        },
        {
          type = "user_id"
          target_value = [
            "1", "2"
          ]
          operator = "any"
        },
        {
          type         = "email"
          target_value = ["@outlook.com", "@gmail.com"]
          operator     = "str_contains_any"
        },
        {
          type         = "custom_field"
          target_value = [31]
          operator     = "gt"
          field        = "age"
        },
        {
          type         = "app_version"
          target_value = ["1.1.1"]
          operator     = "version_gt"
        },
        {
          type         = "browser_name"
          target_value = ["Firefox", "Chrome"]
          operator     = "any"
        },
        {
          type         = "browser_version"
          target_value = ["94.0.4606.81", "94.0.4606.92"]
          operator     = "any"
        },
        {
          type         = "os_name"
          target_value = ["Android", "Windows"]
          operator     = "none"
        },
        {
          type         = "os_version"
          target_value = ["11.0.0"]
          operator     = "version_lte"
        },
        {
          type         = "country"
          target_value = ["NZ", "US"]
          operator     = "any"
        },
        {
          type         = "passes_gate"
          target_value = ["my_gate_2"]
        },
        {
          type         = "fails_gate"
          target_value = ["a_failing_gate"]
        },
        {
          type         = "time"
          target_value = [1643070357193]
          operator     = "after"
        },
        {
          type         = "environment_tier"
          target_value = ["production"]
          operator     = "any"
        },
        {
          type         = "passes_segment"
          target_value = ["growth_org"]
        },
        {
          type         = "fails_segment"
          target_value = ["promo_id_list"]
        },
        {
          type         = "ip_address"
          target_value = ["1.1.1.1", "8.8.8.8"]
          operator     = "any"
        }
      ]
      return_value       = jsonencode({ extra_field = 12, my_field = "My Other Value" })
      return_value_json5 = jsonencode({ extra_field = 12, my_field = "My Other Value" })
    },
    {
      name            = "Development Conditions"
      pass_percentage = 10
      environments = ["development"]
      conditions = [
        {
          type         = "public"
          target_value = []
        }
      ]
      return_value       = jsonencode({ my_field = "My Other Value" })
      return_value_json5 = jsonencode({ my_field = "My Other Value" })
    }
  ]
  default_value       = jsonencode({ my_field = "My Value" })
  default_value_json5 = jsonencode({ my_field = "My Value" })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The dynamic config display name

### Optional

- `creator_email` (String)
- `creator_id` (String)
- `default_value` (String) The fallback JSON object when no rules are triggered. Use `jsonencode()` to pass a value. Supports arbitrary nested objects.
- `default_value_json5` (String) Can include comments. If provided with defaultValue, must parse to the same JSON
- `description` (String)
- `id` (String) The dynamic config name ID
- `id_type` (String) The type of ID which the dynamic config is based on.
- `is_enabled` (Boolean) Is the dynamic config enabled
- `is_template` (Boolean)
- `rules` (Attributes List) An array of Rule objects (see [below for nested schema](#nestedatt--rules))
- `schema` (String) A schema using JSON Schema Draft 2020-12 to enforce return values of this dynamic config's rules.
- `schema_json5` (String) `schema` except with Json5 comments. Optional and should parse to same json as `schema`.
- `tags` (List of String) The list of tag names attached to the dynamic config
- `target_apps` (List of String)
- `team` (String)

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `conditions` (Attributes List) An array of Condition objects. (see [below for nested schema](#nestedatt--rules--conditions))
- `name` (String) The name of this rule.
- `pass_percentage` (Number) Of the users that meet the conditions of this rule, what percent should return true.

Optional:

- `base_id` (String) The base ID of this rule, i.e. without any added metadata. Will remain the exact same throughout
- `environments` (List of String)
- `id` (String) The Statsig ID of this rule.
- `return_value` (String) The return value JSON object for this rule. Use `jsonencode()` to pass a value. Supports arbitrary nested objects.
- `return_value_json5` (String)

<a id="nestedatt--rules--conditions"></a>
### Nested Schema for `rules.conditions`

Required:

- `type` (String)

Optional:

- `custom_id` (String)
- `field` (String)
- `operator` (String)
- `target_value` (List of String)
