terraform {
  required_providers {
    statsig = {
      version = "0.1.0"
      source  = "statsig/statsig"
    }
  }
}

variable "gate_id" {
  type    = string
  default = "default_gate_name"
}

data "statsig_gates" "all" {}

# Returns all coffees
output "all_gates" {
  value = data.statsig_gates.all.gates
}

# Only returns packer spiced latte
output "gates" {
  value = data.statsig_gates.all.gates
}
