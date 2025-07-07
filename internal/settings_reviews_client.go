package statsig

import (
	"context"
	"terraform-provider-statsig/internal/resource_settings_reviews"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type settingsReviewsClient struct {
	endpoint  string
	transport *Transport
}

func newSettingsReviewsClient(transport *Transport) *settingsReviewsClient {
	return &settingsReviewsClient{
		endpoint:  "settings/reviews",
		transport: transport,
	}
}

func (c *settingsReviewsClient) read(ctx context.Context, settingsReviews *resource_settings_reviews.SettingsReviewsModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_reviews.SettingsReviewsAPIModel
		res, err := c.transport.Get(c.endpoint, "", &data)
		resource_settings_reviews.SettingsReviewsFromAPIModel(ctx, diags, settingsReviews, data)
		return res, err
	})
}

func (c *settingsReviewsClient) update(ctx context.Context, settingsReviews *resource_settings_reviews.SettingsReviewsModel) diag.Diagnostics {
	return runWithDiagnostics(func(diags diag.Diagnostics) (*APIResponse, error) {
		var data resource_settings_reviews.SettingsReviewsAPIModel
		res, err := c.transport.Post(c.endpoint, resource_settings_reviews.SettingsReviewsToAPIModel(ctx, settingsReviews), &data)
		resource_settings_reviews.SettingsReviewsFromAPIModel(ctx, diags, settingsReviews, data)
		return res, err
	})
}
