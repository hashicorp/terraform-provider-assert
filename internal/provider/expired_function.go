// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = ExpiredFunction{}
)

func NewExpiredFunction() function.Function {
	return ExpiredFunction{}
}

type ExpiredFunction struct{}

func (r ExpiredFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "expired"
}

func (r ExpiredFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a timestamp in RFC3339 format is expired",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "timestamp",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r ExpiredFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var timestamp string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &timestamp))
	if resp.Error != nil {
		return
	}

	expired, err := isExpired(timestamp)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, expired))
}

func isExpired(timestamp string) (bool, error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return false, err
	}

	return t.Before(time.Now()), nil
}
