terraform {
  required_providers {
    statsig = {
      version = "0.2.0"
      source  = "statsig-io/statsig"
    }
  }
}

module "gates" {
  source = "./gates"
}
