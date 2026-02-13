terraform {
  source = "${get_repo_root()}/deployment/modules/aws//tesseract/conformance"
}

include "root" {
  path   = find_in_parent_folders("root.hcl")
  expose = true
}

inputs = merge(
  include.root.locals,
  {
    roots_remote_fetch_url      = "http://localhost:8080/roots.csv"
    roots_remote_fetch_interval = "10s"
    roots_reject_fingerprints   = "e4286b40367091e7151f06d2164f9dea3472983e934a2b860bbc4909e070cd8a" # internal/testdata/test_root_ca_cert.pem
    # This hack makes it so that the antispam tables are created in the main
    # tessera DB. We strongly recommend that the antispam DB is separate, but
    # creating a second DB from OpenTofu is too difficult without a large
    # rewrite. For CI purposes, testing antispam, even if in the same DB, is
    # preferred compared to not testing antispam at all.
    antispam_database_name = "tesseract"
    create_antispam_db     = false
  }
)
