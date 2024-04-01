variable "db_instance_size" {
  type = number
  validation {
    condition     = provider::assert::less_or_equal(var.db_instance_size, 1000)
    error_message = "DB instance size must be less than or equal to 1000"
  }
}
