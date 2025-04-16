// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = SemVerCheckFunction{}
)

func NewSemVerCheckFunction() function.Function {
	return SemVerCheckFunction{}
}

type SemVerCheckFunction struct{}

func (r SemVerCheckFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "semver_check"
}

func (r SemVerCheckFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check if a semver matches a constraint",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The constraint to check against",
				Name:               "constraint",
			},
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The version to check",
				Name:               "semver",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r SemVerCheckFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var c, v string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &c, &v))
	if resp.Error != nil {
		return
	}

	semver, err := version.NewVersion(v)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	constraints, err := version.NewConstraint(c)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, constraints.Check(semver)))
}
