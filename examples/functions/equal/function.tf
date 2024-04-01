output "test" {
  value = provider::assert::equal(1000000, 1000000)
}