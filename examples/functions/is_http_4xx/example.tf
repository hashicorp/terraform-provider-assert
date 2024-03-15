data "http" "example" {
  url = "https://checkpoint-api.hashicorp.com/v1/check/terraform"

  # Optional request headers
  request_headers = {
    Accept = "application/json"
  }
}

output "http_client_error" {
  value = provider::assert::http_client_error(data.http.example.status_code)
}
