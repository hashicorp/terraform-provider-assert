// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = NotEqualFunction{}
)

func NewNotEqualFunction() function.Function {
	return NotEqualFunction{}
}

type NotEqualFunction struct{}

func (r NotEqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "not_equal"
}

func (r NotEqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is not equal to another number",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to compare against",
				Name:               "compare_against",
			},
			function.NumberParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to compare",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r NotEqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var comparisonTarget *big.Float
	var value *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &comparisonTarget, &value))
	if resp.Error != nil {
		return
	}

	if comparisonTarget == nil || value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value.Cmp(comparisonTarget) != 0))
}
