data "http" "down" {
  url = "https://my.down.website.com"
}

output "is_redirected" {
  value = provider::assert::http_server_error(data.http.down.status_code)
}
