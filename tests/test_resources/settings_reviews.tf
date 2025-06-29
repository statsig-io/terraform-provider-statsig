resource "statsig_settings_reviews" "example" {
  is_config_review_required                   = true
  is_metric_review_required                   = true
  is_metric_review_required_on_verified_only  = true
}
