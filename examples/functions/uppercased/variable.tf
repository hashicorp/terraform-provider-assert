variable "example" {
  type = string
  validation {
    condition     = provider::assert::uppercased(var.example)
    error_message = "Example must be uppercased"
  }
}
