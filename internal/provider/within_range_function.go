// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = WithinRangeFunction{}
)

func NewWithinRangeFunction() function.Function {
	return WithinRangeFunction{}
}

type WithinRangeFunction struct{}

func (r WithinRangeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "within_range"
}

func (r WithinRangeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is within a given range",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The beginning of the range",
				Name:               "begin",
			},
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The end of the range",
				Name:               "end",
			},
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The number to check",
				Name:               "number",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r WithinRangeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var begin, end, number types.Number
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &begin, &end, &number))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isInRange(number, begin, end)))
}

func isInRange(number, start, end types.Number) bool {
	return number.ValueBigFloat().Cmp(start.ValueBigFloat()) != -1 && number.ValueBigFloat().Cmp(end.ValueBigFloat()) != 1
}
