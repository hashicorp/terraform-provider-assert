variable "example" {
  type = object({
    name  = string
    value = optional(string)
  })
  validation {
    condition     = provider::assert::null(var.example.value)
    error_message = "Value must be null"
  }
}
