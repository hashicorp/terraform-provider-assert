# Terraform Provider: Assert

The [Assert Terraform provider]((https://registry.terraform.io/providers/bschaatsbergen/assert/latest/docs)) is intended for use when writing [Terraform tests](https://developer.hashicorp.com/terraform/language/tests). It serves as a way to verify that the values in your Terraform configuration meet specific criteria. The provider only contains functions to assert values, and does not manage any resources.

* [Terraform Registry](https://registry.terraform.io/providers/bschaatsbergen/assert/latest/docs)
* [Contributor Guide](https://bschaatsbergen.github.io/terraform-provider-assert/)

To use provider functions, declare the provider as a required provider in your Terraform configuration:

```hcl
terraform {
  required_providers {
    assert = {
      source = "bschaatsbergen/assert"
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

Simplify writing test assertions for your Terraform configuration:

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

Make writing validations for Terraform variables easier:

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
