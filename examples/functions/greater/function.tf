output "test" {
  value = provider::assert::greater(500, 200)
}