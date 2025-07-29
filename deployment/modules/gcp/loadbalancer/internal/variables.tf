variable "project_id" {
  description = "GCP project ID where the log is hosted."
  type        = string
}

variable "base_name" {
  description = "Base name to use when naming resources."
  type        = string
}

variable "location" {
  description = "Location in which to create resources."
  type        = string
}

variable "env" {
  description = "Unique identifier for the env, e.g. dev or ci or prod."
  type        = string
}

variable "backend_group" {
  description = "Backend group to wire the load balancer to."
  type        = string
}
