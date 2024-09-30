run "check_valid_ip_aws_subnet" {
  command = plan
  assert {
    condition     = provider::assert::cidr(aws_subnet.example.cidr_block)
    error_message = "Subnet is not in valid CIDR notation"
  }
}
