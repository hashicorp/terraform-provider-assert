run "number_of_glue_job_workers_is_5" {

  command = plan

  assert {
    condition     = provider::assert::equal(5, aws_glue_job.example.number_of_workers)
    error_message = "Number of Glue job workers must be 5"
  }
}
