variable "tags" {
  type = map(string)
  validation {
    condition     = provider::assert::key(var.tags, "key1")
    error_message = "The tags map must contain the key 'key1'"
  }
}
