// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure AssertProvider satisfies various provider interfaces.
var _ provider.Provider = &AssertProvider{}
var _ provider.ProviderWithFunctions = &AssertProvider{}

// AssertProvider defines the provider implementation.
type AssertProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *AssertProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "assert"
	resp.Version = p.version
}

func (p *AssertProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *AssertProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *AssertProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *AssertProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *AssertProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewNotNullFunction,
		NewIsNullFunction,
		NewBetweenFunction,
		NewContainsFunction,
		NewIsHTTPSuccessFunction,
		NewIsHTTPRedirectFunction,
		NewIsHTTPClientErrorFunction,
		NewIsHTTPServerErrorFunction,
		NewGreaterFunction,
		NewGreaterOrEqualFunction,
		NewLessFunction,
		NewLessOrEqualFunction,
		NewEqualFunction,
		NewNotEqualFunction,
		NewTrueFunction,
		NewFalseFunction,
		NewValidJSONFunction,
		NewValidYAMLFunction,
		NewIPv4Function,
		NewIPv6Function,
		NewIPFunction,
		NewCIDRFunction,
		NewCIDRv4Function,
		NewCIDRv6Function,
		NewStartsWithFunction,
		NewEndsWithFunction,
		NewUppercasedFunction,
		NewLowercasedFunction,
		NewNegativeFunction,
		NewPositiveFunction,
		NewKeyFunction,
		NewValueFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AssertProvider{
			version: version,
		}
	}
}
