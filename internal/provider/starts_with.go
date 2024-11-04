// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = StartsWithFunction{}
)

func NewStartsWithFunction() function.Function {
	return StartsWithFunction{}
}

type StartsWithFunction struct{}

func (r StartsWithFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "starts_with"
}

func (r StartsWithFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string starts with another string",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The prefix to check for",
				Name:               "prefix",
			},
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r StartsWithFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var prefix, value *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &prefix, &value))
	if resp.Error != nil {
		return
	}

	if prefix == nil || value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, strings.HasPrefix(*value, *prefix)))
}
