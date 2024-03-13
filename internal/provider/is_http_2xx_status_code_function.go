// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IsHTTP2XXStatusCodeFunction{}
)

func NewIsHTTP2XXStatusCodeFunction() function.Function {
	return IsHTTP2XXStatusCodeFunction{}
}

type IsHTTP2XXStatusCodeFunction struct{}

func (r IsHTTP2XXStatusCodeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "is_http_2xx_status_code"
}

func (r IsHTTP2XXStatusCodeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Checks whether the HTTP status code is a valid 2xx status code",
		MarkdownDescription: "Checks whether the HTTP status code is a valid 2xx status code",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The HTTP status code",
				Name:               "status_code",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsHTTP2XXStatusCodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode int

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))

	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, is2xxStatusCode(statusCode)))
}

// isValid2xxStatusCode checks if an HTTP status code is within the 2xx range
func is2xxStatusCode(statusCode int) bool {
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
