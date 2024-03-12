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

type WithinRangeFunctioneModel struct {
	Input int64   `cty:"input"`
	Range []int64 `cty:"range"`
}

func (r WithinRangeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "within_range"
}

func (r WithinRangeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Within range function",
		MarkdownDescription: "Checks whether the input is within the range",
		Parameters: []function.Parameter{
			function.SetParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The range to check if the input is within",
				Name:               "range",
				ElementType:        types.Int64Type,
			},
			function.Int64Parameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The input to check if it is within the range",
				Name:               "input",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r WithinRangeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data WithinRangeFunctioneModel

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	for _, i := range data.Range {
		if data.Input == i {
			resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, true))
			return
		}
	}

	resp.Error = function.ConcatFuncErrors()
}
