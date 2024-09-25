run "check_http_success" {
  command = plan
  assert {
    condition     = provider::assert::http_success(data.http.hashicorp.status_code)
    error_message = "HashiCorp's website must return a 2xx status code"
  }
}
