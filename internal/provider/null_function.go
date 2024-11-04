// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = IsNullFunction{}
)

func NewIsNullFunction() function.Function {
	return IsNullFunction{}
}

type IsNullFunction struct{}

func (r IsNullFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "null"
}

func (r IsNullFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a given value is null",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				AllowNullValue:     true,
				AllowUnknownValues: true,
				Description:        "The value to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsNullFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value types.Dynamic

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	if value.UnderlyingValue() == nil {
		if err := resp.Result.Set(ctx, true); err == nil {
			return
		}
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
}
