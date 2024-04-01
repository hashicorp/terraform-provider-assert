run "check_http_redirect" {

  command = plan

  assert {
    condition     = provider::assert::http_redirect(data.http.hashicorp.status_code)
    error_message = "HashiCorp's website must return a 3xx status code, when using HTTP instead of HTTPS"
  }
}
