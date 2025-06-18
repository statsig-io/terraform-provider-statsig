package resource_tag

import (
	"context"
	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type TagAPIModel struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsCore      bool   `json:"isCore"`
}

func TagToAPIModel(ctx context.Context, tag *TagModel) TagAPIModel {
	return TagAPIModel{
		Name:        tag.Name.ValueString(),
		Description: tag.Description.ValueString(),
		IsCore:      utils.BoolFromBoolValue(tag.IsCore),
	}
}

func TagFromAPIModel(ctx context.Context, diags diag.Diagnostics, tag *TagModel, res TagAPIModel) {
	tag.Name = utils.StringToNilableValue(res.Name)
	tag.Description = utils.StringToNilableValue(res.Description)
	tag.IsCore = utils.BoolToBoolValue(res.IsCore)
}
