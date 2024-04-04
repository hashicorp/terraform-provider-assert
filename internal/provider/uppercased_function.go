// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = UppercasedFunction{}
)

func NewUppercasedFunction() function.Function {
	return UppercasedFunction{}
}

type UppercasedFunction struct{}

func (r UppercasedFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "uppercased"
}

func (r UppercasedFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is uppercased",
		Parameters: []function.Parameter{
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

func (r UppercasedFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var s string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &s))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isUpper(s)))
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
