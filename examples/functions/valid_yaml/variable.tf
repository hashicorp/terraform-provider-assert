variable "config" {
  type = string
  validation {
    condition     = provider::assert::valid_yaml(var.config)
    error_message = "Config is not a valid YAML"
  }
}
