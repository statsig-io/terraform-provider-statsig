package resource_unit_id_type

import (
	"context"
	"terraform-provider-statsig/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// API data model for UnitIdTypeModel
type UnitIdTypeAPIModel struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func UnitIdTypeToAPIModel(ctx context.Context, unitIdType *UnitIdTypeModel) UnitIdTypeAPIModel {
	return UnitIdTypeAPIModel{
		Name:        utils.StringFromNilableValue(unitIdType.Name),
		Description: utils.StringFromNilableValue(unitIdType.Description),
	}
}

func UnitIdTypeFromAPIModel(ctx context.Context, diags diag.Diagnostics, unitIdType *UnitIdTypeModel, res UnitIdTypeAPIModel) {
	unitIdType.Name = utils.StringToStringValue(res.Name)
	unitIdType.Description = utils.StringToNilableValue(res.Description)
}
