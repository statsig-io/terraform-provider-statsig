terraform {
  required_providers {
    statsig = {
      version = "~> 1.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_experiment" "simple" {
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
}
