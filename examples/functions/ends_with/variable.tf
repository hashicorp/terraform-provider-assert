variable "webhook_url" {
  type = string
  validation {
    condition     = provider::assert::ends_with("/events", var.webhook_url)
    error_message = "Webhook URL must end with /events"
  }
}
