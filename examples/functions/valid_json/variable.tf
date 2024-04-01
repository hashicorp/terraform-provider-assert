variable "config" {
  type = string
  validation {
    condition     = provider::assert::valid_json(var.config)
    error_message = "Config is not a valid JSON"
  }
}
