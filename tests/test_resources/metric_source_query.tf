resource "statsig_metric_source" "query" {
  name         = "query_metric_source"
  description  = "A short description of this metric source."
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
