resource "statsig_dynamic_config" "my_dynamic_config" {
  id          = "my_dynamic_config"
  name        = "my_dynamic_config"
  description = "A short description of what this Dynamic Config is used for."
  is_enabled  = true
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
      return_value = {
        extra_field = 12
        my_field    = "My Other Value"
      }
      return_value_json5 = "{extra_field: \"12\",\n  my_field: \"My Other Value\"}"
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
      return_value = {
        my_field = "My Other Value"
      }
      return_value_json5 = "{\"my_field\":\"My Other Value\"}"
    }
  ]
  default_value = {
    my_field = "My Value"
  }
  default_value_json5 = "{\"my_field\":\"My Value\"}"
}
