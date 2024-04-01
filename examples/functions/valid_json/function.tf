locals {
  json = jsonencode({
    foo = "bar"
  })
}
output "test" {
  value = provider::assert::valid_json(local.json)
}