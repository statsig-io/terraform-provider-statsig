resource "statsig_keys" "client_key" {
  description  = "A short description of what this client key is used for."
  type         = "CLIENT"
  environments = ["production"]
  scopes       = ["client_download_config_specs"]
}

output "client_key" {
  value = statsig_keys.client_key
}
