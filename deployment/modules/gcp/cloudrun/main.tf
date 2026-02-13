terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.50.0"
    }
  }
}

# Cloud Run

locals {
  cloudrun_service_account_id = var.env == "" ? "cloudrun-sa" : "cloudrun-${var.env}-sa"
  spanner_log_db_path         = "projects/${var.project_id}/instances/${var.log_spanner_instance}/databases/${var.log_spanner_db}"
  spanner_antispam_db_path    = "projects/${var.project_id}/instances/${var.log_spanner_instance}/databases/${var.antispam_spanner_db}"
}

resource "google_project_service" "cloudrun_api" {
  service            = "run.googleapis.com"
  disable_on_destroy = false
}

resource "google_cloud_run_v2_service" "default" {
  name         = var.base_name
  location     = var.location
  launch_stage = "GA"

  template {
    service_account                  = "${local.cloudrun_service_account_id}@${var.project_id}.iam.gserviceaccount.com"
    max_instance_request_concurrency = 1000
    timeout                          = "30s"

    scaling {
      max_instance_count = 2
      min_instance_count = 2
    }

    containers {
      image = var.server_docker_image
      name  = "tesseract"
      args = flatten([
        "--logtostderr",
        "--v=1",
        "--http_endpoint=:6962",
        "--bucket=${var.bucket}",
        "--spanner_db_path=${local.spanner_log_db_path}",
        "--spanner_antispam_db_path=${local.spanner_antispam_db_path}",
        "--roots_pem_file=/bin/test_root_ca_cert.pem",
        formatlist("--roots_reject_fingerprints=%s", var.roots_reject_fingerprints),
        "--origin=${var.origin}",
        "--path_prefix=${var.origin}",
        "--signer_public_key_secret_name=${var.signer_public_key_secret_name}",
        "--signer_private_key_secret_name=${var.signer_private_key_secret_name}",
        "--inmemory_antispam_cache_size=256k",
        "--not_after_start=${var.not_after_start}",
        "--not_after_limit=${var.not_after_limit}",
        "--trace_fraction=${var.trace_fraction}",
        "--batch_max_size=${var.batch_max_size}",
        "--batch_max_age=${var.batch_max_age}",
        "--roots_remote_fetch_url=${var.roots_remote_fetch_url}",
        "--roots_remote_fetch_interval=${var.roots_remote_fetch_interval}",
      ])
      ports {
        container_port = 6962
      }

      resources {
        limits = {
          cpu    = "8"
          memory = "8Gi"
        }
      }

      startup_probe {
        initial_delay_seconds = 1
        timeout_seconds       = 1
        period_seconds        = 10
        failure_threshold     = 30
        tcp_socket {
          port = 6962
        }
      }
    }

    dynamic "containers" {
      for_each = var.additional_containers
      content {
        name  = containers.value.name
        image = containers.value.image
        args  = containers.value.args
        dynamic "ports" {
          for_each = containers.value.ports != null ? containers.value.ports : []
          content {
            container_port = ports.value.container_port
            name           = ports.value.name
          }
        }
        dynamic "resources" {
          for_each = containers.value.resources != null ? [containers.value.resources] : []
          content {
            limits = resources.value.limits
          }
        }
        dynamic "env" {
          for_each = containers.value.env != null ? containers.value.env : []
          content {
            name  = env.value.name
            value = env.value.value
          }
        }
        dynamic "volume_mounts" {
          for_each = containers.value.volume_mounts != null ? containers.value.volume_mounts : []
          content {
            name       = volume_mounts.value.name
            mount_path = volume_mounts.value.mount_path
          }
        }
      }
    }
  }

  deletion_protection = false

  client = "terraform"

  depends_on = [
    google_project_service.cloudrun_api,
  ]
}
