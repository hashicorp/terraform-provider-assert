variable "ip_subnet" {
  type = string
  validation {
    condition     = provider::assert::cidr(var.ip_subnet)
    error_message = "Invalid IP subnet"
  }
}
