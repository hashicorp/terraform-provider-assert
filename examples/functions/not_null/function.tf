run "check_if_range_key_is_not_null" {

  command = plan

  assert {
    condition     = provider::assert::not_null(aws_dynamodb_table.example.range_key)
    error_message = "DynamoDB table range key must not be null"
  }
}
