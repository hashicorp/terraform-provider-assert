variable "db_instance_size" {
  type = number
  validation {
    condition     = provider::assert::less_or_equal(1000, var.db_instance_size)
    error_message = "DB instance size must be less than or equal to 1000"
  }
}
