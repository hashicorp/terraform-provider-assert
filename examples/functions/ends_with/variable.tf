variable "webhook_url" {
  type = string
  validation {
    condition     = provider::assert::ends_with(var.webhook_url, "/events")
    error_message = "Webhook URL must end with /events"
  }
}
