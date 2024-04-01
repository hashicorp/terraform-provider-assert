locals {
  data = null
}
output "test" {
  value = provider::assert::not_null(local.data)
}