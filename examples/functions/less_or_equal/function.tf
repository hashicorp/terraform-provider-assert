run "check_aws_db_instance_size" {

  command = plan

  assert {
    condition     = provider::assert::less_or_equal(1000, aws_db_instance.example.instance_class)
    error_message = "DB instance size must be less than or equal to 1000"
  }
}
