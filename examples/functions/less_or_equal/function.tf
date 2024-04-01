output "test" {
  value = provider::assert::less(100, 100)
}