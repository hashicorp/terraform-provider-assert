// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = GreaterThanOrEqualFunction{}
)

func NewGreaterThanOrEqualFunction() function.Function {
	return GreaterThanOrEqualFunction{}
}

type GreaterThanOrEqualFunction struct{}

func (r GreaterThanOrEqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "greater_than_or_equal"
}

func (r GreaterThanOrEqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is greater than or equal to a given number",
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

func (r GreaterThanOrEqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var number types.Number
	var compareAgainst types.Number
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &number, &compareAgainst))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isGreaterThanOrEqual(number, compareAgainst)))
}

func isGreaterThanOrEqual(number, compareAgainst types.Number) bool {
	return number.ValueBigFloat().Cmp(compareAgainst.ValueBigFloat()) >= 0
}
