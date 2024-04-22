variable "ipv4_subnet" {
  type = string
  validation {
    condition     = provider::assert::cidrv4(var.ipv4_subnet)
    error_message = "Invalid IPv4 subnet"
  }
}
