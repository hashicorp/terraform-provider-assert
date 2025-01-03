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

func TestNotNullFunction(t *testing.T) {
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
  value = provider::assert::not_null(local.obj)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_object(t *testing.T) {
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
  value = provider::assert::not_null(local.obj)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_string(t *testing.T) {
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
  name = "John Doe"
}
output "test" {
  value = provider::assert::not_null(local.name)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_int(t *testing.T) {
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
  number = 14
}
output "test" {
  value = provider::assert::not_null(local.number)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_function(t *testing.T) {
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
  number = max(1, 2)
}
output "test" {
  value = provider::assert::not_null(local.number)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_list(t *testing.T) {
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
  fruits = ["apple", "banana", "cherry"]
}
output "test" {
  value = provider::assert::not_null(local.fruits)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_map(t *testing.T) {
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
  kvmap = {
    "first_name" = "John"
    "last_name"  = "Doe"
  }
}
output "test" {
  value = provider::assert::not_null(local.kvmap)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_set(t *testing.T) {
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
  set = toset(["apple", "banana", "cherry"])
}
output "test" {
  value = provider::assert::not_null(local.set)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestNotNullFunction_compoundValidation(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
variable "ipv4_ipam_pool_id" {
  default     = null
  description = "ID of the IPv4 IPAM pool to use for the VPC."
  type        = string
}

variable "cidr_block" {
  default     = null
  description = "CIDR block for the VPC."
  type        = string

  validation {
    condition = provider::assert::cidr(var.cidr_block)
    error_message = "CIDR block must be a valid CIDR range."
  }

  validation {
    condition = anytrue([
	  provider::assert::not_null(var.cidr_block), 
	  provider::assert::not_null(var.ipv4_ipam_pool_id)
	])
    error_message = "Exactly one of cidr_block or ipv4_ipam_pool_id must be provided."
  }
}
				`,
				ConfigVariables: config.Variables{
					"cidr_block": config.StringVariable("10.0.42.0/24"),
				},
				Check: resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestNotNullFunction_falseCases(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"wireguard": {
				Source:            "OJFord/wireguard",
				VersionConstraint: "0.3.1",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `
locals {
  object = null
}
output "test" {
  value = provider::assert::not_null(local.object)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
			{
				Config: `
resource "wireguard_asymmetric_key" "main" {}

data "wireguard_config_document" "main" {
  private_key = wireguard_asymmetric_key.main.private_key
}

output "test" {
  // .addresses is always null in this configuration
  value = provider::assert::not_null(data.wireguard_config_document.main.addresses)
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
		},
	})
}
