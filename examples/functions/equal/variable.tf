variable "number_of_glue_job_workers" {
  type = number
  validation {
    condition     = provider::assert::equal(5, var.number_of_glue_job_workers)
    error_message = "Number of Glue job workers must be 5"
  }
}
