resource "statsig_experiment" "my_experiment" {
  name        = "my_experiment"
  description = "A short description of what we are experimenting on."
  id_type     = "userID"
  allocation  = 10
  groups = [
    {
      name             = "Test Group"
      size             = 50
      parameter_values = { "a_string" : "test_string", "a_bool" : true }
    },
    {
      name             = "Control Group"
      size             = 50
      parameter_values = { "a_string" : "control_string", "a_bool" : false }
    }
  ]
}
