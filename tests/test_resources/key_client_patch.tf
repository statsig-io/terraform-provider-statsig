variable "key" {
  type = string
}

resource "statsig_keys" "client_key" {
  description  = "A short description of what this client key is used for."
  type         = "CLIENT"
  environments = ["production", "staging"]
  scopes       = ["client_download_config_specs"]
  key          = var.key
}

output "client_key" {
  value = statsig_keys.client_key
}
