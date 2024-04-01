variable "db_instance_size" {
  type = number
  validation {
    condition     = provider::assert::greater(var.db_instance_size, 100)
    error_message = "DB instance size must be greater than 100"
  }
}
