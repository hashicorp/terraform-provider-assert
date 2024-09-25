run "check_http_server_error" {
  command = plan
  assert {
    condition     = provider::assert::http_client_error(data.http.down.status_code)
    error_message = "My down website must return an HTTP server error status code"
  }
}
