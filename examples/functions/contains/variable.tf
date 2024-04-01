variable "environment" {
  type = string
  validation {
    condition     = provider::assert::contains(["dev", "test", "prod"], var.environment)
    error_message = "Environment must be one of dev, test, or prod"
  }
}
