variable "example" {
  type = number
  validation {
    condition     = provider::assert::negative(var.example)
    error_message = "Expected ${var.example} to be negative"
  }
}
