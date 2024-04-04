---
page_title: "negative function - terraform-provider-assert"
subcategory: "Numeric Functions"
description: |-
  Checks whether a number is negative
---

# function: negative



## Variable Validation Example

```terraform
variable "example" {
  type = number
  validation {
    condition     = provider::assert::negative(var.example)
    error_message = "Expected ${var.example} to be negative"
  }
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
negative(number number) bool
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `number` (Number) The number to check


## Return Type

The return type of `negative` is a boolean.