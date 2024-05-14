# Terraform Provider Statsig

The Statsig Terrafrom Provider allows you to configure your Statsig project entities (e.g. gates, experiments, keys) with Terraform. This is achieved by leveraging the Console API. If there is something you need to perform that isn't supported by the Terraform Provider, checkout the [Console API](https://docs.statsig.com/console-api/introduction).

It is hosted on the Terraform registry at https://registry.terraform.io/providers/statsig-io/statsig

## Gates

You can create a .tf file (Terraform File) to configure how your gate works. All features of [console/v1/gates](https://docs.statsig.com/console-api/gates) are supported. The layout is very similar to the JSON body of a /gates request.

Requiring the Statsig provider. (You will need to change the version).

```go
terraform {
  required_providers {
    statsig = {
      version = "x.x.x"
      source  = "statsig-io/statsig"
    }
  }
}
```

Creating a basic gate resource

```go
resource "statsig_gate" "my_gate" {
  name        = "my_gate"
  description = "A short description of what this Gate is used for."
  is_enabled  = true
  id_type     = "userID"
  rules {
    name            = "Public"
    pass_percentage = 100
    conditions {
      type = "public"
    }
  }
}
```

### Conditions

All Console API conditions are supported but the syntax needs a little tweaking.


| Attributes | Description | Type | Required | Default |
| ---------- | ----------- | ---- | -------- | ------- |
| type | The [type](https://docs.statsig.com/console-api/rules#all-conditions) of condition it is. | string | `true` | - |
| operator | What form of evaluation should be run against the **target_value**. | string | `false` | "" |
| target_value | The value or values you wish to evaluate. Note: This must be an array, and elements should be of string type. (You can put quotes on numbers. 31 -> "31") | []string | `true` | [] |
| field | Only for custom_field condition type. The name of the field you wish to pull for evaluation from the "custom" object on a user. | string | `false` | "" |

```go
conditions {
  type         = "custom_field"
  target_value = ["31"]
  operator     = "gt"
  field        = "age"
}
```

See the full list of conditions [here](https://docs.statsig.com/console-api/rules#all-conditions).

## Keys

The `statsig_keys` resource allows you to manage SDK keys and Console API keys. Reference to [Keys API docs](https://docs.statsig.com/console-api/keys)

| Attributes | Description | Type | Required | Default |
| ---------- | ----------- | ---- | -------- | ------- |
| type | The type of key. | string | `true` | - |
| description | The description of the key. | string | `true` | - |
| key | The actual value of the key. | string | `false` | (computed) |
| scopes | The list of scopes granted to the key. | []string | `false` | [] |
| environments | The list of environments accessible by the key. | []string | `false` | [] |
| target_app_id | The primary target app assigned to the key. | string | `false` | "" |
| secondary_target_app_ids | The list of secondary target apps assigned to the key. | []string | `false` | [] |

### Example
```go
resource "statsig_keys" "server_key" {
  description  = "A short description of what this server key is used for."
  type         = "SERVER"
  environments = ["production"]
}

resource "statsig_keys" "client_key" {
  description  = "A short description of what this client key is used for."
  type         = "CLIENT"
  environments = ["production"]
  scopes       = ["client_download_config_specs"]
}

resource "statsig_keys" "console_key" {
  description = "A short description of what this console key is used for."
  type        = "CONSOLE"
  scopes      = ["omni_read_only"]
}
```
