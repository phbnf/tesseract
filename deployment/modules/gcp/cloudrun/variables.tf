variable "project_id" {
  description = "GCP project ID where the log is hosted"
  type        = string
}

variable "base_name" {
  description = "Base name to use when naming resources"
  type        = string
}

variable "origin_suffix" {
  description = "Origin suffix, appended to base_name"
  type        = string
}

variable "location" {
  description = "Location in which to create resources"
  type        = string
}

variable "env" {
  description = "Unique identifier for the env, e.g. dev or ci or prod"
  type        = string
}

variable "server_docker_image" {
  description = "The full image URL (path & tag) for the Docker image to deploy in Cloud Run"
  type        = string
}

variable "bucket" {
  description = "Log GCS bucket"
  type        = string
}

variable "log_spanner_instance" {
  description = "Log Spanner instance"
  type        = string
}

variable "log_spanner_db" {
  description = "Log Spanner database"
  type        = string
}

variable "antispam_spanner_db" {
  description = "Antispam Spanner database"
  type        = string
}

variable "signer_public_key_secret_name" {
  description = "Public key secret name for checkpoints and SCTs signer. Format: projects/{projectId}/secrets/{secretName}/versions/{secretVersion}."
  type        = string
}

variable "signer_private_key_secret_name" {
  description = "Private key secret name for checkpoints and SCTs signer. Format: projects/{projectId}/secrets/{secretName}/versions/{secretVersion}."
  type        = string
}

variable "not_after_start" {
  description = "Start of the range of acceptable NotAfter values, inclusive. Leaving this empty implies no lower bound to the range. RFC3339 UTC format, e.g: 2024-01-02T15:04:05Z."
  default     = ""
  type        = string
}

variable "not_after_limit" {
  description = "Cut off point of notAfter dates - only notAfter dates strictly *before* notAfterLimit will be accepted. Leaving this empty means no upper bound on the accepted range. RFC3339 UTC format, e.g: 2024-01-02T15:04:05Z."
  default     = ""
  type        = string
}
