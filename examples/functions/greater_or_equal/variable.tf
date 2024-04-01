variable "db_instance_size" {
  type = number
  validation {
    condition     = provider::assert::greater_or_eqal(var.db_instance_size, 100)
    error_message = "DB instance size must be greater than or equal to 100"
  }
}
