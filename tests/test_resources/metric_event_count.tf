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
