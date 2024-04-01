run "number_of_glue_job_workers_is_5" {

  command = plan

  assert {
    condition     = provider::assert::equal(aws_glue_job.example.number_of_workers, 5)
    error_message = "Number of Glue job workers must be 5"
  }
}
