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
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to compare against",
				Name:               "compare_against",
			},
			function.NumberParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to compare",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r EqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
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

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value.Cmp(compareAgainst) == 0))
}
