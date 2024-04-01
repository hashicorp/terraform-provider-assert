locals {
  obj = {
    foo = "Foo"
    bar = "Bar"
  }
}
output "test" {
  value = provider::assert::not_null(local.obj)
}