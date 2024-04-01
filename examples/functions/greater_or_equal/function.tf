output "test" {
  value = provider::assert::greater_or_equal(1000000, 1000000)
}