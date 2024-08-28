variable "service_account_email" {
  type = string
  validation {
    condition     = provider::assert::regex("^[^@]+@[^@]+\.[^@]+$", var.webhook_url)
    error_message = "Service account email must be in the format of <name>@<domain>.<tld>"
  }
}
