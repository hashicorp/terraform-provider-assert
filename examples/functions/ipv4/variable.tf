variable "ipv4_address" {
  type = string
  validation {
    condition     = provider::assert::ipv4(var.ipv4_address)
    error_message = "Invalid IPv4 address"
  }
}
