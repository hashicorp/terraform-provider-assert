# Terraform Provider: Assert

The [Assert Terraform provider]((https://registry.terraform.io/providers/hashicorp/assert/latest/docs)) is intended for use when writing [Terraform Tests](https://developer.hashicorp.com/terraform/language/tests), [Variable Validation](https://developer.hashicorp.com/terraform/language/values/variables#custom-validation-rules), [Preconditions and Postconditions](https://developer.hashicorp.com/terraform/language/expressions/custom-conditions#preconditions-and-postconditions), or [Continiuous Validation](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/health#continuous-validation). It serves as a way to verify that the values in your Terraform configuration meet specific criteria. The provider only contains functions to assert values, and does not manage any resources.

* [Terraform Registry](https://registry.terraform.io/providers/hashicorp/assert/latest/docs)
* [Contributor Guide](https://hashicorp.github.io/terraform-provider-assert/)

To use provider functions, declare the provider as a required provider in your Terraform configuration:

```hcl
terraform {
  required_providers {
    assert = {
      source = "hashicorp/assert"
    }
  }
}
```

## Continuous Validation

Simplify continuous validation checks that run as part of your Terraform workflow:

```hcl
data "http" "terraform_io" {
  url = "https://www.terraform.io"
}

check "health_check" {
  assert {
    condition     = provider::assert::http_success(data.http.terraform_io.status_code)
    error_message = "${data.http.terraform_io.url} returned an unhealthy status code"
  }
}
```

## Terraform Test

Test assertions in your Terraform configuration should be simple and easy to read:

```hcl
run "ebs_volume_size" {
  command = plan
  assert {
    condition     = provider::assert::between(1, 100, aws_ebs_volume.example.size)
    error_message = "EBS volume size must be between 1 and 100 GiB"
  }
}
```

## Variable Validation

Write simple validation rules for your Terraform variables:

```hcl
variable "ebs_volume_size" {
  type = number
  validation {
    condition     = provider::assert::between(1, 100, var.ebs_volume_size)
    error_message = "EBS volume size must be between 1 and 100 GiB"
  }
}
```

## License

[Mozilla Public License v2.0](./LICENSE)
