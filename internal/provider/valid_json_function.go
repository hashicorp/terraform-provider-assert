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
		Summary: "Checks whether a string is valid JSON",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The JSON string to check",
				Name:               "json",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ValidJSONFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var JSON *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &JSON))
	if resp.Error != nil {
		return
	}

	isValidJSON, err := isValidJSON(JSON)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isValidJSON))
}

func isValidJSON(JSON *string) (bool, error) {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(*JSON), &js)
	if err != nil {
		return false, nil
	}
	return true, nil
}
