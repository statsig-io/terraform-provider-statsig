terraform {
  required_providers {
    statsig = {
      version = "~> 0.1.2"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_experiment" "my_experiment" {
  name        = "my_experiment"
  description = "A short description of what we are experimenting on."
  id_type     = "userID"
  allocation  = 12.3
  status      = "setup"
  hypothesis  = "Move some metrics"
  layer_id    = "a_layer"
  tags        = []
  primary_metrics_json   = "[]"
  primary_metric_tags    = []
  secondary_metrics_json = "[]"
  secondary_metric_tags  = []
  groups {
    name                  = "Test"
    size                  = 50
    parameter_values_json = jsonencode({ "a_string" : "test" })
  }
  groups {
    name                  = "Control"
    size                  = 50
    parameter_values_json = jsonencode({ "a_string" : "control" })
  }
  default_confidence_interval = "80"
  bonferroni_correction       = true
  duration                    = 10
  launched_group_id           = ""
  targeting_gate_id           = "my_gate"
}

output "my_experiment" {
  value = statsig_experiment.my_experiment
}
