# Local Kubernetes Setup Guide

This guide describes how to set up a local Kubernetes environment for testing the Cloud-Barista CM-Beetle deployment.

## Recommended Tools

### 1. Kind (Kubernetes in Docker) - **Recommended for beginners**

Kind is very fast and runs a Kubernetes cluster using Docker containers as "nodes".

**Installation:**

```bash
# For Linux
[ $(uname -m) = x86_64 ] && curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

**Create Cluster:**

```bash
kind create cluster --name cloud-barista
```

---

### 2. k3d (k3s in Docker) - **Best for Ingress testing**

k3d is a lightweight wrapper for k3s that makes it easy to manage multi-node clusters in Docker.

**Installation:**

```bash
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

**Create Cluster with Port Mapping:**

```bash
k3d cluster create cloud-barista -p "8080:80@loadbalancer" -p "8443:443@loadbalancer"
```

---

### 3. Minikube - **Highly versatile**

Minikube is the most mature tool and supports various drivers (Docker, VirtualBox, etc.).

**Installation:**

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

**Start Cluster:**

```bash
minikube start --driver=docker
minikube addons enable ingress
```

## WSL2 Specific Considerations

If you are running on **Windows with WSL2**, pay attention to the following for a better experience:

### 1. Resource Management (`.wslconfig`)

Kubernetes can be hungry for RAM. Limit WSL2 memory usage to prevent Windows from slowing down. Create or edit `%USERPROFILE%\.wslconfig` in Windows:

```ini
[wsl2]
memory=8GB   # Adjust based on your total RAM
cpus=4
```

### 2. Filesystem Performance

**CRITICAL**: Keep your project files and Kubernetes manifests inside the Linux filesystem (e.g., `/home/ubuntu/...`). Accessing files under `/mnt/c/` is significantly slower and can cause timeout issues with etcd or other databases.

### 3. Networking & Service Access

- **Localhost**: WSL2 usually maps `localhost` to Windows, so `kubectl port-forward` should work directly from Windows browsers.
- **Service Type**: For local testing on WSL2, `NodePort` or `kubectl port-forward` is easier to manage than `LoadBalancer` unless you use a tool like `metallb`.

### 4. Docker Backend

If using **Docker Desktop**, ensure "Use the WSL 2 based engine" is enabled in Settings. If using **Docker-CE** directly in WSL2, ensure your user is in the `docker` group.

---

## Performance Tip: 2-Core Systems (WSL2)

Your environment has **16GB RAM (Excellent)** but **2 CPU cores (Minimum)**. Running a full Kubernetes cluster (Control Plane + App Pods) on 2 cores can be tight.

### Recommended Choice: k3d / k3s

For 2-core systems, **k3d (k3s in Docker)** is significantly lighter than KIND. It shares a single process for the control plane and is optimized for low-resource environments.

### KIND Optimization (If you use KIND)

If you prefer KIND, limit the resource overhead:

- Use a **single-node** cluster (default).
- Disable non-essential add-ons (like heavy dashboards).
- Be patient during `kubectl apply` as the control plane handles API requests.

---

## Useful Tools

- **kubectl**: The standard CLI for Kubernetes.
- **k9s**: A terminal-based UI for managing Kubernetes clusters.
- **Lens**: A feature-rich IDE for Kubernetes.
