package statsig

import "github.com/hashicorp/terraform-plugin-framework/diag"

func runWithDiagnostics(handleRequest func(diags diag.Diagnostics) (*APIResponse, error)) diag.Diagnostics {
	var diags diag.Diagnostics
	res, err := handleRequest(diags)
	if err != nil {
		diags.Append(InternalAPIErrorDiagnostic(err))
		return diags
	}

	// The Console API reports 4xx as {message, status} and never populates
	// "errors", so the status has to be checked as well or the failure is
	// swallowed and the zero-value response is written back into state.
	if res.StatusCode >= 400 || res.Errors != nil {
		diags.Append(APIErrorDiagnostic(res))
		return diags
	}

	return diags
}
