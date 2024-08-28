// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestRegexMatchesFunction(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
output "test" {
  value = provider::assert::regex("needle", "hay needle stack")
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestRegexMatchesFunction_trueCases(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		pattern string
		s       string
	}{
		{
			name:    "simple regex matches",
			pattern: "needle",
			s:       "hay needle stack",
		},
		{
			name:    "regex matches exact string",
			pattern: "needle",
			s:       "needle",
		},
		{
			name:    "regex matches with special characters",
			pattern: "needle\\s+",
			s:       "hay needle    stack",
		},
		{
			name:    "regex matches with positional anchors in front",
			pattern: "^needle",
			s:       "needle hay stack",
		},
		{
			name:    "regex matches with positional anchors behind",
			pattern: "needle$",
			s:       "hay stack needle",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			resource.UnitTest(t, resource.TestCase{
				TerraformVersionChecks: []tfversion.TerraformVersionCheck{
					tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
				},
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
output "test" {
  value = provider::assert::regex("%s", "%s")
}
				`, test.pattern, test.s),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckOutput("test", "true"),
						),
					},
				},
			})
		})
	}
}

func TestRegexMatchesFunction_falseCases(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		pattern string
		s       string
	}{
		{
			name:    "simple regex does not match",
			pattern: "needle",
			s:       "hay hay stack",
		},
		{
			name:    "regex does not match exact string",
			pattern: "needle",
			s:       "needles",
		},
		{
			name:    "regex does not match with special characters",
			pattern: "needle\\s+",
			s:       "hay needlestack",
		},
		{
			name:    "regex does not match with positional anchors in front",
			pattern: "^needle",
			s:       "hay needle stack",
		},
		{
			name:    "regex does not match with positional anchors behind",
			pattern: "needle$",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			resource.UnitTest(t, resource.TestCase{
				TerraformVersionChecks: []tfversion.TerraformVersionCheck{
					tfversion.SkipBelow(version.Must(version.NewVersion(MinimalRequiredTerraformVersion))),
				},
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
output "test" {
  value = provider::assert::regex("%s", "%s")
}
				`, test.pattern, test.s),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckOutput("test", "false"),
						),
					},
				},
			})
		})
	}
}
