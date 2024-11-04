// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"gopkg.in/yaml.v3"
)

var (
	_ function.Function = ValidYAMLFunction{}
)

func NewValidYAMLFunction() function.Function {
	return ValidYAMLFunction{}
}

type ValidYAMLFunction struct{}

func (r ValidYAMLFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "valid_yaml"
}

func (r ValidYAMLFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a value is valid YAML",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The YAML value to check",
				Name:               "yaml",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ValidYAMLFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	if value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	isValidYAML, err := isValidYAML(value)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isValidYAML))
}

func isValidYAML(v *string) (bool, error) {
	var js map[string]interface{}
	err := yaml.Unmarshal([]byte(*v), &js)
	if err != nil {
		return false, err
	}
	return true, nil
}
