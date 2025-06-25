variable "version" {
  type = string
  validation {
    condition     = provider::assert::semver_check("~> 1.0", var.version)
    error_message = "The provided version is not supported"
  }
}
