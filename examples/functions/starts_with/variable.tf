variable "webhook_url" {
  type = string
  validation {
    condition     = provider::assert::starts_with(var.webhook_url, "https://")
    error_message = "Webhook URL must start with https://"
  }
}
