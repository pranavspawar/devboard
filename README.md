# 🚀 DevBoard

**Production-style microservices application built to demonstrate an end-to-end DevOps workflow.**

☁️ `AWS EKS` · ☸️ `Kind` · ⚙️ `Kubernetes` · 🐳 `Docker` · 🔄 `GitHub Actions` · 🚀 `ArgoCD` · 🏗️ `Terraform` · 📈 `Prometheus` · 📊 `Grafana` · 📝 `Loki` · 🔭 `OpenTelemetry`

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
  🚀 ArgoCD
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

| 🔧 Area          | 🚀 Technologies                          |
| ---------------- | ---------------------------------------- |
| ☁️ Cloud         | AWS EKS                                  |
| ☸️ Kubernetes    | Kind, AWS EKS                            |
| 🐳 Containers    | Docker                                   |
| 🔄 CI/CD         | GitHub Actions                           |
| 🚀 GitOps        | ArgoCD                                   |
| 🏗️ IaC          | Terraform                                |
| 🌐 Gateway       | Envoy Gateway                            |
| 📊 Observability | OpenTelemetry, Prometheus, Grafana, Loki |
| 🐘 Database      | PostgreSQL                               |
| 🔐 Security      | Trivy, Gitleaks                          |

---

## 📁 Project Structure

```text
DevBoard/
├── 🔄 .github/workflows/   # CI/CD & DevSecOps
├── 💻 frontend/             # React/Vite
├── ⚙️ backend/              # Go/Gin API
├── ☸️ Kubernetes/           # Kubernetes manifests
├── 🚀 argocd/               # GitOps configuration
├── 🏗️ terraform/            # AWS infrastructure
├── 📊 monitoring/           # Prometheus/Grafana/Loki
└── 🐳 docker-compose.yml    # Local environment
```

---

## 🔄 Deployment Flow

```text
💻 Code Push
     ↓
🔄 GitHub Actions
     ↓
🐳 Docker Build & 🔐 Security Scan
     ↓
📦 Docker Hub
     ↓
🔁 GitOps Image Update
     ↓
🚀 ArgoCD Sync
     ↓
☁️ AWS EKS
```

---

## ☸️ Kubernetes

### 💻 Local Development

```text
🐳 Docker → 🟢 Kind → ☸️ Kubernetes
```

Used **Kind** to develop and test Kubernetes deployments locally.

### ☁️ Cloud Deployment

```text
🐳 Docker → 📦 Docker Hub → 🚀 ArgoCD → ☁️ AWS EKS
```

Used **AWS EKS** for cloud-based Kubernetes deployment.

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

**OpenTelemetry** is used as the **observability instrumentation and collection layer**, with **Prometheus, Loki, and Grafana** used for monitoring, logging, and visualization.

---

## 👨‍💻 DevOps Highlights

* 🐳 Containerized frontend and backend services
* ☸️ Local Kubernetes development using **Kind**
* ☁️ Cloud Kubernetes deployment using **AWS EKS**
* 🔄 Automated CI/CD with **GitHub Actions**
* 🔐 Security scanning with **Trivy & Gitleaks**
* 🚀 GitOps deployment using **ArgoCD**
* 🏗️ Terraform-based AWS infrastructure
* 🌐 Envoy Gateway for HTTP routing
* 🔭 OpenTelemetry-based observability
* 📈 Prometheus & Grafana monitoring
* 📝 Loki centralized logging
* 🐘 PostgreSQL persistent storage

---

## 👨‍💻 Author

**Pranav Pawar**

🚀 **DevOps Engineer** | ☁️ AWS | ☸️ Kubernetes | 🔄 CI/CD
