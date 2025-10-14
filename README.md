# Kubernetes Go Server Project

This project contains a complete Kubernetes setup for deploying a Go web server application with various Kubernetes resources and configurations.

## Project Structure

```
├── Dockerfile                     # Docker image for Go application
├── server.go                      # Go web server source code
└── k8s/                          # Kubernetes manifests
    ├── cluster-issuer.yaml       # Cert-manager cluster issuer for SSL
    ├── configmap-env.yaml        # Environment variables ConfigMap
    ├── configmap-family.yaml     # Family data ConfigMap
    ├── deployment.yaml           # Main application deployment
    ├── hpa.yaml                  # Horizontal Pod Autoscaler
    ├── ingress.yaml              # Ingress configuration
    ├── kind.yaml                 # Kind cluster configuration
    ├── metrics-server.yaml       # Metrics server for HPA
    ├── mysql-service-h.yaml      # Headless service for MySQL
    ├── pod.yaml                  # Standalone pod definition
    ├── pv.yaml_                  # Persistent Volume (needs fixing)
    ├── pvc.yaml                  # Persistent Volume Claim
    ├── replicaset.yaml           # ReplicaSet configuration
    ├── secret.yaml               # Secret with credentials
    ├── service.yaml              # Service for Go application
    ├── statefulset.yaml          # MySQL StatefulSet
    └── namespaces/               # Namespace-specific resources
        ├── deployment.yaml       # Deployment with RBAC
        └── security.yaml         # ServiceAccount and RBAC
```

## Application Overview

The Go server provides the following endpoints:

- `/` - Returns a greeting with name and age from environment variables
- `/healthz` - Health check endpoint (returns 500 for first 10 seconds, then 200)
- `/secret` - Displays user credentials from Kubernetes secrets
- `/configmap` - Shows family data from mounted ConfigMap file

## Prerequisites

- Docker
- Kubernetes cluster (local or cloud)
- kubectl configured
- Kind (for local development)

## Quick Start

### 1. Build Docker Image

```bash
docker build -t luizcurti/hello-go:latest .
```

### 2. Create Kind Cluster (Local Development)

```bash
kind create cluster --config k8s/kind.yaml
```

### 3. Deploy to Kubernetes

Apply the Kubernetes manifests in the following order:

```bash
# Create ConfigMaps and Secrets
kubectl apply -f k8s/configmap-env.yaml
kubectl apply -f k8s/configmap-family.yaml
kubectl apply -f k8s/secret.yaml

# Create Storage
kubectl apply -f k8s/pv.yaml
kubectl apply -f k8s/pvc.yaml

# Deploy Application
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Optional: Setup Autoscaling
kubectl apply -f k8s/metrics-server.yaml
kubectl apply -f k8s/hpa.yaml

# Optional: Setup Ingress
kubectl apply -f k8s/cluster-issuer.yaml
kubectl apply -f k8s/ingress.yaml
```

### 4. Access the Application

```bash
# Port forward to access locally
kubectl port-forward service/goserver-service 8080:80

# Access the application
curl http://localhost:8080
```

## Configuration Details

### Environment Variables (ConfigMap)
- `NAME`: Person's name (default: "Wesley")
- `AGE`: Person's age (default: "36")

### Secrets
- `USER`: Base64 encoded username
- `PASSWORD`: Base64 encoded password

### Storage
- Persistent Volume: 5Gi storage for application data
- ConfigMap Volume: Family data mounted at `/go/myfamily/family.txt`

### Resource Limits
- CPU Request: 0.05 cores
- CPU Limit: 0.05 cores  
- Memory Request: 20Mi
- Memory Limit: 25Mi

### Health Checks
- **Startup Probe**: Checks `/healthz` every 4s, allows 30 failures
- **Readiness Probe**: Checks `/healthz` every 3s, 1 failure threshold
- **Liveness Probe**: Checks `/healthz` every 5s, 1 failure threshold

### Auto Scaling
- **HPA**: Scales between 1-30 replicas based on 25% CPU utilization
- Requires metrics-server to be installed

## MySQL StatefulSet

The project includes a MySQL StatefulSet with:
- 4 replicas
- Headless service (`mysql-h`)
- Persistent storage (5Gi per instance)
- Root password: "root"

## Ingress Configuration

- Host: `ingress.fullcycle.com.br`
- Path: `/admin`
- SSL/TLS with Let's Encrypt
- Nginx ingress class

## RBAC Configuration

Located in `k8s/namespaces/`:
- ServiceAccount: `server`
- ClusterRole: Read access to pods, services, and deployments
- ClusterRoleBinding: Binds role to service account

## Known Issues

⚠️ **Important**: Some files need corrections before deployment:

1. **Dockerfile**: Line 3 has syntax error - needs to be split
2. **pv.yaml_**: Wrong file extension and invalid YAML comments
3. **ingress.yaml**: Uses deprecated API version `networking.k8s.io/v1beta1`
4. **cluster-issuer.yaml**: Uses deprecated API version `cert-manager.io/v1alpha2`

## Monitoring

The project includes metrics-server configuration for monitoring CPU/memory usage and enabling HPA functionality.

## Development

### Local Testing

```bash
# Build and run locally
go run server.go

# Test endpoints
curl http://localhost:8000
curl http://localhost:8000/healthz
```

### Docker Testing

```bash
docker run -p 8000:8000 \
  -e NAME="Test User" \
  -e AGE="25" \
  luizcurti/hello-go:latest
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is for educational purposes and demonstrates Kubernetes deployment patterns.