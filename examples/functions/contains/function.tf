run "check_example_subnet_cidr_block" {

  command = plan

  assert {
    condition     = provider::assert::contains(cidrsubnets("10.1.0.0/16", 4, 4, 8, 4), aws_subnet.example.cidr_block)
    error_message = "CIDR block is not in the list of allowed subnets"
  }
}
