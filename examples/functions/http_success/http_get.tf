data "http" "hashicorp" {
  url = "https://developer.hashicorp.com"
}

output "is_redirected" {
  value = provider::assert::http_success(data.http.hashicorp.status_code)
}
