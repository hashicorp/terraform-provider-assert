// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ function.Function = ValueFunction{}
)

func NewValueFunction() function.Function {
	return ValueFunction{}
}

type ValueFunction struct{}

func (r ValueFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "value"
}

func (r ValueFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a value exists in a map",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to check",
				Name:               "value",
			},
			function.MapParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The map to check",
				Name:               "map",
				ElementType:        basetypes.StringType{},
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ValueFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var v *string
	var m *map[string]string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &v, &m))
	if resp.Error != nil {
		return
	}

	if v == nil || m == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, hasValue(v, m)))
}

func hasValue(v *string, m *map[string]string) bool {
	if m == nil {
		return false
	}

	for _, value := range *m {
		if value == *v {
			return true
		}
	}

	return false
}
