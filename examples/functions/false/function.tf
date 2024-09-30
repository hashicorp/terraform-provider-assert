run "check_rds_global_cluster_deletion_protection" {
  command = plan
  assert {
    condition     = provider::assert::false(aws_rds_global_cluster.example.deletion_protection)
    error_message = "Cluster deletion protection must be false, this is a dev environment"
  }
}
