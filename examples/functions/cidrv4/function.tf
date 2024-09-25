run "check_valid_ipv4_aws_subnet" {
  command = plan
  assert {
    condition     = provider::assert::cidrv4(aws_subnet.example.cidr_block)
    error_message = "Subnet is not in valid IPv4 CIDR notation"
  }
}
