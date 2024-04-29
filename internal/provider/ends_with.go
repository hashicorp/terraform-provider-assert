// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = EndsWithFunction{}
)

func NewEndsWithFunction() function.Function {
	return EndsWithFunction{}
}

type EndsWithFunction struct{}

func (r EndsWithFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ends_with"
}

func (r EndsWithFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string ends with another string",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The suffix to check for",
				Name:               "suffix",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "string",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r EndsWithFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var suffix, s string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &suffix, &s))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, strings.HasSuffix(s, suffix)))
}
