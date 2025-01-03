// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/config"
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
  something = null
}
output "test" {
  value = provider::assert::null(local.something)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNullFunction_crossObjectValidation(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
		},
		ExternalProviders: map[string]resource.ExternalProvider{
			"wireguard": {
				Source:            "OJFord/wireguard",
				VersionConstraint: "0.3.1",
			},
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "wireguard_asymmetric_key" "main" {}

data "wireguard_config_document" "main" {
  private_key = wireguard_asymmetric_key.main.private_key
}

output "test" {
  // .addresses is always null in this configuration
  value = provider::assert::null(data.wireguard_config_document.main.addresses)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNullFunction_compoundValidation(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.2.0"))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
variable "example_value_a" {
  default     = null
  description = "Example option A for testing null validation."
  type        = string
}

variable "example_value_b" {
  default     = null
  description = "Example option B for testing null validation."
  type        = string

  validation {
    condition = anytrue([
      !provider::assert::null(var.example_value_a),
      !provider::assert::null(var.example_value_b)
    ])
    error_message = "Exactly one of example_value_a or example_value_b must be provided."
  }
}
				`,
				ConfigVariables: config.Variables{
					"example_value_b": config.StringVariable("example-value"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestNullFunction_falseCases(t *testing.T) {
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
  obj = {
    foo = "Foo"
    bar = "Bar"
  }
}
output "test" {
  value = provider::assert::null(local.obj)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
locals {
  name = "John Doe"
}
output "test" {
  value = provider::assert::null(local.name)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
locals {
  number = 14
}
output "test" {
  value = provider::assert::null(local.number)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
locals {
  number = max(1, 2)
}
output "test" {
  value = provider::assert::null(local.number)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
locals {
  fruits = ["apple", "banana", "cherry"]
}
output "test" {
  value = provider::assert::null(local.fruits)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
locals {
  kvmap = {
    "first_name" = "John"
    "last_name"  = "Doe"
  }
}
output "test" {
  value = provider::assert::null(local.kvmap)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
locals {
  set = toset(["apple", "banana", "cherry"])
}
output "test" {
  value = provider::assert::null(local.set)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
		},
	})
}
