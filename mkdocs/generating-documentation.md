# Generating Documentation

This provider uses [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs/) to generate documentation and store it in the [docs/](https://github.com/hashicorp/terraform-provider-assert/tree/main/docs) directory. Once a release is cut, the Terraform Registry will download the documentation from [docs/](https://github.com/hashicorp/terraform-provider-assert/tree/main/docs) and associate it with the release version. Read more about how this works on the official page.

### Adding documentation for a new function

To add documentation for a new function, you need to do 2 things:

* Add a `<function>.md.tmpl` file in the [templates/](https://github.com/hashicorp/terraform-provider-assert/tree/main/templates) directory. This file is used by [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs/) to generate the documentation. You can refer to one of the existing files to create a new one.

* Add examples to the [examples/](https://github.com/hashicorp/terraform-provider-assert/tree/main/examples) directory. These examples are referenced in the template files and are used to generate the examples section in the documentation. We recommend adding an example for Terraform Test and Variable Validation. You can refer to one of the existing examples to create a new one.


### Generating documentation

After adding a new function to the [templates/](https://github.com/hashicorp/terraform-provider-assert/tree/main/templates) and [examples/](https://github.com/hashicorp/terraform-provider-assert/tree/main/examples) directories, run:

```sh
make generate
```
