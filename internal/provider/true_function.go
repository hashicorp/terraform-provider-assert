// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = TrueFunction{}
)

func NewTrueFunction() function.Function {
	return TrueFunction{}
}

type TrueFunction struct{}

func (r TrueFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "true"
}

func (r TrueFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a boolean value is true",
		Parameters: []function.Parameter{
			function.BoolParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The boolean value to check",
				Name:               "bool",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r TrueFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value bool

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))
}
