resource "google_compute_managed_ssl_certificate" "example" {
  name = "example"

  managed {
    domains = ["example.com"]
  }
}

check "certificate_valid" {
  assert {
    // Add 336 hours (14 days) to the expiration time, making sure we have enough time to renew the certificate
    condition     = !provider::assert::expired(timeadd(google_compute_managed_ssl_certificate.example.expire_time, "336h"))
    error_message = "Example certificate needs to be renewed"
  }
}
