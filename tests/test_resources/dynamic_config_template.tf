resource "statsig_dynamic_config" "example_template" {
  id                    = "test_dynamic_config"
  name                  = "test_dynamic_config"
  id_type               = "userID"
  description           = "A test dynamic config"
  is_enabled            = true
  is_template           = true
  rules                 = []
  default_value         = {}
  default_value_json5   = "{}"
}
