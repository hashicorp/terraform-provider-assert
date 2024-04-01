output "test" {
  value = provider::assert::http_client_error(301)
}