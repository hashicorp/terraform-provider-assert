run "check_security_group_description" {

  command = plan

  assert {
    condition     = provider::assert::not_empty(aws_security_group.example.description)
    error_message = "Description can not be empty"
  }
}