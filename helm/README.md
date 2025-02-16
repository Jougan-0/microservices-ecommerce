# Microservices Ecommerce Helm Chart

This Helm chart is designed to deploy the Microservices Ecommerce project on a Kubernetes cluster using Kubernetes manifests for backend, frontend, and database services.

## Project Structure

```
helm/
  charts/
    templates/
      backend-deployment.yaml
      backend-service.yaml
      configmap.yaml
      db-deployment.yaml
      db-pvc.yaml
      db-service.yaml
      frontend-deployment.yaml
      frontend-service.yaml
      ingress.yaml
      namespace.yaml
      secret.yaml
    Chart.yaml
    values.yaml.gotmpl
```

### Key Files

- **Chart.yaml**: Helm chart metadata including name, description, and version.
- **values.yaml.gotmpl**: Template for Helm values, used during releases to set image tags and secrets dynamically. `values.yaml` is `.gitignore`d to avoid committing sensitive data.
- **templates/**: Kubernetes resource templates such as Deployments, Services, ConfigMaps, Secrets, Ingress, and PVCs.

---

## Deployment Steps

### Prerequisites
- Kubernetes cluster
- Helm installed

### Steps

1. **Navigate to the Helm charts directory:**
   ```bash
   cd helm/charts
   ```

2. **Create a `values.yaml` from `values.yaml.gotmpl` and set values as needed:**
   ```bash
   cp values.yaml.gotmpl values.yaml
   ```
   - Set `BACKEND_TAG`, `FRONTEND_TAG`, `INGRESS_HOST`, `DB_USER`, `DB_PASSWORD`, and `JWT_SECRET`.

3. **Deploy using Helm with custom values:**
   ```bash
   helm install commerce . -f values.yaml
   ```

4. **Verify deployment:**
   ```bash
   kubectl get all -n commerce
   ```

5. **Access the application via the configured ingress host.**

### Automated Releases
- On every new release, a new Helm chart is deployed with the latest Docker images and configurations.
- Ensure that `values.yaml.gotmpl` is updated with the correct tags and secrets before creating a new release.

---

## Helm Chart Explanation

- **Backend Deployment and Service**:
  - Uses the backend Docker image, sets environment variables from ConfigMap and secrets.
- **Frontend Deployment and Service**:
  - Deploys the frontend UI container and exposes it via a ClusterIP service.
- **Database Deployment and PVC**:
  - Deploys PostgreSQL with persistent storage.
- **Ingress**:
  - Configures NGINX ingress for routing traffic to the frontend service.
- **Namespace and Secrets**:
  - Creates the `commerce` namespace and stores sensitive data securely.

---

This chart simplifies the Kubernetes deployment of the microservices ecommerce project using Helm. Adjust values and apply to your cluster seamlessly!

