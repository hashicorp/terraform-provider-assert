// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = NotNullFunction{}
)

func NewNotNullFunction() function.Function {
	return NotNullFunction{}
}

type NotNullFunction struct{}

func (r NotNullFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "notnull"
}

func (r NotNullFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Not null function",
		MarkdownDescription: "Checks that the input is not null",
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

func (r NotNullFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data types.Object

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, !data.IsNull()))
}
