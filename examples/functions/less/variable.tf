variable "db_instance_size" {
  type = number
  validation {
    condition     = provider::assert::less(var.db_instance_size, 1000)
    error_message = "DB instance size must be less than 1000"
  }
}
