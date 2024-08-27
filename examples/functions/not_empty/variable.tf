variable "bucket_name" {
  type = string

  validation {
    condition     = provider::assert::not_empty(var.bucket_name)
    error_message = "Bucket name must not be empty"
  }
}
