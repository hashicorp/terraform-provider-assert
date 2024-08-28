// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"regexp"
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
		Summary: "Checks whether a string matches a regular expression pattern",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The regular expression pattern to check for",
				Name:               "pattern",
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

func (r RegexMatchesFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var patternRaw, s string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &patternRaw, &s))
	if resp.Error != nil {
		return
	}

	patternCompiled, err := regexp.Compile(patternRaw)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, patternCompiled.MatchString(s)))
}
