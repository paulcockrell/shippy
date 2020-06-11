// General variables

variable "linux_admin_username" {
    type = "string"
    description = "User name for authentication to the Kubernetes linux agent virtual machines in the cluster"
}

variable "linux_admin_password" {
    type = "string"
    description = "Password for the linux admin account"
}

// GCP variables

variable "gcloud_project" {
    type = "string"
    description = "GCP project ID (not project name!)"
}

variable "gcloud_region" {
    type = "string"
    description = "GCP project region"
}
variable "gcloud_zone" {
    type = "string"
    description = "GCP project zone"
}
variable "platform_name" {
    type = "string"
    description = "Platform name"
}