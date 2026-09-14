Hi, I'm Pranav Pawar 👋
DevOps Engineer | AWS | Kubernetes | CI/CD 
  
# 🚀 DevBoard

**Production-style microservices application built to demonstrate an end-to-end DevOps workflow.**

`AWS EKS` · `Kubernetes` · `Docker` · `GitHub Actions` · `ArgoCD` · `Terraform` · `Prometheus` · `Grafana` · `Loki`

## 🏗️ Architecture

```text
Developer
   │
   ▼
 GitHub
   │
   ▼
GitHub Actions
   │
   ├── Security & Tests
   │
   ▼
Docker Hub
   │
   ▼
 ArgoCD
   │
   ▼
 AWS EKS
   │
   ├── Envoy Gateway
   │       │
   │       ├── Frontend
   │       └── Backend ──► PostgreSQL
   │
   └── Prometheus ──► Grafana
       Loki ────────► Grafana
```

## 🛠️ Tech Stack

| Area       | Technologies        |
| ---------- | ------------------- |
| Cloud      | AWS EKS             |
| Containers | Docker, Kubernetes  |
| CI/CD      | GitHub Actions      |
| GitOps     | ArgoCD              |
| IaC        | Terraform           |
| Gateway    | Envoy Gateway       |
| Monitoring | Prometheus, Grafana |
| Logging    | Loki                |
| Database   | PostgreSQL          |
| Security   | Trivy, Gitleaks     |

## 📁 Project Structure

```text
DevBoard/
├── .github/workflows/   # CI/CD & DevSecOps
├── frontend/            # React/Vite
├── backend/             # Go/Gin API
├── Kubernetes/          # K8s manifests
├── argocd/              # GitOps configuration
├── terraform/           # AWS infrastructure
├── monitoring/          # Prometheus/Grafana/Loki
└── docker-compose.yml   # Local environment
```

## 🔄 Deployment Flow

```text
Code Push
   ↓
GitHub Actions
   ↓
Docker Build & Security Scan
   ↓
Docker Hub
   ↓
GitOps Image Update
   ↓
ArgoCD Sync
   ↓
AWS EKS
```

## 👨‍💻 DevOps Highlights

* Containerized frontend and backend services
* Automated CI/CD with security scanning
* GitOps deployment using ArgoCD
* Kubernetes deployment on AWS EKS
* Terraform-based infrastructure
* Envoy Gateway for HTTP routing
* Prometheus/Grafana for monitoring
* Loki for centralized logging
* PostgreSQL persistent storage

**Pranav Pawar**
DevOps Engineer | AWS | Kubernetes | CI/CD
