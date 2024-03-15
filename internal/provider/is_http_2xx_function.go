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
	_ function.Function = IsHTTP2XXFunction{}
)

func NewIsHTTP2XXFunction() function.Function {
	return IsHTTP2XXFunction{}
}

type IsHTTP2XXFunction struct{}

func (r IsHTTP2XXFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "is_http_2xx"
}

func (r IsHTTP2XXFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether the HTTP status code is a valid 2xx status code",
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

func (r IsHTTP2XXFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode types.Int64
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, is2xxStatusCode(statusCode.ValueInt64())))
}

// isValid2xxStatusCode checks if an HTTP status code is within the 2xx range
func is2xxStatusCode(statusCode int64) bool {
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
