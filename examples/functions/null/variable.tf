variable "example" {
  type = string
  validation {
    condition     = provider::assert::null(var.example) || length(var.example_variable) > 0
    error_message = "Must either be a non-empty string or null"
  }
}
