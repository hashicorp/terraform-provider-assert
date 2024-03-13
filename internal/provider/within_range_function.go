// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
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
		Summary:             "Within range function",
		MarkdownDescription: "Checks whether the input is within the range",
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
				Description:        "The number to check if it is within the range",
				Name:               "number",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r WithinRangeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var begin int
	var end int
	var number int

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &begin, &end, &number))

	if resp.Error != nil {
		return
	}

	isInRange := func(number, start, end int) bool {
		return number >= start && number <= end
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isInRange(number, begin, end)))
}
