variable "dynamodb_range_key" {
  type    = string
  default = null
  validation {
    condition     = provider::assert::not_null(var.dynamodb_range_key)
    error_message = "DynamoDB table range key must not be null"
  }
}
