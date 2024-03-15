// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = LessThanOrEqualFunction{}
)

func NewLessThanOrEqualFunction() function.Function {
	return LessThanOrEqualFunction{}
}

type LessThanOrEqualFunction struct{}

func (r LessThanOrEqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "less_than_or_equal"
}

func (r LessThanOrEqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is less than or equal to a given number",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     true,
				AllowUnknownValues: true,
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

func (r LessThanOrEqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var number types.Number
	var compareAgainst types.Number
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &number, &compareAgainst))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isLessThanOrEqual(number, compareAgainst)))
}

func isLessThanOrEqual(number, compareAgainst types.Number) bool {
	return number.ValueBigFloat().Cmp(compareAgainst.ValueBigFloat()) <= 0
}
