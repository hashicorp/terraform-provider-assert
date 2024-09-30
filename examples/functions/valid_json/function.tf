run "check_config" {
  command = plan
  assert {
    condition     = provider::assert::valid_json(data.local_file.config.content)
    error_message = "Config is not a valid JSON"
  }
}
