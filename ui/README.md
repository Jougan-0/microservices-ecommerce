
# UI for Microservices Ecommerce

This document provides details on the **UI component** of the Microservices Ecommerce project, including how JWT tokens are managed and steps to run the UI locally.

---

## **JWT Token Management**

- Once a user logs in, a **JWT token** is generated and stored securely in the browser's `localStorage`.
- This token is automatically retrieved and added to API requests where authentication is required, ensuring seamless communication with backend services.

---

## **How to Run the UI Locally**

### **Prerequisites**
- **Node.js** installed (version 18 or later).
- **Docker and Kubernetes** if you are running the backend via Kubernetes.

### **Steps to Run the UI**

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/microservices-ecommerce.git
   cd microservices-ecommerce/ui
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Start the UI locally**:
   ```bash
   npm start
   ```
   The UI will be available at `http://localhost:3000`.

---

## **Accessing the UI After Kubernetes Deployment**
- If you have deployed the backend using Kubernetes, use the following command to expose the backend API locally:
  ```bash
  kubectl port-forward svc/backend 3001:80 -n commerce
  ```
- Then, start the UI as mentioned above.

---

## **Notes**
- Ensure the backend is running (either locally or on Kubernetes) before starting the UI.
- The JWT token once generated is automatically stored and used in subsequent API calls.

This `README.md` ensures that developers can easily run the UI, understand how tokens are managed, and interact with the backend securely.

