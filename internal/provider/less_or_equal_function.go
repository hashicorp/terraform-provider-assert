// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = LessOrEqualFunction{}
)

func NewLessOrEqualFunction() function.Function {
	return LessOrEqualFunction{}
}

type LessOrEqualFunction struct{}

func (r LessOrEqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "less_or_equal"
}

func (r LessOrEqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is less than or equal to a given number",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The number to check",
				Name:               "number",
			},
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The number to compare against",
				Name:               "compare_against",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r LessOrEqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var number *big.Float
	var compareAgainst *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &number, &compareAgainst))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isLessOrEqual(number, compareAgainst)))
}

func isLessOrEqual(number, compareAgainst *big.Float) bool {
	return number.Cmp(compareAgainst) <= 0
}
