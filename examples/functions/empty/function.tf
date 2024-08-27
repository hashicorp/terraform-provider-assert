run "check_cloudwatch_log_subscription_match_all" {

  command = plan

  assert {
    condition     = provider::assert::empty(aws_cloudwatch_log_subscription_filter.example.filter_pattern)
    error_message = "Filter pattern must be empty to match all log events"
  }
}
