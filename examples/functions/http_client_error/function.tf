run "check_http_client_error" {

  command = plan

  assert {
    condition     = provider::assert::http_client_error(data.http.secured.status_code)
    error_message = "My secure website must return an HTTP client error status code"
  }
}
