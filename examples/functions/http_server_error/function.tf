output "test" {
  value = provider::assert::http_server_error(504)
}