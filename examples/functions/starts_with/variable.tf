variable "webhook_url" {
  type = string
  validation {
    condition     = provider::assert::starts_with("https://", var.webhook_url)
    error_message = "Webhook URL must start with https://"
  }
}
