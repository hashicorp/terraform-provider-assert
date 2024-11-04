// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = BetweenFunction{}
)

func NewBetweenFunction() function.Function {
	return BetweenFunction{}
}

type BetweenFunction struct{}

func (r BetweenFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "between"
}

func (r BetweenFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is within a given range",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The beginning of the range",
				Name:               "begin",
			},
			function.NumberParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The end of the range",
				Name:               "end",
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

func (r BetweenFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var begin, end, value *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &begin, &end, &value))
	if resp.Error != nil {
		return
	}

	// Check if any of the values are nil, and if so, set the result to false and return early
	if begin == nil || end == nil || value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isInRange(value, begin, end)))
}

func isInRange(v, start, end *big.Float) bool {
	return v.Cmp(start) != -1 && v.Cmp(end) != 1
}
