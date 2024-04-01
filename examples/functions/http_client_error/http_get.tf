data "http" "secure" {
  url = "https://my.secure.website.com"
}

output "is_redirected" {
  value = provider::assert::http_client_error(data.http.secure.status_code)
}
