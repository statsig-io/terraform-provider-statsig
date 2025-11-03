package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccSettingsReviews(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/settings_reviews.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifySettingsReviewsSetup(t, "statsig_settings_reviews.example"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/settings_reviews.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifySettingsReviewsSetup(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "false", local["is_config_review_required"])
		assert.Equal(t, "true", local["is_metric_review_required"])
		assert.Equal(t, "true", local["is_metric_review_required_on_verified_only"])

		return nil
	}
}

func TestAccSettingsReviewsWHN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviders(t, TestOptions{isWHN: true}),
		Steps: []resource.TestStep{
			{
				ConfigFile: config.StaticFile("test_resources/settings_reviews_whn.tf"),
				Check: resource.ComposeTestCheckFunc(
					verifySettingsReviewsSetupWHN(t, "statsig_settings_reviews.example_whn"),
				),
			},
			{
				ConfigFile: config.StaticFile("test_resources/settings_reviews_whn.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func verifySettingsReviewsSetupWHN(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, _ := s.RootModule().Resources[name]
		local := rs.Primary.Attributes

		assert.Equal(t, "false", local["is_config_review_required"])
		assert.Equal(t, "true", local["is_metric_review_required"])
		assert.Equal(t, "true", local["is_metric_review_required_on_verified_only"])
		assert.Equal(t, "true", local["is_whn_analysis_only_review_required"])
		assert.Equal(t, "true", local["is_whn_source_review_required"])

		return nil
	}
}
