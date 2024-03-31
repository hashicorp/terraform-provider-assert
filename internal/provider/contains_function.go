// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	tpftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = ContainsFunction{}
)

func NewContainsFunction() function.Function {
	return ContainsFunction{}
}

type ContainsFunction struct{}

func (r ContainsFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "contains"
}

func (r ContainsFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether an element is in a list",
		Parameters: []function.Parameter{
			function.ListParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The list to check",
				Name:               "list",
				ElementType:        tpftypes.StringType,
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The element to check",
				Name:               "element",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ContainsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var list []string
	var value string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &list, &value))
	if resp.Error != nil {
		return
	}

	for _, item := range list {
		if item == value {
			resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, true))
			return
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
}
