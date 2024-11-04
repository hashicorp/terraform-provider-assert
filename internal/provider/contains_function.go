// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The list to check",
				Name:               "list",
				ElementType:        types.StringType,
			},
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The element to check",
				Name:               "element",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ContainsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var list *[]string
	var element *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &list, &element))
	if resp.Error != nil {
		return
	}

	// Return false if list is empty or element is empty
	if list == nil || element == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	for _, item := range *list {
		if item == *element {
			resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, true))
			return
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
}
