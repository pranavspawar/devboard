# ☸️ DevBoard — Kubernetes

This branch contains the Kubernetes deployment configuration for DevBoard.

## 🏗️ Kubernetes Architecture

```text
                         AWS EKS
                           │
                           ▼
                    Envoy Gateway
                           │
                           ▼
                       HTTPRoute
                           │
                ┌──────────┴──────────┐
                ▼                     ▼
           Frontend Service      Backend Service
                │                     │
                ▼                     ▼
           Frontend Pods         Backend Pods
                                      │
                                      ▼
                                PostgreSQL
                                      │
                                      ▼
                                Persistent
                                  Storage
```

## 📁 Structure

```text
Kubernetes/
│
├── namespace.yaml
│
├── frontend/
│   ├── deployment.yaml
│   └── service.yaml
│
├── backend/
│   ├── deployment.yaml
│   └── service.yaml
│
├── postgres/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── pvc.yaml
│   └── secret.yaml
│
└── gateway/
    ├── gateway.yaml
    └── httproute.yaml
```

> Update the filenames above to exactly match the repository.

## 🚀 Deployment

```bash
kubectl apply -f Kubernetes/
```

Verify:

```bash
kubectl get pods -n devboard
kubectl get svc -n devboard
kubectl get gateway -n devboard
kubectl get httproute -n devboard
```

## 🔍 Main Kubernetes Concepts

* Deployments
* Services
* ConfigMaps
* Secrets
* PersistentVolumeClaims
* Gateway API
* HTTPRoute
* Container health checks
* Resource management
* Kubernetes networking

## 🎯 Objective

Deploy a multi-service application on Kubernetes with service discovery, persistent storage, HTTP routing, and scalable workloads.
