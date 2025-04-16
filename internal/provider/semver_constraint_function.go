// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = SemVerConstraintFunction{}
)

func NewSemVerConstraintFunction() function.Function {
	return SemVerConstraintFunction{}
}

type SemVerConstraintFunction struct{}

func (r SemVerConstraintFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "semver_constraint"
}

func (r SemVerConstraintFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check if a semver constraint is valid",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The constraint to validate",
				Name:               "constraint",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r SemVerConstraintFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var constraint string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &constraint))
	if resp.Error != nil {
		return
	}

	_, err := version.NewConstraint(constraint)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, err == nil))
}
