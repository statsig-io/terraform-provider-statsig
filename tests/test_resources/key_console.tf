resource "statsig_keys" "console_key" {
  description = "A short description of what this console key is used for."
  type        = "CONSOLE"
  scopes      = ["omni_read_only"]
}

output "console_key" {
  value = statsig_keys.console_key
}
