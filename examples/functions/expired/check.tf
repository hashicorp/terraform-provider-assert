resource "google_compute_managed_ssl_certificate" "example" {
  name = "example"

  managed {
    domains = ["example.com"]
  }
}

check "certificate_valid" {
  assert {
    condition     = provider::assert::expired(timeadd(google_compute_managed_ssl_certificate.example.expire_time, "7d"))
    error_message = "Certificate needs to be renewed"
  }
}
