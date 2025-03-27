resource "statsig_metric" "warehouse_native_metric" {
  description    = "A short description of this metric."
  directionality = "increase"
  is_permanent   = false
  is_read_only   = false
  is_verified    = false
  warehouse_native = {
    metric_source_name = "shoppy_events"
    aggregation        = "count"
    criteria = [
      {
        type      = "metadata",
        column    = "event",
        condition = "=",
        values    = ["add_to_cart"],
      }
    ]
    cuped_attribution_window = 7
    cap                      = 150
  }
  name       = "Warehouse Native Metric"
  tags       = ["test-tag"]
  type       = "user_warehouse"
  unit_types = ["userID"]
}
