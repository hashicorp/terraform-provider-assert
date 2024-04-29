variable "tags" {
  type = map(string)
  validation {
    condition     = provider::assert::value("value1", var.tags)
    error_message = "The tags map must contain the value 'value1'"
  }
}
