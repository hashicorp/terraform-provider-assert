variable "db_instance_size" {
  type = number
  validation {
    condition     = provider::assert::greater(100, var.db_instance_size)
    error_message = "DB instance size must be greater than 100"
  }
}
