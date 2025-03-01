terraform {
  required_providers {
    statsig = {
      version = "~> 2.0.0"
      source  = "statsig-io/statsig"
    }
  }
}

resource "statsig_keys" "server_key" {
  description  = "A short description of what this server key is used for."
  type         = "SERVER"
  environments = ["production"]
}

resource "statsig_keys" "client_key" {
  description  = "A short description of what this client key is used for."
  type         = "CLIENT"
  environments = ["production"]
  scopes       = ["client_download_config_specs"]
}

resource "statsig_keys" "console_key" {
  description = "A short description of what this console key is used for."
  type        = "CONSOLE"
  scopes      = ["omni_read_only"]
}
