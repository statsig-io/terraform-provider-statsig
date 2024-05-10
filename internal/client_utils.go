package statsig

import "github.com/hashicorp/terraform-plugin-framework/diag"

func runWithDiagnostics(handleRequest func(diags diag.Diagnostics) (*APIResponse, error)) diag.Diagnostics {
	var diags diag.Diagnostics
	res, err := handleRequest(diags)
	if err != nil {
		diags.Append(InternalAPIErrorDiagnostic(err))
		return diags
	}

	if res.Errors != nil {
		diags.Append(APIErrorDiagnostic(res))
		return diags
	}

	return diags
}
