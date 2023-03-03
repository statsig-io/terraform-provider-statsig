variable "launched_group_id" {
  type = string
}

resource "statsig_experiment" "my_experiment" {
  name              = "my_experiment"
  description       = "A short description of what we are experimenting on."
  id_type           = "userID"
  allocation        = 0
  status            = "decision_made"
  launched_group_id = var.launched_group_id
  groups {
    name                  = "Test Group"
    size                  = 50
    parameter_values_json = jsonencode({ "a_string" : "test_string", "a_bool" : true })
  }
  groups {
    name                  = "Control Group"
    size                  = 50
    parameter_values_json = jsonencode({ "a_string" : "control_string", "a_bool" : false })
  }
}
