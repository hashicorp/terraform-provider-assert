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

func NewEmptyFunction() function.Function {
	return EmptyFunction{}
}

type EmptyFunction struct{}

func (r EmptyFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "empty"
}

func (r EmptyFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks wether a given string is empty",
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

func (r EmptyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var argument string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &argument))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isEmpty(argument)))
}

func isEmpty(argument string) bool {
	return len(argument) == 0
}
