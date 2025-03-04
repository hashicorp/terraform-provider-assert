variable "version" {
  type = string
  validation {
    condition     = provider::assert::semver_version(var.version)
    error_message = "The provided version is not a valid SemVer version"
  }
}
