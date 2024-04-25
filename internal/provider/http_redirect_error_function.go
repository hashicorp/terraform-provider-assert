// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-3.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = IsHTTPRedirectFunction{}
)

func NewIsHTTPRedirectFunction() function.Function {
	return IsHTTPRedirectFunction{}
}

type IsHTTPRedirectFunction struct{}

func (r IsHTTPRedirectFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "http_redirect"
}

func (r IsHTTPRedirectFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether an HTTP status code is a redirect status code",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The HTTP status code to check",
				Name:               "status_code",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsHTTPRedirectFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var statusCode types.Int64

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &statusCode))
	if resp.Error != nil {
		return
	}

	result := isRedirectStatusCode(statusCode.ValueInt64())

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// isValid3xxStatusCode checks if an HTTP status code is within the 3xx range.
func isRedirectStatusCode(statusCode int64) bool {
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
