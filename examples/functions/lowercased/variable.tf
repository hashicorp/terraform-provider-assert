variable "example" {
  type = string
  validation {
    condition     = provider::assert::lowercased(var.example)
    error_message = "Example must be lowercased"
  }
}
