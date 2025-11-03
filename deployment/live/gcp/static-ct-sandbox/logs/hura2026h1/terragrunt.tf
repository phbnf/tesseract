module "log" {
  source = "../../../../../../deployment/modules/gcp/tesseract/gce"

  env                                        = "sandbox"
  docker_env                                 = "sandbox"
  base_name                                  = "hura2026h1"
  project_id                                 = "static-ct-sandbox"
  location                                   = "us-central1"
  origin_suffix                              = ".sandbox.ct.transparency.dev"
  not_after_start                            = "2026-01-01T00:00:00Z"
  not_after_limit                            = "2026-07-01T00:00:00Z"
  server_docker_image                        = "us-central1-docker.pkg.dev/static-ct-sandbox/docker-sandbox/tesseract-gcp:latest"
  spanner_pu                                 = 500
  trace_fraction                             = 0.1
  create_internal_load_balancer              = true
  public_bucket                              = true
  rate_limit_old_not_before                  = "28h:150"
  additional_signer_private_key_secret_names = [
  //  "projects/781477119959/secrets/arche2026h1-ed25519-private-key/versions/1"
  ]
}

