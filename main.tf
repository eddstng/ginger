# Use the local kubeconfig file to connect to your k3d cluster
provider "kubernetes" {
  config_path = "~/.kube/config"
}

# Define the Deployment for ginger-api
resource "kubernetes_deployment" "ginger_api" {
  metadata {
    name = "ginger-api" # Name of the deployment
    labels = {
      app = "ginger-api" # Label used for selector and service matching
    }
  }

  spec {
    replicas = 1 # Run one instance (pod) of this service

    selector {
      match_labels = {
        app = "ginger-api" # Select pods with this label
      }
    }

    template {
      metadata {
        labels = {
          app = "ginger-api" # Assign this label to the pod
        }
      }

      spec {
        container {
          name  = "ginger-api" # Name of the container in the pod
          image = "ginger-api:latest" # Local image name (not from Docker Hub)
          image_pull_policy = "Never" # Tell K8s not to try pulling from a registry

          port {
            container_port = 3000 # Port the app runs on inside the container
          }

          env {
            name  = "DATABASE_URL" # Environment variable for DB connection
            value = "postgres://postgres:postgres@psql:5432/postgres?sslmode=disable"
          }

          resources {
            limits = {
              cpu = "500m" # Max CPU allowed
            }
            requests = {
              cpu = "100m" # Guaranteed CPU reservation
            }
          }
        }

        restart_policy = "Always" # Automatically restart the pod if it crashes
      }
    }
  }
}

# Define a Service to expose the ginger-api Deployment
resource "kubernetes_service" "ginger_api" {
  metadata {
    name = "ginger-api" # Name of the service
  }

  spec {
    selector = {
      app = "ginger-api" # Match pods with this label (from Deployment)
    }

    port {
      name       = "http" # Optional port name
      port       = 3000   # Service port (inside the cluster)
      target_port = 3000  # Pod container port to forward to
    }
  }
}

resource "kubernetes_horizontal_pod_autoscaler_v2" "ginger_api_hpa" {
  metadata {
    name      = "ginger-api-hpa"
    namespace = "default"
  }

  spec {
    scale_target_ref {
      kind        = "Deployment"
      name        = kubernetes_deployment.ginger_api.metadata[0].name  # Dynamic link to your deployment
      api_version = "apps/v1"
    }

    min_replicas = 1
    max_replicas = 5

    metric {
      type = "Resource"

      resource {
        name = "cpu"

        target {
          type                = "Utilization"
          average_utilization = 50  # HPA tries to keep pods ~50% CPU usage
        }
      }
    }
  }
}
