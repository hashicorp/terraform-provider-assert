// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IsIPFunction{}
)

func NewIsIPFunction() function.Function {
	return IsIPFunction{}
}

type IsIPFunction struct{}

func (r IsIPFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ip"
}

func (r IsIPFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid IP address (IPv4 or IPv6)",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "ip_address",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IsIPFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var ip string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &ip))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isIP(ip)))
}

func isIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
