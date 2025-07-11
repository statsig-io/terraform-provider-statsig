resource "statsig_settings_project" "example" {
  name               = "Console API Test"
  visibility         = "CLOSED"
  default_unit_type  = "user_id"
}
