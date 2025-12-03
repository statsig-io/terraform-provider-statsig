terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_role" "example" {
  name = "Custom Role"
  permissions = {
    invitation_access                                = false
    create_configs                                   = false
    edit_or_delete_configs                           = false
    launch_to_production                             = false
    launch_or_disable_configs                        = false
    start_experiments                                = false
    create_or_edit_templates                         = false
    create_or_edit_dashboards                        = false
    create_teams                                     = false
    edit_dynamic_config_schemas                      = false
    create_release_pipelines                         = false
    approve_required_review_release_pipeline_phase   = false
    self_approve_review                              = false
    approve_reviews                                  = false
    bypass_reviews_for_overrides                     = false
    bypass_precommit_webhook                         = false
    metric_management                                = false
    verify_metrics                                   = false
    use_metrics_explorer                             = false
    local_metrics                                    = false
    manage_alerts                                    = false
    integrations_edit_access                         = false
    source_connection_and_creation                   = false
    data_warehouse_ingestion_and_exports_edit_access = false
    edit_and_tag_configs_with_core_tag               = false
    reset_experiments                                = false
    event_dimensions_access                          = false
    manual_whn_reload                                = false
  }
}
