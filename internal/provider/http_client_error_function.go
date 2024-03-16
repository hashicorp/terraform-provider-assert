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
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The HTTP status code to check",
				Name:               "status_code",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsHTTPClientErrorFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode types.Int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isHTTPClientError(statusCode.ValueInt64())))
}

// isHTTPClientError checks if an HTTP status code is within the 4xx range.
func isHTTPClientError(statusCode int64) bool {
	switch statusCode {
	case
		// 4XX status codes
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
		http.StatusUnavailableForLegalReasons,
		// 5XX status codes
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
