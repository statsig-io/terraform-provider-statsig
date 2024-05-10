resource "statsig_keys" "server_key" {
  description  = "A short description of what this server key is used for."
  type         = "SERVER"
  environments = ["production"]
}

output "server_key" {
  value = statsig_keys.server_key
}
