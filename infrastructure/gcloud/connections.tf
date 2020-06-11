provider "google" {
    credentials = "${file("google-cred.json")}"
    project = "${var.gcloud_project}"
    region = "${var.gcloud_region}"
}