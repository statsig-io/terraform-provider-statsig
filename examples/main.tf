terraform {
  required_providers {
    statsig = {
      version = "2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

module "gates" {
  source = "./gates"
}
