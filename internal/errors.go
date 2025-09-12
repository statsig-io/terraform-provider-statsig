package statsig

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternalAPIErrorDiagnostic(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Internal API Error",
		"While calling the API, an unexpected error occurred. "+
			"Please contact support if you are unsure how to resolve the error.\n\n"+
			"Error: "+err.Error(),
	)
}

func APIErrorDiagnostic(res *APIResponse) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		fmt.Sprintf("API Error (%d): %s", res.StatusCode, res.Message),
		"While calling the API, an error was returned in the response. "+
			"Please contact support if you are unsure how to resolve the error.\n\n"+
			"Error: "+fmt.Sprintf("%+v", res.Errors),
	)
}
