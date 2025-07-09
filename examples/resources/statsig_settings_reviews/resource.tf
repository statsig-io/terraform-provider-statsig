terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_settings_reviews" "main" {
  is_config_review_required                   = true
  is_metric_review_required                   = true
  is_metric_review_required_on_verified_only  = true
  is_whn_analysis_only_review_required        = true
  is_whn_source_review_required               = true
}
