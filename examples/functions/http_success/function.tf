output "test" {
  value = provider::assert::http_success(200)
}