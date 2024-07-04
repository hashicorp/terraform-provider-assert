resource "aws_acm_certificate" "example" {
  domain_name       = "example.com"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

check "certificate_valid" {
  assert {
    // Add 336 hours (14 days) to the expiration time, making sure we have enough time to renew the certificate
    condition     = !provider::assert::expired(timeadd(aws_acm_certificate.example.not_after, "336h"))
    error_message = "Example certificate needs to be renewed"
  }
}
