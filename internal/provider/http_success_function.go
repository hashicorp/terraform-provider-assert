// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IsHTTPSuccessFunction{}
)

func NewIsHTTPSuccessFunction() function.Function {
	return IsHTTPSuccessFunction{}
}

type IsHTTPSuccessFunction struct{}

func (r IsHTTPSuccessFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "http_success"
}

func (r IsHTTPSuccessFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether an HTTP status code is a success status code",
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

func (r IsHTTPSuccessFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value *int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	result := isSuccessStatusCode(value)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// isSuccessStatusCode checks if an HTTP status code is within the 2xx range.
func isSuccessStatusCode(v *int64) bool {
	// Check if statusCode is nil
	if v == nil {
		return false
	}

	// Check if the status code is in the 2xx range
	return *v >= 200 && *v < 300
}
