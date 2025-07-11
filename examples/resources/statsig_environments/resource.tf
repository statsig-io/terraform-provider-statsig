terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_environments" "main" {
  environments = [
    {
      name                      = "development"
      id                        = "0.35629335902367476"
      is_production             = false
      requires_review           = false
      requires_release_pipeline = true
    },
    {
      name                      = "staging"
      id                        = "0.015089163460661137"
      is_production             = false
      requires_review           = false
      requires_release_pipeline = true
    },
    {
      name                      = "production"
      id                        = "0.4067426155658289"
      is_production             = true
      requires_review           = true
      requires_release_pipeline = true
    }
  ]
}
