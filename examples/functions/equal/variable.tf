variable "number_of_glue_job_workers" {
  type = number
  validation {
    condition     = provider::assert::equal(var.number_of_glue_job_workers, 5)
    error_message = "Number of Glue job workers must be 5"
  }
}
