run "check_rds_global_cluster_deletion_protection" {
  command = plan
  assert {
    condition     = provider::assert::true(aws_rds_global_cluster.example.deletion_protection)
    error_message = "Cluster deletion protection must be enabled, because this is a prod environment"
  }
}
