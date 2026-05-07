# GCP Deployment Comparison: GKE vs. Google Cloud Run

This project was developed as part of a university course on DevOps. It demonstrates a technical comparison between two different deployment strategies on Google Cloud Platform (GCP): **Container Orchestration** with Kubernetes and **Serverless** deployment with Cloud Run.

The core objective was to implement a "Build Once, Run Anywhere" workflow using a single Go-based microservice.

## 🚀 Project Overview

The application is a lightweight Go web server that identifies its deployment environment dynamically. This allows us to use the exact same container image across different infrastructure targets while maintaining environment-specific behavior.

### Key Features
* **Unified Application Base**: One shared Go application and Dockerfile used across both deployment targets (GKE and Cloud Run).
* **Platform Awareness**: Uses environment variables to detect if it is running on GKE or Cloud Run.
* **Automated CI/CD**: Fully configured pipelines using Google Cloud Build.
* **Optimized Images**: Multi-stage Docker builds based on `debian:bullseye-slim` for security and performance.

## 🔄 CI/CD Pipeline (Google Cloud Build)

A core part of this DevOps project is the automated CI/CD workflow. We use Google Cloud Build to automate the path from code to production:

**Build**: The pipeline triggers on every push, building the Docker image from the multi-stage Dockerfile.

**Store**: Images are pushed to the Google Artifact Registry.

**Deploy**:

`cloudbuild-gke.yaml` handles the rolling update to the GKE cluster.

`cloudbuild-run.yaml` manages the revision deployment to Cloud Run.

## 🛠 Tech Stack
* **Language:** Go 1.22
* **Containerization:** Docker
* **Orchestration:** Google Kubernetes Engine (GKE)
* **Serverless:** Google Cloud Run
* **CI/CD:** Google Cloud Build

## 📂 Repository Structure

```text
├── main.go                 # Go application logic
├── go.mod                  # Dependencies and module definition
├── Dockerfile              # Multi-stage build configuration
├── cloudbuild-gke.yaml     # CI/CD pipeline for Kubernetes deployment
├── cloudbuild-run.yaml     # CI/CD pipeline for Cloud Run deployment
├── kubernetes/             # Infrastructure-as-Code manifests
│   ├── deployment.yaml     # K8s Deployment (Replicas, Env-Vars)
│   └── service.yaml        # K8s LoadBalancer Service
└── .gitignore              # Standard exclusions (OS, IDE files)
```

## 💻 Local Development

To run the application locally for testing purposes:

1. Clone the repository:

```bash
git clone https://github.com/Fuerueze/my-devops-app-GKE.git
```

2. Navigate to the project folder:

```bash
cd my-devops-app-GKE
```

3. Run the server:
   
```bash
go run main.go
```

Note: By default, the application will listen on port 8080 and display "Unknown" as the platform unless the `PLATFORM` environment variable is set locally.

**Authors**: Fürüze Saritoprak & Nhi Nguyen

Developed as a collaborative project for the DevOps module.
