// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = FalseFunction{}
)

func NewFalseFunction() function.Function {
	return FalseFunction{}
}

type FalseFunction struct{}

func (r FalseFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "false"
}

func (r FalseFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a boolean value is false",
		Parameters: []function.Parameter{
			function.BoolParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The boolean value to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r FalseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value *bool

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	if value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, !*value))
}
