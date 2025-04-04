resource "statsig_qualifying_event" "query" {
  name         = "query_qualifying_event"
  description  = "A short description of this qualifying event."
  is_read_only = false
  source_type  = "query"
  sql          = "SELECT * FROM `shoppy-sales.kenny_dev.shoppy-events`"
  id_type_mapping = [
    {
      statsig_unit_id = "userID"
      column          = "user_id"
    }
  ]
  timestamp_column = "ts"
  tags             = ["test-tag"]
}
