// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IPFunction{}
)

func NewIPFunction() function.Function {
	return IPFunction{}
}

type IPFunction struct{}

func (r IPFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ip"
}

func (r IPFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid IP address (IPv4 or IPv6)",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The value to check",
				Name:               "value",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IPFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var value *string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &value))
	if resp.Error != nil {
		return
	}

	if value == nil {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, false))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isIP(value)))
}

func isIP(v *string) bool {
	return net.ParseIP(*v) != nil
}
