// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ function.Function = KeyFunction{}
)

func NewKeyFunction() function.Function {
	return KeyFunction{}
}

type KeyFunction struct{}

func (r KeyFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "key"
}

func (r KeyFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a key exists in a map",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The key to check",
				Name:               "key",
			},
			function.MapParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The map to check",
				Name:               "map",
				ElementType:        basetypes.StringType{},
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r KeyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var key *string
	var mapValue *map[string]string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &key, &mapValue))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, hasKey(key, mapValue)))
}

func hasKey(key *string, mapValue *map[string]string) bool {
	if mapValue == nil {
		return false
	}

	_, ok := (*mapValue)[*key]
	return ok
}
