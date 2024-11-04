// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = ValidJSONFunction{}
)

func NewValidJSONFunction() function.Function {
	return ValidJSONFunction{}
}

type ValidJSONFunction struct{}

func (r ValidJSONFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "valid_json"
}

func (r ValidJSONFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a value is valid JSON",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The JSON value to check",
				Name:               "json",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ValidJSONFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	if value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	isValidJSON, err := isValidJSON(value)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isValidJSON))
}

func isValidJSON(v *string) (bool, error) {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(*v), &js)
	if err != nil {
		return false, err
	}
	return true, nil
}
