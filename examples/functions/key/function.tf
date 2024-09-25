run "check_if_lambda_function_tags_has_key" {
  command = plan
  assert {
    condition     = provider::assert::key("key1", aws_lambda_function.example.tags)
    error_message = "The tags map must contain the key 'key1'"
  }
}
