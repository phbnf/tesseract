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

variable "target_log_uri" {
  description = "URL of the destination log to preload to."
  type        = string
}

variable "source_log_uri" {
  description = "URL of the source RFC6962 log."
  type        = string
}

variable "start_index" {
  description = "Index to start preloading from on the source log."
  type        = number
}
