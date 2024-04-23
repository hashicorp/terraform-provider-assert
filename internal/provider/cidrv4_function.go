// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = CIDRv4Function{}
)

func NewCIDRv4Function() function.Function {
	return CIDRv4Function{}
}

type CIDRv4Function struct{}

func (r CIDRv4Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cidrv4"
}

func (r CIDRv4Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid CIDR notation (IPv4)",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "prefix",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r CIDRv4Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var prefix string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &prefix))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isCIDRv4(prefix)))
}

func isCIDRv4(prefix string) bool {
	ip, _, err := net.ParseCIDR(prefix)
	if err != nil {
		return false
	}
	return ip.To4() != nil
}
