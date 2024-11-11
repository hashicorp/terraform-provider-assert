// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestTimeCheckFunction(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
output "test1" {
 value = provider::assert::timecheck(["ALL"],["ALL"],"America/New_York","00:00","23:59")
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test1", "true"),
				),
			},
			{
				Config: `
output "test2" {
  value = provider::assert::timecheck(["ALL"],["ALL"],"India/New_York","00:00","23:59")
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test2", "false"),
				),
			},
		},
	})
}
