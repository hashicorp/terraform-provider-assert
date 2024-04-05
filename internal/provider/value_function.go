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
			function.MapParameter{
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The map to check",
				Name:               "map",
				ElementType:        basetypes.StringType{},
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The value to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ValueFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var mapValue *map[string]string
	var value *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &mapValue, &value))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, hasValue(mapValue, value)))
}

func hasValue(mapValue *map[string]string, value *string) bool {
	if mapValue == nil {
		return false
	}

	for _, v := range *mapValue {
		if v == *value {
			return true
		}
	}

	return false
}
