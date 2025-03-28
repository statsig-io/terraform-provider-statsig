terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_metric" "custom_event_count_metric" {
  custom_roll_up_end   = 14
  custom_roll_up_start = 0
  description          = "A short description of this metric."
  directionality       = "decrease"
  is_permanent         = false
  is_read_only         = false
  is_verified          = false
  metric_events = [
    {
      criteria = []
      name     = "test_event_1"
    }
  ]
  name               = "Custom Event Count Metric"
  rollup_time_window = "custom"
  tags               = ["test-tag"]
  type               = "event_count_custom"
  unit_types         = ["userID"]
}

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
