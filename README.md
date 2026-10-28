# 🚀 DevBoard

**Production-style application built to demonstrate an end-to-end DevOps and DevSecOps workflow.**

[![GitHub Actions](https://github.com/pranavspawar/devboard/actions/workflows/Devsecops.yml/badge.svg)](https://github.com/pranavspawar/devboard/actions/workflows/Devsecops.yml)
![Docker](https://img.shields.io/badge/Docker-Containerization-2496ED?logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-Orchestration-326CE5?logo=kubernetes&logoColor=white)
![AWS EKS](https://img.shields.io/badge/AWS-EKS-FF9900?logo=amazoneks&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/CI-GitHub_Actions-2088FF?logo=githubactions&logoColor=white)
![Argo CD](https://img.shields.io/badge/GitOps-Argo_CD-EF7B4D?logo=argo&logoColor=white)
![Terraform](https://img.shields.io/badge/IaC-Terraform-844FBA?logo=terraform&logoColor=white)
![Trivy](https://img.shields.io/badge/Security-Trivy-1904DA?logo=aquasecurity&logoColor=white)
![Prometheus](https://img.shields.io/badge/Metrics-Prometheus-E6522C?logo=prometheus&logoColor=white)
![Grafana](https://img.shields.io/badge/Monitoring-Grafana-F46800?logo=grafana&logoColor=white)
![Loki](https://img.shields.io/badge/Logging-Loki-F4A51C?logo=grafana&logoColor=white)
![OpenTelemetry](https://img.shields.io/badge/Observability-OpenTelemetry-7B61FF?logo=opentelemetry&logoColor=white)

---

## 🏗️ Architecture

```text
👨‍💻 Developer
      │
      ▼
   🐙 GitHub
      │
      ▼
 🔄 GitHub Actions
      │
      ├── 🔐 Security & Tests
      │
      ▼
  🐳 Docker Hub
      │
      ▼
  🚀 Argo CD
      │
      ▼
   ☁️ AWS EKS
      │
      ├── 🌐 Envoy Gateway
      │       │
      │       ├── 💻 Frontend
      │       └── ⚙️ Backend ──► 🐘 PostgreSQL
      │
      └── 📊 Observability
              │
              ▼
        🔭 OpenTelemetry
           │       │
           ▼       ▼
      📈 Prometheus  📝 Loki
           │       │
           └───┬───┘
               ▼
           📊 Grafana
```

---

## 🛠️ Tech Stack

| 🔧 Area | 🚀 Technologies |
|---|---|
| ☁️ Cloud | AWS EKS |
| ☸️ Kubernetes | Kind, Kubernetes, AWS EKS |
| 🐳 Containers | Docker |
| 🔄 CI | GitHub Actions |
| 🚀 GitOps / CD | Argo CD |
| 🏗️ Infrastructure as Code | Terraform |
| 🌐 Gateway | Envoy Gateway |
| 📊 Observability | OpenTelemetry, Prometheus, Grafana, Loki |
| 🐘 Database | PostgreSQL |
| 🔐 Security | Trivy, Gitleaks |
| 💻 Frontend | React, Vite |
| ⚙️ Backend | Go, Gin |

---

## 📁 Project Structure

```text
DevBoard/
├── .github/workflows/   # CI/CD and DevSecOps workflows
├── frontend/            # React/Vite frontend
├── backend/             # Go/Gin API
├── Kubernetes/          # Kubernetes manifests
├── argocd/              # GitOps configuration
├── terraform/           # AWS infrastructure
├── monitoring/          # Prometheus/Grafana/Loki
└── docker-compose.yml   # Local development environment
```

---

## 🔄 CI/CD and DevSecOps Pipeline

```text
💻 Code Push / Pull Request
          │
          ▼
   🔄 GitHub Actions
          │
          ├── 🧪 Code Validation & Tests
          ├── 🔐 Gitleaks Secret Scanning
          ├── 🔍 Trivy Container Scanning
          └── 🛡️ SonarQube Analysis
                    │
                    ▼
             🐳 Docker Build
                    │
                    ▼
               📦 Docker Hub
                    │
                    ▼
           🔁 GitOps Image Update
                    │
                    ▼
               🚀 Argo CD
                    │
                    ▼
                 ☁️ AWS EKS
```

The workflow is intended to automate code validation, security checks, container image creation, and the image-update process used by the GitOps deployment workflow.

---

## ☸️ Kubernetes Deployment

### 💻 Local Development

```text
🐳 Docker → 🟢 Kind → ☸️ Kubernetes
```

Kind provides a local Kubernetes environment for developing and testing application deployments.

### ☁️ Cloud Deployment

```text
🐳 Docker → 📦 Docker Hub → 🚀 Argo CD → ☁️ AWS EKS
```

AWS EKS provides the managed Kubernetes environment for cloud deployment. Argo CD manages deployment reconciliation based on the configured Git repository and Kubernetes manifests.

---

## 📊 Observability

```text
💻 Application
      │
      ▼
🔭 OpenTelemetry
      │
      ├── 📈 Metrics ──► Prometheus ──► Grafana
      │
      └── 📝 Logs ─────► Loki ────────► Grafana
```

- **OpenTelemetry:** Instrumentation and telemetry collection.
- **Prometheus:** Metrics collection and monitoring.
- **Loki:** Centralized log aggregation.
- **Grafana:** Metrics and log visualization.

---

## 🔐 Security and DevSecOps

- **Gitleaks:** Detects potentially exposed secrets in source code.
- **Trivy:** Scans container images and supported artifacts for security vulnerabilities.
- **SonarQube:** Performs static code analysis and quality checks.
- **GitHub Actions:** Automates the configured validation and security workflow stages.

---

## 👨‍💻 DevOps Highlights

- 🐳 Containerized frontend and backend services.
- ☸️ Developed and tested Kubernetes deployments using Kind.
- ☁️ Configured cloud deployment using AWS EKS.
- 🔄 Automated CI workflows using GitHub Actions.
- 🔐 Integrated security scanning with Trivy and Gitleaks.
- 🚀 Used Argo CD for GitOps-based deployment.
- 🏗️ Managed AWS infrastructure using Terraform.
- 🌐 Configured Envoy Gateway for HTTP routing.
- 🔭 Integrated OpenTelemetry into the observability architecture.
- 📈 Used Prometheus and Grafana for metrics and visualization.
- 📝 Used Loki for centralized logging.
- 🐘 Integrated PostgreSQL for persistent application data.

---

## 👨‍💻 Author

**Pranav Pawar**

[![GitHub](https://img.shields.io/badge/GitHub-pranavspawar-181717?logo=github&logoColor=white)](https://github.com/pranavspawar)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-0A66C2?logo=linkedin&logoColor=white)](https://linkedin.com/in/devopsengineerpranav)

**DevOps Engineer | AWS | Kubernetes | CI/CD | DevSecOps | GitOps**
