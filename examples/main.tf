terraform {
  required_providers {
    statsig = {
      version = "0.1.2"
      source  = "statsig-io/statsig"
    }
  }
}

module "gates" {
  source = "./gates"
}

output "gates" {
  value = module.gates
}

module "experiments" {
  source = "./experiments"
}

output "experiments" {
  value = module.experiments
}
