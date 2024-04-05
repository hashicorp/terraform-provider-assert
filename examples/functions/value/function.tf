run "check_if_lambda_function_tags_has_value" {
  command = plan

  assert {
    condition     = provider::assert::value(aws_lambda_function.example.tags, "value1")
    error_message = "The tags map must contain the value 'value1'"
  }
}
