terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_tag" "example" {
  name        = "New Tag"
  description = "A short description of what this tag is used for."
  is_core     = false
}
