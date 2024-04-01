output "test" {
  value = provider::assert::between(1, 10, 5)
}
