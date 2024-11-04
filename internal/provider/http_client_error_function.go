// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-4.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IsHTTPClientErrorFunction{}
)

func NewIsHTTPClientErrorFunction() function.Function {
	return IsHTTPClientErrorFunction{}
}

type IsHTTPClientErrorFunction struct{}

func (r IsHTTPClientErrorFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "http_client_error"
}

func (r IsHTTPClientErrorFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether an HTTP status code is a client error status code",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The HTTP status code to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsHTTPClientErrorFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value *int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	if value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	result := isHTTPClientError(value)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// isHTTPClientError checks if an HTTP status code is within the 4xx range.
func isHTTPClientError(v *int64) bool {
	// Check if statusCode is nil
	if v == nil {
		return false
	}

	// Check if the status code is in the 4xx range
	return *v >= 400 && *v < 500
}
