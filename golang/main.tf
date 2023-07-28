
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.51.0"
    }
  }
}


data "google_client_config" "default" {}





data "google_container_cluster" "helloworld-gke" {
  name     = "helloworld-gke"
  project            = "wed-practice"
  location = "asia-northeast1"
}

provider "google" {
  credentials = file("wed-practice-381bc81c536a.json")

  project = "wed-practice"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

provider "kubernetes" {
  host                   = "https://${data.google_container_cluster.helloworld-gke.endpoint}"
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(data.google_container_cluster.helloworld-gke.master_auth[0].cluster_ca_certificate)
}

resource "kubernetes_namespace" "kn" {
  metadata {
    name = "go-app-namespace"
  }
}



resource "kubernetes_deployment" "kd" {
  metadata {
    name      = "go-app"
    namespace = kubernetes_namespace.kn.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "MyTestApp"
      }
    }
    template {
      metadata {
        labels = {
          app = "MyTestApp"
        }
      }
      spec {
        container {
          image = "asia-northeast1-docker.pkg.dev/wed-practice/go-app-repo/docker-gs-ping:latest"
          name  = "docker-gs-ping"
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "test" {
  metadata {
    name      = "go-app"
    namespace = kubernetes_namespace.kn.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.kd.spec.0.template.0.metadata.0.labels.app
    }
    type = "LoadBalancer"
    port {
      port        = 80
      target_port = 8080
    }
  }
}
  
