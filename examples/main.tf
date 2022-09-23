terraform {
  required_providers {
    statsig = {
      version = "0.1.0"
      source  = "statsig/statsig"
    }
  }
}

provider "statsig" {}

module "psl" {
  source  = "./gate"
  gate_id = "a_gate"
}

output "psl" {
  value = module.psl.gates
}
