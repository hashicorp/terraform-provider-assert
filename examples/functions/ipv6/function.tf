run "check_valid_ipv6_google_compute_address" {

  command = plan

  assert {
    condition     = provider::assert::ipv6(google_compute_address.example.address)
    error_message = "Address is not a valid IPv6 address"
  }
}
