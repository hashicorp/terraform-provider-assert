run "check_if_lambda_function_tags_has_key" {
  command = plan

  assert {
    condition     = provider::assert::key(aws_lambda_function.example.tags, "key1")
    error_message = "The tags map must contain the key 'key1'"
  }
}
