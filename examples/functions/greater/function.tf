run "check_aws_db_instance_size" {

  command = plan

  assert {
    condition     = provider::assert::greater(aws_db_instance.example.instance_class, 100)
    error_message = "DB instance size must be greater than 100"
  }
}
