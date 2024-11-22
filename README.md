# Terraform Provider: Assert

The [Assert Terraform provider]((https://registry.terraform.io/providers/hashicorp/assert/latest/docs)) is intended for use when writing [Terraform Tests](https://developer.hashicorp.com/terraform/language/tests), [Variable Validation](https://developer.hashicorp.com/terraform/language/values/variables#custom-validation-rules), [Preconditions and Postconditions](https://developer.hashicorp.com/terraform/language/expressions/custom-conditions#preconditions-and-postconditions), or [Continuous Validation](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/health#continuous-validation). It serves as a way to verify that the values in your Terraform configuration meet specific criteria. The provider only contains functions to assert values, and does not manage any resources.

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

## Contributing

This provider is a HashiCorp utility provider, which means any bug fix and feature has to be considered in the context of the thousands/millions of configurations in which this provider is used. This is great as your contribution can have a big positive impact, but we have to assess potential negative impact too (e.g. breaking existing configurations). Stability over features.

To provide some safety to the wider provider ecosystem, we strictly follow semantic versioning and HashiCorp's own versioning specification. Any changes that could be considered as breaking will only be included as part of a major release. In case multiple breaking changes need to happen, we will group them in the next upcoming major release.

If you’re looking to contribute, thank you for investing your time and energy into this project! Please make sure you’re familiar with the [HashiCorp Code of Conduct](https://www.hashicorp.com/community-guidelines) and the Assert Provider [Contributor Guide](https://hashicorp.github.io/terraform-provider-assert/).

## License

[Mozilla Public License v2.0](./LICENSE)
