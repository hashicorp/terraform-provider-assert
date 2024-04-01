// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = NotEqualFunction{}
)

func NewNotEqualFunction() function.Function {
	return NotEqualFunction{}
}

type NotEqualFunction struct{}

func (r NotEqualFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "not_equal"
}

func (r NotEqualFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether an element is not equal to another element",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: true,
				Description:        "The element to compare",
				Name:               "number",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The element to compare against",
				Name:               "compare_against",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r NotEqualFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var element string
	var compareAgainst string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &element, &compareAgainst))
	if resp.Error != nil {
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, element != compareAgainst))
}
