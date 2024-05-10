resource "statsig_experiment" "full_experiment" {
  name        = "full_experiment"
  description = "A short description of what we are experimenting on."
  id_type     = "userID"
  allocation  = 12.3
  status      = "setup"
  hypothesis  = "Move some metrics"
  layer_id    = "a_layer"
  tags        = [
    "test-tag-a",
    "test-tag-b"
  ]
  primary_metrics_json = jsonencode([
    {
      name : "d1_retention_rate",
      type : "user"
    }
  ])
  primary_metric_tags = [
    "test-tag-a"
  ]
  secondary_metrics_json = jsonencode([
    {
      name : "new_dau",
      type : "user"
    }
  ])
  secondary_metric_tags = [
    "test-tag-b"
  ]
  groups {
    name                  = "Test A"
    size                  = 33.3
    parameter_values_json = jsonencode({ "a_string" : "test_a" })
  }
  groups {
    name                  = "Test B"
    size                  = 33.3
    parameter_values_json = jsonencode({ "a_string" : "test_b" })
  }
  groups {
    name                  = "Control"
    size                  = 33.4
    parameter_values_json = jsonencode({ "a_string" : "control" })
  }
  default_confidence_interval = "80"
  bonferroni_correction       = true
  duration                    = 10
  launched_group_id           = ""
  targeting_gate_id           = "targeting_gate"
}
