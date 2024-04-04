variable "example" {
  type = number
  validation {
    condition     = provider::assert::positive(var.example)
    error_message = "Expected ${var.example} to be positive"
  }
}
