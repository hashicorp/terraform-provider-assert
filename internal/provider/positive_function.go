// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = PositiveFunction{}
)

func NewPositiveFunction() function.Function {
	return PositiveFunction{}
}

type PositiveFunction struct{}

func (r PositiveFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "positive"
}

func (r PositiveFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a number is positive",
		Parameters: []function.Parameter{
			function.NumberParameter{
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The number to check",
				Name:               "number",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r PositiveFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var number *big.Float

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &number))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, number.Cmp(big.NewFloat(0)) == 1))
}
