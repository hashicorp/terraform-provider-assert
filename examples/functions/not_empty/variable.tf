variable "example" {
  type = string

  validation {
    condition     = provider::assert::not_empty(var.example)
    error_message = "Variable 'example' must not be empty."
  }
}
