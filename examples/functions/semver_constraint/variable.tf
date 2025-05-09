variable "version_constraint" {
  type = string
  validation {
    condition     = provider::assert::semver_constraint("~> 1.0", var.version_constraint)
    error_message = "The provided version constraint is not valid"
  }
}
