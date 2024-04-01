---
page_title: "Provider: Assert"
description: |-
  The Assert provider provides functions to verify values in your Terraform configurations to make sure they meet specific criteria.
---

# Assert Provider

The Assert provider is a utility provider that helps practitioners 
to simplify the way they write assertions in their Terraform configurations.

This provider is typically used in Terraform test assertions or variable validation.

This provider does not manage any infrastructure, but instead provides a set of provider-defined functions
that can be used to assert that values in Terraform configurations.

Use the navigation to the left to read about the available resources.

## Example Usage

As of Terraform 1.8 and later, providers can implement functions that you can call from the Terraform configuration. 

Define the provider as a `required_provider` to use its functions

```terraform
terraform {
  required_providers {
    assert = {
      source = "bschaatsbergen/assert"
    }
  }
}
```

## Limitations

