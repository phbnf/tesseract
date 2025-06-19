terraform {
  backend "gcs" {}

  required_providers {
    google = {
      source  = "registry.terraform.io/hashicorp/google"
      version = "6.12.0"
    }
  }
}

# Cloud Build

locals {
  cloudbuild_service_account   = "cloudbuild-${var.env}-sa@${var.project_id}.iam.gserviceaccount.com"
  scheduler_service_account    = "scheduler-${var.env}-sa@${var.project_id}.iam.gserviceaccount.com"
}

resource "google_project_service" "cloudbuild_api" {
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

## Service usage API is required on the project to enable APIs.
## https://cloud.google.com/apis/docs/getting-started#enabling_apis
## serviceusage.googleapis.com acts as a central point for managing the API 
## lifecycle within your project. By ensuring the required APIs are enabled 
## and accessible, it allows Cloud Build to function seamlessly and interact 
## with other Google Cloud services as needed.
## 
## The Cloud Build service account also needs roles/serviceusage.serviceUsageViewer.
resource "google_project_service" "serviceusage_api" {
  service            = "serviceusage.googleapis.com"
  disable_on_destroy = false
}

resource "google_cloudbuild_trigger" "preloader_trigger" {
  name            = "preloader-${var.base_name}"
  service_account = "projects/${var.project_id}/serviceAccounts/${local.cloudbuild_service_account}"
  location        = var.location

  # TODO(phboneff): use a better mechanism to trigger releases that re-uses Docker containters, or based on branches rather.
  # This is a temporary mechanism to speed up development.
  github {
    owner = var.github_owner
    name  = "tesseract"
    push {
      tag = "^staging-deploy-preloader-(.+)$"
    }
  }

  build {
    ## TODO(phbnf): add the step that actually build the cotainer.

    ## TODO(phboneff): move to its own container / cloudrun / batch job.
    ## Preload entries.
    ## Leave enough time for the preloader to run, until the token expires.
    ## Stop after 360k entries, this is what gets copied within 60 minutes.
    timeout = "4200s" // 60 minutes
    step {
      id       = "start_index"
      name     = "golang"
      script   = <<EOT
	      echo $(($(curl -H "Authorization: Bearer $(cat /workspace/cb_access)" ${var.monitoring_url}/checkpoint | head -2 | tail -1)+${var.start_index_offset})) > /workspace/start_index
	      echo "Will run preloader from $(cat /workspace/start_index)."
      EOT
      wait_for = ["bearer_token"]
    }

    ## Apply the deployment/live/gcp/static-staging/preloader/XXX terragrunt config.
    ## This will bring up or update TesseraCT's infrastructure, including a service
    ## running the server docker image built above.

   step {
      id     = "terraform_apply_preloader_${tg_path.key}"
      name   = "alpine/terragrunt"
      script = <<EOT
        terragrunt --terragrunt-non-interactive --terragrunt-no-color apply -auto-approve -no-color -var "start_index=$(cat /workspace/start_index)" 2>&1
      EOT
      dir    = tg_path.value
      env = [
        "GOOGLE_PROJECT=${var.project_id}",
        "TF_IN_AUTOMATION=1",
        "TF_INPUT=false",
        "TF_VAR_project_id=${var.project_id}",
        "DOCKER_CONTAINER_TAG=$SHORT_SHA"
      ]
      wait_for = ["start_index"]
    }

    options {
      logging      = "CLOUD_LOGGING_ONLY"
      machine_type = "E2_HIGHCPU_8"
    }
  }
}
