run "check_valid_ipv6_aws_subnet" {

  command = plan

  assert {
    condition     = provider::assert::cidrv6(aws_subnet.example.ipv6_cidr_block)
    error_message = "Subnet is not in valid IPv6 CIDR notation"
  }
}
