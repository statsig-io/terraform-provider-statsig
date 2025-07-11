package resource_settings_reviews

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/statsig-io/terraform-provider-statsig/internal/utils"
)

// API data model for SettingsReviewsModel
type SettingsReviewsAPIModel struct {
	IsConfigReviewRequired               bool `json:"is_config_review_required"`
	IsMetricReviewRequired               bool `json:"is_metric_review_required"`
	IsMetricReviewRequiredOnVerifiedOnly bool `json:"is_metric_review_required_on_verified_only"`
	IsWhnAnalysisOnlyReviewRequired      bool `json:"is_whn_analysis_only_review_required"`
	IsWhnSourceReviewRequired            bool `json:"is_whn_source_review_required"`
}

func SettingsReviewsToAPIModel(ctx context.Context, settingsReviews *SettingsReviewsModel) SettingsReviewsAPIModel {
	return SettingsReviewsAPIModel{
		IsConfigReviewRequired:               utils.BoolFromBoolValue(settingsReviews.IsConfigReviewRequired),
		IsMetricReviewRequired:               utils.BoolFromBoolValue(settingsReviews.IsMetricReviewRequired),
		IsMetricReviewRequiredOnVerifiedOnly: utils.BoolFromBoolValue(settingsReviews.IsMetricReviewRequiredOnVerifiedOnly),
		IsWhnAnalysisOnlyReviewRequired:      utils.BoolFromBoolValue(settingsReviews.IsWhnAnalysisOnlyReviewRequired),
		IsWhnSourceReviewRequired:            utils.BoolFromBoolValue(settingsReviews.IsWhnSourceReviewRequired),
	}
}

func SettingsReviewsFromAPIModel(ctx context.Context, diags diag.Diagnostics, settingsReviews *SettingsReviewsModel, res SettingsReviewsAPIModel) {
	settingsReviews.IsConfigReviewRequired = utils.BoolToBoolValue(res.IsConfigReviewRequired)
	settingsReviews.IsMetricReviewRequired = utils.BoolToBoolValue(res.IsMetricReviewRequired)
	settingsReviews.IsMetricReviewRequiredOnVerifiedOnly = utils.BoolToBoolValue(res.IsMetricReviewRequiredOnVerifiedOnly)
	settingsReviews.IsWhnAnalysisOnlyReviewRequired = utils.BoolToBoolValue(res.IsWhnAnalysisOnlyReviewRequired)
	settingsReviews.IsWhnSourceReviewRequired = utils.BoolToBoolValue(res.IsWhnSourceReviewRequired)
}
