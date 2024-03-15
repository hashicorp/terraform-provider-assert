// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-5.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = IsHTTP5XXFunction{}
)

func NewIsHTTP5XXFunction() function.Function {
	return IsHTTP5XXFunction{}
}

type IsHTTP5XXFunction struct{}

func (r IsHTTP5XXFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "is_http_5xx"
}

func (r IsHTTP5XXFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether the HTTP status code is a valid 5xx status code",
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

func (r IsHTTP5XXFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode types.Int64
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, is5xxStatusCode(statusCode.ValueInt64())))
}

// isValid5xxStatusCode checks if an HTTP status code is within the 5xx range
func is5xxStatusCode(statusCode int64) bool {
	switch statusCode {
	case
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected,
		http.StatusNotExtended,
		http.StatusNetworkAuthenticationRequired:
		return true
	default:
		return false
	}
}
