data "http" "hashicorp" {
  url = "https://hashicorp.com"
}

output "is_redirected" {
  value = provider::assert::http_redirect(data.http.hashicorp.status_code)
}
