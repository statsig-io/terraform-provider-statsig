terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_entity_property" "example" {
  name         = "my_entity_property"
  description  = "A short description of this entity property."
  is_read_only = false
  sql          = "SELECT * FROM `shoppy-sales.kenny_dev.users`"
  id_type_mapping = [
    {
      statsig_unit_id = "userID"
      column          = "user_id"
    }
  ]
  timestamp_column = "ts"
  tags             = ["test-tag"]
}

