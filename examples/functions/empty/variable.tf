variable "example" {
  type = string

  validation {
    condition     = provider::assert::empty(var.example)
    error_message = "Variable 'example' must be empty."
  }
}
