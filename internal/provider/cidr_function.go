// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = CIDRFunction{}
)

func NewCIDRFunction() function.Function {
	return CIDRFunction{}
}

type CIDRFunction struct{}

func (r CIDRFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cidr"
}

func (r CIDRFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid CIDR notation (IPv4 or IPv6)",
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

func (r CIDRFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var prefix string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &prefix))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isCIDR(prefix)))
}

func isCIDR(prefix string) bool {
	_, _, err := net.ParseCIDR(prefix)
	return err == nil
}
