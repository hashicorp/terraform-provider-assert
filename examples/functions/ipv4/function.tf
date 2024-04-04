run "check_valid_ipv4_google_compute_address" {

  command = plan

  assert {
    condition     = provider::assert::ipv4(google_compute_address.example.address)
    error_message = "Address is not a valid IPv4 address"
  }
}
