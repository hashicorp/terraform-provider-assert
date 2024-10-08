---
page_title: "cidrv6 function - terraform-provider-assert"
subcategory: "Network Functions"
description: |-
  Checks whether a string is a valid CIDR notation (IPv6)
---

# function: cidrv6



The network function `cidrv6` returns true if the provided CIDR range is a valid IPv6 CIDR notation. Otherwise, it returns `false`.

It parses the `prefix` as a CIDR notation IP address and prefix length, such as “2001:db8::/32,” as defined in RFC 4291.

To validate a CIDR range regardless of the IP version, use the `cidr` function.

## Terraform Test Example

```terraform
run "check_valid_ipv6_aws_subnet" {
  command = plan
  assert {
    condition     = provider::assert::cidrv6(aws_subnet.example.ipv6_cidr_block)
    error_message = "Subnet is not in valid IPv6 CIDR notation"
  }
}
```

## Variable Validation Example

```terraform
variable "ipv6_subnet" {
  type = string
  validation {
    condition     = provider::assert::cidrv6(var.ipv6_subnet)
    error_message = "Invalid IPv6 subnet"
  }
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
cidrv6(prefix string) bool
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `prefix` (String) The string to check


## Return Type

The return type of `cidrv6` is a boolean.
