---
page_title: "semver_check function - terraform-provider-assert"
subcategory: "SemVer Functions"
description: |-
  Check if a semver matches a constraint
---

# function: semver_check





## Variable Validation Example

```terraform
variable "version" {
  type = string
  validation {
    condition     = provider::assert::semver_check("~> 1.0", var.version)
    error_message = "The provided version is not supported"
  }
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
semver_check(constraint string, semver string) bool
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `constraint` (String) The constraint to check against
1. `semver` (String) The version to check


## Return Type

The return type of `semver_check` is a boolean.