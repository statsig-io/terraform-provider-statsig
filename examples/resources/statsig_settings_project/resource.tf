terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_settings_project" "main" {
  name               = "My Statsig Project"
  visibility         = "CLOSED"
  default_unit_type  = "user_id"
}
