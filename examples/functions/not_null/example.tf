locals {
  person = {
    name = "John Doe"
    age  = 30
  }
}

output "example" {
  value = provider::assert::not_null(local.person)
}
