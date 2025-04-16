// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = SemVerVersionFunction{}
)

func NewSemVerVersionFunction() function.Function {
	return SemVerVersionFunction{}
}

type SemVerVersionFunction struct{}

func (r SemVerVersionFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "semver_version"
}

func (r SemVerVersionFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check if a semver version is valid",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The version to validate",
				Name:               "version",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r SemVerVersionFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var semver string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &semver))
	if resp.Error != nil {
		return
	}

	_, err := version.NewVersion(semver)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, err == nil))
}
