// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = CIDRv6Function{}
)

func NewCIDRv6Function() function.Function {
	return CIDRv6Function{}
}

type CIDRv6Function struct{}

func (r CIDRv6Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cidrv6"
}

func (r CIDRv6Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid CIDR notation (IPv6)",
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

func (r CIDRv6Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var prefix string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &prefix))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isCIDRv6(prefix)))
}

func isCIDRv6(prefix string) bool {
	ip, _, err := net.ParseCIDR(prefix)
	if err != nil {
		return false
	}

	// Because IPv4 addresses can also have a 16-byte representation,
	// we need to check that the address is not an IPv4 address first.
	if ip.To4() != nil {
		return false
	}
	return ip.To16() != nil
}
