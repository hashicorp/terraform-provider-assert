run "check_valid_ip_google_compute_address" {

  command = plan

  assert {
    condition     = provider::assert::ip(google_compute_address.example.address)
    error_message = "Address is not a valid IP address"
  }
}
