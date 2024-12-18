variable "time" {
  type        = bool
  description = "Check whether the current time matches the provider time slot"
  validation {
    condition     = var.time == provider::assert::timecheck(["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"], ["ALL"], "America/New_York", "09:00", "17:00")
    error_message = "true"
  }
  default = true
}
