output "example" {
  value = provider::assert::within_range(1, 10, 4)
}
