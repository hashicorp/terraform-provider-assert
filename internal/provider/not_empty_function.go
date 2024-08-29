// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = EmptyFunction{}
)

func NewNotEmptyFunction() function.Function {
	return NotEmptyFunction{}
}

type NotEmptyFunction struct{}

func (r NotEmptyFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "not_empty"
}

func (r NotEmptyFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks wether a given string is not empty",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "s",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r NotEmptyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var argument string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &argument))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isNotEmpty(argument)))
}

func isNotEmpty(argument string) bool {
	return len(argument) > 0
}
