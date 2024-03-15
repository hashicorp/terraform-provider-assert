// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-3.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IsHTTP3XXFunction{}
)

func NewIsHTTP3XXFunction() function.Function {
	return IsHTTP3XXFunction{}
}

type IsHTTP3XXFunction struct{}

func (r IsHTTP3XXFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "is_http_3xx"
}

func (r IsHTTP3XXFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether the HTTP status code is a valid 3xx status code",
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

func (r IsHTTP3XXFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode int
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, is3xxStatusCode(statusCode)))
}

// isValid3xxStatusCode checks if an HTTP status code is within the 3xx range
func is3xxStatusCode(statusCode int) bool {
	switch statusCode {
	case
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect:
		return true
	default:
		return false
	}
}
