locals {
  person = null
}

output "null" {
  value = provider::assert::null(local.person)
}
