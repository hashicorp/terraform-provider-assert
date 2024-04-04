# Adding a New Function

!!! tip
    Provider-defined function support is in technical preview and offered without compatibility promises until Terraform 1.8 is generally available.

Provider-defined functions were introduced with Terraform 1.8, enabling provider developers to expose functions specific to a given cloud provider or use case.
Functions in the Assert provider provide a utility that is valuable when paired with Terraform tests or variable validation.

See the Terraform Plugin Framework [Function documentation](https://developer.hashicorp.com/terraform/plugin/framework/functions) for additional details.

## Prerequisites

The only prerequisite for creating a function is ensuring the desired functionality is appropriate for a provider-defined function.
Functions must be reproducible across executions ("pure" functions), where the same input always results in the same output.
This requirement precludes the use of network calls, time-based operations, or other non-deterministic operations.

## Steps to add a function

### Fork the provider and create a feature branch

For a new function use a branch named `f-{function name}`, for example, `f-equals`.

### Create and name the function

The function name should be descriptive of its functionality and succinct.
Existing examples include `null`, `equals`, and `contains`.

New functions can be created by copying the format of an existing function inside `internal/provider`.

### Fill out the function parameter(s) and return value

The function struct's `Definition` method will document the expected parameters and return value.
Parameter names and return values should be specified in `camel_case`.

```go
func (r NullFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a given argument is null",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				AllowNullValue:     true,
				AllowUnknownValues: true,
				Description:        "The argument to check",
				Name:               "argument",
			},
		},
		Return: function.BoolReturn{},
	}
}
```

The example above defines a function which accepts a dynamic parameter named `argument` and returns a boolean value.

### Implement the function logic

The function struct's `Run` method will contain the function logic.
This includes processing the arguments, setting the return value, and any data processing that needs to happen in between.

```go
func (r NullFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var argument types.Dynamic

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &argument))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, !argument.IsNull()))
}
```

### Register function to the provider

Once the function is implemented, it must be registered to the provider to be used.
This is done by adding the function to the `Functions` method in `internal/provider/provider.go`.

```go
func (p *AssertProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewNullFunction,
	}
}
```

### Write passing tests

All functions should have corresponding acceptance tests.
For functions with variadic arguments, or which can potentially return an error, tests should be written to exercise those conditions.

An example outline is included below:

```go
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestNullFunction(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
locals {
  example = null
}
output "is_null" {
  value = provider::assert::null(locals.example)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}
```

With Terraform 1.8+ installed, individual tests can be run like:

```console
go test -run='^TestExample' -v ./internal/provider/
```
