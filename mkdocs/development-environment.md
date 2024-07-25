# Development Environment Setup

If running tests and acceptance tests isn't enough, it's possible to set up a local terraform configuration to use a development builds of the provider. This can be achieved by leveraging the Terraform CLI configuration file development overrides.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.8+ (to run acceptance tests)
- [Go](https://golang.org/doc/install) >=1.21 (to build the provider plugin)

### Building the Provider

To place a fresh development build of the provider in your `${GOBIN}` (defaults to `${GOPATH}/bin` or `${HOME}/go/bin` if `${GOPATH}` is not set), run:

```console
make build
```

This will build the provider and put the provider binary in the `${GOBIN}` directory.

```console
ls -la ./$GOPATH/bin/terraform-provider-assert
```

### Testing the Provider

In order to test the provider, you can run:

```console
make test
```

### Using the Provider

With Terraform 1.8 and later, [development overrides for provider developers](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers) can be leveraged in order to use the provider built from source.

To do this, populate a Terraform CLI configuration file (`~/.terraformrc` for all platforms other than Windows; `terraform.rc` in the `%APPDATA%` directory when using Windows) with at least the following options:

```terraform
provider_installation {
  dev_overrides {
    "hashicorp/assert" = "[REPLACE WITH GOPATH]/bin"
  }
  direct {}
}
```
