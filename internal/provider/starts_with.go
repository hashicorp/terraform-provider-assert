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
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "string",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The prefix to check for",
				Name:               "prefix",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r StartsWithFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var s, prefix string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &s, &prefix))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, strings.HasPrefix(s, prefix)))
}
