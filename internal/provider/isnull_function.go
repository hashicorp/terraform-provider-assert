// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = IsNullFunction{}
)

func NewIsNullFunction() function.Function {
	return IsNullFunction{}
}

type IsNullFunction struct{}

func (r IsNullFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "isnull"
}

func (r IsNullFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Is null function",
		MarkdownDescription: "Checks that the input is null",
		Parameters: []function.Parameter{
			function.ObjectParameter{
				AllowNullValue:     true,
				AllowUnknownValues: true,
				Description:        "The input to check",
				Name:               "input",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsNullFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data types.Object

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, data.IsNull()))
}
