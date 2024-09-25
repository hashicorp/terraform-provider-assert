run "check_https_google_pubsub_subscription_push_endpoint" {
  command = plan
  assert {
    condition     = provider::assert::starts_with("https://", google_pubsub_subscription.example.push_config.push_endpoint)
    error_message = "Push endpoint must start with https://"
  }
}
