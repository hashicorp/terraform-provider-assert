// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = EqualFunction{}
)

func NewEqualFunction() function.Function {
	return EqualFunction{}
}

type EqualFunction struct{}

func (r EqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "equal"
}

func (r EqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is equal to another number",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The number to compare",
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

func (r EqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var number *big.Float
	var compareAgainst *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &number, &compareAgainst))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, number.Cmp(compareAgainst) == 0))
}
