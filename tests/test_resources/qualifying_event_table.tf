resource "statsig_qualifying_event" "table" {
  name         = "table_qualifying_event"
  description  = "A short description of this qualifying event."
  is_read_only = false
  source_type  = "table"
  table_name   = "`shoppy-sales.kenny_dev.shoppy-events`"
  sql          = ""
  custom_field_mapping = [
    {
      formula = "price_usd/100"
      key     = "price_usd_cents"
    }
  ]
  id_type_mapping = [
    {
      statsig_unit_id = "userID"
      column          = "user_id"
    }
  ]
  timestamp_column = "ts"
  tags             = ["test-tag"]
}
