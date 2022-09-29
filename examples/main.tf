terraform {
  required_providers {
    statsig = {
      version = "0.1.0"
      source  = "statsig/statsig"
    }
  }
}

#provider "statsig" {}

resource "statsig_gate" "my_gate" {
  name = "my_gate"
  description = "This is my gate"
  is_enabled = true
  id_type = "userID"
}

#module "psl" {
#  source  = "./gate"
#  gate_id = "a_gate"
#}

#output "psl" {
#  value = module.psl.gates
#}

output "my_gate_gate" {
  value = statsig_gate.my_gate
}
