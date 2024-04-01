output "test" {
  value = provider::assert::http_redirect(301)
}