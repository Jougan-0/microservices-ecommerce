# Services (Golang Backend) for Microservices Ecommerce

This document provides details on the **Golang backend services** of the Microservices Ecommerce project, including the project structure, API routes, environment variables, database entity relationship diagram, and steps to run the services locally.

---

## **Database Entity Relationship Diagram**

Below is the entity relationship diagram based on the provided models:

```text
+------------+       +------------+        +-------------+
|   User     |       |   Cart      |        |   Product   |
+------------+       +------------+        +-------------+
| ID (UUID)  |<----->| UserID      |        | ID (UUID)   |
| Name       |       | ProductID   |<------>| Name        |
| Email      |       | Quantity    |        | Description |
| Password   |       | CreatedAt   |        | Price       |
| CreatedAt  |       +------------+         | Stock       |
+------------+                              | CreatedAt   |
                                            +-------------+
```

- **User**: Stores user information.
- **Product**: Stores product details.
- **Cart**: Represents a user's cart, linking users to products with quantities.

---

## **Project Structure**

```
services/
  ├── cmd/                  # Main entry point for the application
  ├── db/                   # Database initialization and connection
  ├── handlers/             # API request handlers
  ├── middleware/           # Middleware for authentication and logging
  ├── models/               # Data models
  ├── repository/           # Database interaction layer
  ├── routes/               # API route definitions
  ├── services/             # Business logic layer
  ├── utils/                # Utility functions (e.g., JWT, logging)
  ├── go.mod/go.sum         # Go dependencies
  └── Dockerfile            # Docker setup for the backend services
```

---

## **API Routes**
- **User Service** (Mixed Authentication):
  - `POST /register` – Register a new user (Unauthenticated).
  - `POST /login` – Authenticate a user and generate JWT (Unauthenticated).
  - `GET /user/profile` – Get user profile (Authenticated).
  - `PUT /user/update-profile` – Update user profile (Authenticated).
  - `DELETE /user/delete-account` – Delete user account (Authenticated).
- **Catalog Service** (Mixed Authentication):
  - `POST /api/product` – Create a product (Authenticated).
  - `PUT /api/product/:id` – Update a product (Authenticated).
  - `DELETE /api/product/:id` – Delete a product (Authenticated).
  - `GET /products` – Fetch all products (Unauthenticated).
  - `GET /product/:id` – Fetch product details (Unauthenticated).
- **Cart Service** (Authenticated):
  - `POST /cart/add` – Add product to cart.
  - `GET /cart` – Retrieve user's cart.
  - `DELETE /cart/remove/:id` – Remove product from cart.
  - `DELETE /cart/clear` – Clear all items from the cart.
- **User Service**:
  - `POST /register` – Register a new user.
  - `POST /login` – Authenticate a user and generate JWT.
- **Catalog Service**:
  - `GET /products` – Fetch all products.
  - `GET /products/:id` – Fetch product details.
- **Cart Service**:
  - `POST /cart/add` – Add product to cart.
  - `GET /cart` – Retrieve user's cart.
  - `DELETE /cart/remove/:id` – Remove product from cart.
  - `DELETE /cart/clear` – Clear all items from the cart.

---

## **Environment Variables**
- `DB_HOST` – Database host.
- `DB_PORT` – Database port.
- `DB_USER` – Database user.
- `DB_PASSWORD` – Database password.
- `DB_NAME` – Database name.
- `JWT_SECRET` – Secret key for JWT authentication.

---

## **How to Run the Backend Services Locally**

### **Prerequisites**
- **Go** installed (version 1.21 or later).
- **Docker** installed and running.

### **Steps to Run**

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/microservices-ecommerce.git
   cd microservices-ecommerce/services
   ```

2. **Set environment variables**:
   - Create a `.env` file in the `services/` directory.
   - Add necessary environment variables (DB credentials, JWT secret).

3. **Run the services locally**:
   ```bash
   go run cmd/main.go
   ```
   The backend API will be available at `http://localhost:3001`.

---

## **Running with Docker**

You can also build and run the backend services using Docker:
```bash
docker build -t services-backend .
docker run -p 3001:3001 services-backend
```

---

## **Notes**
- Ensure the database is running and accessible before starting the backend services.
- JWT tokens are generated during authentication and should be stored securely on the client side for API access.

This `README.md` provides a clear guide to understanding, configuring, and running the Golang backend services for the Microservices Ecommerce project.

