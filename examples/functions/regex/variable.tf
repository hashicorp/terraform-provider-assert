variable "user_email" {
  description = "The email address for the user."
  type        = string

  validation {
    condition     = provider::assert::regex("^[^@]+@[^@]+\\.[^@]+$", var.user_email)
    error_message = "The user email must be in a valid format, such as 'name@domain.tld'."
  }
}
