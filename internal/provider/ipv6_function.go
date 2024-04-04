// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IPv6Function{}
)

func NewIPv6Function() function.Function {
	return IPv6Function{}
}

type IPv6Function struct{}

func (r IPv6Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ipv6"
}

func (r IPv6Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid IPv6 address",
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

func (r IPv6Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var ip string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &ip))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isIPv6(ip)))
}

func isIPv6(ip string) bool {
	// Because IPv4 addresses can also have a 16-byte representation,
	// we need to check that the address is not an IPv4 address first.
	if net.ParseIP(ip).To4() != nil {
		return false
	}
	return net.ParseIP(ip).To16() != nil
}
