locals {
  yaml = yamlencode({
    foo = "bar"
  })
}
output "test" {
  value = provider::assert::valid_yaml(local.yaml)
}