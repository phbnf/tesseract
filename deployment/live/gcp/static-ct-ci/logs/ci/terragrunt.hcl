terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//tesseract/conformance"
}

locals {
  env                             = "ci"
  docker_env                      = local.env
  base_name                       = "${local.env}-conformance"
  origin                          = "${local.base_name}.ct.transparency.dev"
  safe_origin                     = replace("${local.origin}", "/[^-a-zA-Z0-9]/", "-")
  log_public_key_secret_name      = "projects/223810646869/secrets/${local.safe_origin}-log-public/versions/1"
  log_private_key_secret_name     = "projects/223810646869/secrets/${local.safe_origin}-log-secret/versions/1"
  roots_remote_fetch_url          = "http://localhost:8080/roots.csv"
  roots_remote_fetch_interval     = "10s"
  roots_reject_fingerprints       = "e4286b40367091e7151f06d2164f9dea3472983e934a2b860bbc4909e070cd8a"
  server_docker_image             = "${include.root.locals.location}-docker.pkg.dev/${include.root.locals.project_id}/docker-${local.env}/conformance-gcp:latest"
  remote_root_server_docker_image = "${include.root.locals.location}-docker.pkg.dev/${include.root.locals.project_id}/docker-${local.env}/remote-root-server:latest"
  ephemeral                       = true
}

include "root" {
  path   = find_in_parent_folders("root.hcl")
  expose = true
}

inputs = merge(
  local,
  include.root.locals,
)
