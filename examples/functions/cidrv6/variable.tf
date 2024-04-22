variable "ipv6_subnet" {
  type = string
  validation {
    condition     = provider::assert::cidrv6(var.ipv6_subnet)
    error_message = "Invalid IPv6 subnet"
  }
}
