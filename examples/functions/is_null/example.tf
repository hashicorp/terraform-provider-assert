locals {
  person = null
}

output "is_null" {
  value = provider::assert::is_null(local.person)
}
