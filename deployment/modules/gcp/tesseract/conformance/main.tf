terraform {
  backend "gcs" {}
}

module "storage" {
  source = "../../storage"

  project_id = var.project_id
  base_name  = var.base_name
  location   = var.location
  ephemeral  = var.ephemeral
  spanner_pu = var.spanner_pu
}

module "cloudrun" {
  source = "../../cloudrun"

  env                            = var.env
  project_id                     = var.project_id
  base_name                      = var.base_name
  origin                         = var.origin
  location                       = var.location
  server_docker_image            = var.server_docker_image
  not_after_start                = var.not_after_start
  not_after_limit                = var.not_after_limit
  bucket                         = module.storage.log_bucket.id
  log_spanner_instance           = module.storage.log_spanner_instance.name
  log_spanner_db                 = module.storage.log_spanner_db.name
  antispam_spanner_db            = module.storage.antispam_spanner_db.name
  signer_public_key_secret_name  = var.log_public_key_secret_name
  signer_private_key_secret_name = var.log_private_key_secret_name
  trace_fraction                 = var.trace_fraction
  batch_max_age                  = var.batch_max_age
  batch_max_size                 = var.batch_max_size
  roots_remote_fetch_url         = var.roots_remote_fetch_url
  roots_remote_fetch_interval    = var.roots_remote_fetch_interval
  roots_reject_fingerprints      = var.roots_reject_fingerprints

  additional_containers = [
    {
      name  = "remote-root-server"
      image = var.remote_root_server_docker_image
      args  = [
        "--http_endpoint=:8080",
        "--tesseract_url=http://localhost:6962",
        "--exit_on_success=true",
        "--verify_interval=5s",
        "--max_runtime=3m",
        "--v=1"
      ]
      ports = [{
        container_port = 8080
      }]
      resources = {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
      }
    }
  ]

  depends_on = [
    module.storage
  ]
}
