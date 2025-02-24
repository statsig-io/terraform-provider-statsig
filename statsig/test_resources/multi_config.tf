resource "statsig_gate" "simple_gate_a" {
  name        = "simple_gate_a"
  description = "A short description of what this Gate is used for."
  is_enabled  = true
  id_type     = "userID"
  rules {
    name            = "All Conditions"
    pass_percentage = 10
    conditions {
      type = "public"
    }
  }
}

resource "statsig_gate" "simple_gate_b" {
  name        = "simple_gate_b"
  description = "A short description of what this Gate is used for."
  is_enabled  = true
  id_type     = "userID"
}


resource "statsig_experiment" "simple_experiment" {
  name        = "simple_experiment"
  description = "A short description of what this Experiment is used for."
  id_type     = "userID"
  allocation  = 10
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
  lifecycle {
    ignore_changes = ["secondary_metrics_json"] # Automatically attached core tag
  }
}

output "gates" {
  value = [
    statsig_gate.simple_gate_a,
    statsig_gate.simple_gate_b
  ]
}

output "experiments" {
  value = [
    statsig_experiment.simple_experiment
  ]
}
