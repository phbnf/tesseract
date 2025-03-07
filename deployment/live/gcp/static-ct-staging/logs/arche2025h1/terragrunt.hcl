terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//tesseract/conformance"
}

locals {
  docker_env          = "staging"
  base_name           = "arche2025h1"
  origin_suffix       = ".ct.transparency.dev"
  server_docker_image = "us-central1-docker.pkg.dev/${include.root.locals.project_id}/docker-staging/conformance-gcp:latest"
}

include "root" {
  path   = find_in_parent_folders()
  expose = true
}

inputs = merge(
  local,
  include.root.locals,
)
