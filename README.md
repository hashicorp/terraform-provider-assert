# Terraform Provider: Assert

The [Assert Terraform provider]((https://registry.terraform.io/providers/bschaatsbergen/assert/latest/docs)) is intended for use when writing [Terraform tests](https://developer.hashicorp.com/terraform/language/tests). It serves as a way to verify that the values in your Terraform configuration meet specific criteria. The provider only contains functions to assert values, and does not manage any resources.

## Examples

## Terraform Test

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

```hcl
variable "ebs_volume_size" {
  type = number
  validation {
    condition     = provider::assert::between(1, 100, var.ebs_volume_size)
    error_message = "EBS volume size must be between 1 and 100 GiB"
  }
}
```

## Documentation

Official documentation on how to use this provider can be found on the 
[Terraform Registry](https://registry.terraform.io/providers/bschaatsbergen/assert/latest/docs).
In case of specific questions or discussions, please [open an issue](https://github.com/bschaatsbergen/terraform-provider-assert/issues/new) in this repository.

## Contributing

Check out the [contributor guide](https://bschaatsbergen.github.io/terraform-provider-assert/).

## License

[Mozilla Public License v2.0](./LICENSE)
