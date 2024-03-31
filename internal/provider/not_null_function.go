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
	resp.Name = "not_null"
}

func (r NotNullFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a given object is not null",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				AllowNullValue:     true,
				AllowUnknownValues: true,
				Description:        "The object to check",
				Name:               "object",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r NotNullFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data types.Dynamic

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, !data.IsNull()))
}
