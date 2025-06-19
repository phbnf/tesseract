terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//cloudbuild/preloader"
}

locals {
  start_index_offset = 300000 # Entries that did not make it into arche2025h1
}

include "root" {
  path   = find_in_parent_folders("root.hcl")
  expose = true
}

inputs = merge(
  local,
  include.root.locals,
)

