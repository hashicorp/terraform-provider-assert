variable "ip_address" {
  type = number
  validation {
    condition     = provider::assert::ip(var.ip_address)
    error_message = "Invalid IP address"
  }
}
