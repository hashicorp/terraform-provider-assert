output "test" {
  value = provider::assert::contains(["a", "b", "c"], "b")
}