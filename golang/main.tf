
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

# google_container_cluster　のところが、Kubernetesにあたる。
# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster

resource "google_container_cluster" "primary" {
  name               = "go-app-gke"
  project            = "wed-practice"
  location           = "asia-northeast1"
  initial_node_count = 2
}

provider "google" {
  credentials = file("wed-practice-381bc81c536a.json")

  project = "wed-practice"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}