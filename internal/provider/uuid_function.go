// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = UUIDFunction{}
)

func NewUUIDFunction() function.Function {
	return UUIDFunction{}
}

type UUIDFunction struct{}

func (r UUIDFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "uuid"
}

func (r UUIDFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid UUID",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "uuid",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r UUIDFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var uuid string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &uuid))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isUUID(uuid)))
}

func isUUID(uuidStr string) bool {
	_, err := uuid.ParseUUID(uuidStr)
	return err == nil
}
