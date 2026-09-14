# 🔄 DevBoard — GitOps with ArgoCD

This branch contains the **GitOps deployment configuration** for DevBoard.

The desired Kubernetes state is stored in Git, and **ArgoCD continuously synchronizes that state with the Kubernetes cluster**.

## 🔄 GitOps Flow

```text
              Developer
                  │
                  ▼
             Application
                  │
                  ▼
               GitHub
                  │
                  ▼
          Image / Manifest Update
                  │
                  ▼
               ArgoCD
                  │
          ┌───────┴───────┐
          │               │
       Compare          Sync
          │               │
          └───────┬───────┘
                  ▼
              AWS EKS
                  │
                  ▼
            DevBoard Pods
```

## 🏗️ GitOps Architecture

```text
Git Repository
      │
      │ Desired State
      ▼
    ArgoCD
      │
      │ Reconciliation
      ▼
 Kubernetes API
      │
      ▼
   AWS EKS
      │
      ▼
 Application
```

## 📁 Structure

```text
argocd/
│
├── application.yaml
│
└── Kubernetes/
    ├── frontend/
    ├── backend/
    ├── postgres/
    └── gateway/
```

## 🔑 GitOps Principles

* Git is the source of truth
* Kubernetes desired state is version controlled
* ArgoCD continuously reconciles the cluster
* Deployments are auditable through Git history
* Application configuration is managed declaratively

## 🔄 Image Update Flow

```text
New Application Code
        │
        ▼
GitHub Actions
        │
        ▼
Docker Image Build
        │
        ▼
Docker Hub
        │
        ▼
GitOps Image Tag Update
        │
        ▼
Git
        │
        ▼
ArgoCD Detects Change
        │
        ▼
AWS EKS Deployment
```

## 🎯 Objective

Implement a declarative deployment model where Git defines the desired state and ArgoCD continuously reconciles Kubernetes with that state.
