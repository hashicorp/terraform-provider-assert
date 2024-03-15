// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-4.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = IsHTTP4XXFunction{}
)

func NewIsHTTP4XXFunction() function.Function {
	return IsHTTP4XXFunction{}
}

type IsHTTP4XXFunction struct{}

func (r IsHTTP4XXFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "is_http_4xx"
}

func (r IsHTTP4XXFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether the HTTP status code is a valid 4xx status code",
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

func (r IsHTTP4XXFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode types.Int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, is4xxStatusCode(statusCode.ValueInt64())))
}

// isValid4xxStatusCode checks if an HTTP status code is within the 4xx range
func is4xxStatusCode(statusCode int64) bool {
	switch statusCode {
	case
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusPaymentRequired,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusProxyAuthRequired,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed,
		http.StatusTeapot,
		http.StatusMisdirectedRequest,
		http.StatusUnprocessableEntity,
		http.StatusLocked,
		http.StatusFailedDependency,
		http.StatusTooEarly,
		http.StatusUpgradeRequired,
		http.StatusPreconditionRequired,
		http.StatusTooManyRequests,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnavailableForLegalReasons:
		return true
	default:
		return false
	}
}
