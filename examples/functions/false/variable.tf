variable "rds_global_cluster_deletion_protection" {
  type = bool
  validation {
    condition     = provider::assert::false(var.rds_global_cluster_deletion_protection)
    error_message = "Cluster deletion protection must be false, this is a dev environment"
  }
}
