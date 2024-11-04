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
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The key to check",
				Name:               "key",
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

func (r KeyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var k *string
	var m *map[string]string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &k, &m))
	if resp.Error != nil {
		return
	}

	if k == nil || m == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, hasKey(k, m)))
}

func hasKey(k *string, m *map[string]string) bool {
	if m == nil {
		return false
	}

	_, ok := (*m)[*k]
	return ok
}
