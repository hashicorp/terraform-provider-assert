variable "example" {
  type = string

  validation {
    condition     = provider::assert::not_empty(var.example)
    error_message = "Value must not be empty"
  }
}
