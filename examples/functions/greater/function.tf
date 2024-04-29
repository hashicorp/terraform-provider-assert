run "check_aws_db_instance_size" {

  command = plan

  assert {
    condition     = provider::assert::greater(100, aws_db_instance.example.instance_class)
    error_message = "DB instance size must be greater than 100"
  }
}
