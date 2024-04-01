variable "rds_global_cluster_deletion_protection" {
  type = bool
  validation {
    condition     = provider::assert::true(var.rds_global_cluster_deletion_protection)
    error_message = "Cluster deletion protection must be enabled, because this is a prod environment"
  }
}
