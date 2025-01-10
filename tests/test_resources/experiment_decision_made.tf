variable "test_group_id" {
  type = string
}

variable "control_group_id" {
  type = string
}

resource "statsig_experiment" "my_experiment" {
  id                = "my_experiment"
  name              = "my_experiment"
  description       = "A short description of what we are experimenting on."
  id_type           = "userID"
  allocation        = 100
  status            = "decision_made"
  launched_group_id = var.test_group_id
  groups = [
    {
      id               = var.test_group_id
      name             = "Test Group"
      size             = 50
      parameter_values = { "a_string" : "test_string", "a_bool" : true }
    },
    {
      id               = var.control_group_id
      name             = "Control Group"
      size             = 50
      parameter_values = { "a_string" : "control_string", "a_bool" : false }
    }
  ]
}
