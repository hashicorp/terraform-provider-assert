variable "ipv6_address" {
  type = number
  validation {
    condition     = provider::assert::ipv6(var.ipv6_address)
    error_message = "Invalid IPv6 address"
  }
}
