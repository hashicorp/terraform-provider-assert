// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = LessFunction{}
)

func NewLessFunction() function.Function {
	return LessFunction{}
}

type LessFunction struct{}

func (r LessFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "less"
}

func (r LessFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is less than a given number",
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

func (r LessFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
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

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isLessThan(value, comparisonTarget)))
}

func isLessThan(value, comparisonTarget *big.Float) bool {
	return value.Cmp(comparisonTarget) == -1
}
