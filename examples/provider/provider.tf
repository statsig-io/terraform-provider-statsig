terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

# Configure the Statsig provider
provider "statsig" {
  console_api_key = var.statsig_api_key
}

# Create a new gate
resource "statsig_gate" "example_gate" {
  # ...
}
