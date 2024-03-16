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
		Summary: "Checks whether a string is valid YAML",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: true,
				Description:        "The YAML string to check",
				Name:               "yaml",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ValidYAMLFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var YAML *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &YAML))
	if resp.Error != nil {
		return
	}

	isValidYAML, err := isValidYAML(YAML)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isValidYAML))
}

func isValidYAML(YAML *string) (bool, error) {
	var js map[string]interface{}
	err := yaml.Unmarshal([]byte(*YAML), &js)
	if err != nil {
		return false, err
	}
	return true, nil
}
