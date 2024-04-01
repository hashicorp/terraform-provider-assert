variable "rds_global_cluster_deletion_protection" {
  type = bool
  validation {
    condition     = provider::assert::not_equal(var.rds_global_cluster_deletion_protection, false)
    error_message = "Make sure that deletion protection is not set to false"
  }
}
