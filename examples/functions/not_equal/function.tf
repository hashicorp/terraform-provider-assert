run "check_rds_global_cluster_deletion_protection" {

  command = plan

  assert {
    condition     = provider::assert::not_equal(var.rds_global_cluster_deletion_protection, false)
    error_message = "Make sure that deletion protection is not set to false"
  }
}
