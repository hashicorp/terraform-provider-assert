// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The HTTP status code to check",
				Name:               "status_code",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsHTTPSuccessFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode types.Int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	result := isSuccessStatusCode(statusCode.ValueInt64())

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// isValid2xxStatusCode checks if an HTTP status code is within the 2xx range.
func isSuccessStatusCode(statusCode int64) bool {
	switch statusCode {
	case
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed:
		return true
	default:
		return false
	}
}
