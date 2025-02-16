# Kubernetes Deployment for Microservices Ecommerce

This folder contains Kubernetes manifests to deploy the **Microservices Ecommerce project** locally or on a Kubernetes cluster.

---

## **Folder Structure**

```
k8s/
  ├── namespace.yaml         # Namespace for the project (optional)
  ├── configmap.yaml         # ConfigMap for non-sensitive environment variables
  ├── secret.yaml            # Secret for sensitive data like DB credentials
  ├── db-pvc.yaml            # PersistentVolumeClaim for PostgreSQL database storage
  ├── db-deployment.yaml     # Deployment for PostgreSQL database
  ├── db-service.yaml        # Service for PostgreSQL database
  ├── backend-deployment.yaml # Deployment for backend microservices
  ├── backend-service.yaml    # Service for backend
  ├── frontend-deployment.yaml # Deployment for frontend service
  ├── frontend-service.yaml    # Service for frontend
  └── ingress.yaml           # Ingress configuration for external access (optional)
```

---

## **How to Deploy Kubernetes Resources Locally**

### **Prerequisites**

- **Docker Desktop with Kubernetes enabled** or **Minikube** installed.
- **kubectl** installed and configured.
- **Make** installed for running `make` commands.

---

### **Steps to Deploy**

1. **Ensure Kubernetes is running locally**:
   - Start Docker Desktop with Kubernetes enabled or start Minikube:
     ```bash
     minikube start
     ```

2. **Run the `make k8s` Command**:
   ```bash
   make k8s
   ```
   This command will apply all the manifests from the `k8s/` folder:
   - Create the namespace (if provided).
   - Create ConfigMaps and Secrets.
   - Deploy PostgreSQL with persistent storage.
   - Deploy the backend and frontend services.
   - Optionally, configure an Ingress for external access.

---

### **Verify the Deployment**

Once the resources are deployed, verify them using:

```bash
kubectl get pods -n commerce
kubectl get services -n commerce
```

---

### **Accessing the Application**

- **Expose the Backend**:
  - Use `kubectl port-forward` to access the backend API locally.
  ```bash
  kubectl port-forward svc/backend 3001:80 -n commerce
  ```
  Access the backend at `http://localhost:3001`.


- **Without Ingress**:
  - Use `kubectl port-forward` to access services locally.
  ```bash
  kubectl port-forward svc/frontend 3000:80 -n commerce
  ```
  Access the frontend at `http://localhost:3000`.

- **With Ingress**:
  - Update your `/etc/hosts` to map your domain to `127.0.0.1` (for local testing).
    ```bash
    127.0.0.1 myapp.example.com
    ```
  - Access the app at `http://myapp.example.com`.

---

### **Deleting the Kubernetes Resources**

To delete all the resources deployed from the `k8s/` folder:
```bash
kubectl delete -f k8s/
```

---

### **Notes**

- **Secrets**: Make sure to update `secret.yaml` with your own database credentials and JWT secret before deployment.
- **Ingress**: If you're not using Ingress, you can omit the `ingress.yaml` file during deployment.

---

### **Makefile Command Reference**

- **Deploy**:
  ```bash
  make k8s
  ```
- **Delete Deployment**:
  ```bash
  make k8s-delete
  ```

---

This `README.md` provides detailed instructions for deploying the project using Kubernetes manifests from the `k8s/` folder and using `make k8s` for local development. 😊

