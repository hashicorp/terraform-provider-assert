// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = RegexMatchesFunction{}
)

func NewRegexMatchesFunction() function.Function {
	return RegexMatchesFunction{}
}

type RegexMatchesFunction struct{}

func (r RegexMatchesFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "regex"
}

func (r RegexMatchesFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check if a string matches a regular expression",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The regular expression pattern to match against",
				Name:               "pattern",
			},
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to match against the regular expression",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r RegexMatchesFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var pattern, value *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &pattern, &value))
	if resp.Error != nil {
		return
	}

	if pattern == nil || value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	patternCompiled, err := regexp.Compile(*pattern)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, patternCompiled.MatchString(*value)))
}
