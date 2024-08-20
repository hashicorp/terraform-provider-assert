run "check_https_google_pubsub_subscription_push_endpoint" {

  command = plan

  assert {
      condition     = provider::assert::regex_matches("^[^@]+@[^@]+\.[^@]+$", var.webhook_url)
      error_message = "Service account email must be in the format of <name>@<domain>.<tld>"
  }
}
