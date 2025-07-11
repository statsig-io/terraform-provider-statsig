terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_unit_id_type" "example" {
  name        = "my_unit_id_type"
  description = "An example unit ID type"
}
