// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-4.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IsHTTPServerErrorFunction{}
)

func NewIsHTTPServerErrorFunction() function.Function {
	return IsHTTPServerErrorFunction{}
}

type IsHTTPServerErrorFunction struct{}

func (r IsHTTPServerErrorFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "http_server_error"
}

func (r IsHTTPServerErrorFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether an HTTP status code is a server error status code",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The HTTP status code to check",
				Name:               "status_code",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsHTTPServerErrorFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode *int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	result := isHTTPServerError(statusCode)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// isHTTPServerError checks if an HTTP status code is within the 5xx range.
func isHTTPServerError(statusCode *int64) bool {
	// Check if statusCode is nil
	if statusCode == nil {
		return false
	}

	// Check if the status code is in the 5xx range
	return *statusCode >= 500 && *statusCode < 600
}
