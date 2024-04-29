variable "tags" {
  type = map(string)
  validation {
    condition     = provider::assert::key("key1", var.tags)
    error_message = "The tags map must contain the key 'key1'"
  }
}
