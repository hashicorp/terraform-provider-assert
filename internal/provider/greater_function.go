// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = GreaterFunction{}
)

func NewGreaterFunction() function.Function {
	return GreaterFunction{}
}

type GreaterFunction struct{}

func (r GreaterFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "greater"
}

func (r GreaterFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is greater than a given number",
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
				Description:        "The value to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r GreaterFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var compareAgainst *big.Float
	var value *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &compareAgainst, &value))
	if resp.Error != nil {
		return
	}

	if compareAgainst == nil || value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isGreaterThan(value, compareAgainst)))
}

func isGreaterThan(v, compareAgainst *big.Float) bool {
	return v.Cmp(compareAgainst) == 1
}
