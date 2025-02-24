resource "statsig_experiment" "full_experiment" {
  name        = "full_experiment"
  description = "A short description of what we are experimenting on."
  id_type     = "userID"
  allocation  = 12.3
  status      = "setup"
  hypothesis  = "Move some metrics"
  layer_id    = "a_layer"
  tags = [
    "test-tag-a",
    "test-tag-b"
  ]
  primary_metrics = [
    {
      name = "d1_retention_rate"
      type = "user"
    }
  ]
  secondary_metrics = [
    {
      name = "dau"
      type = "user"
    },
    {
      name = "new_dau"
      type = "user"
    }
  ]
  groups = [
    {
      name             = "Test A"
      size             = 33.3
      parameter_values = { a_string = "test_a", a_bool = true }
    },
    {
      name             = "Test B"
      size             = 33.3
      parameter_values = { a_string = "test_b" }
    },
    {
      name             = "Control"
      size             = 33.4
      parameter_values = { a_string = "control" }
    }
  ]
  default_confidence_interval = "80"
  bonferroni_correction       = true
  duration                    = 10
  targeting_gate_id           = "targeting_gate"
  lifecycle {
    ignore_changes = [
      "primary_metric_tags", # Metric tags cannot actually be assigned. Associated metrics are exploded from the tags and set on primaryMetrics/secondaryMetrics
      "secondary_metric_tags",
    ]
  }
}
