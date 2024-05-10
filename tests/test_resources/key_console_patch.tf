variable "key" {
  type = string
}

resource "statsig_keys" "console_key" {
  description = "A short description of what this console key is used for."
  type        = "CONSOLE"
  scopes      = ["omni_read_write"]
  key         = var.key
}

output "console_key" {
  value = statsig_keys.console_key
}
