"use client";

import { useState, useEffect } from "react";
import axios from "axios";

const catalogAPIs = [
  {
    name: "Create Product",
    method: "POST",
    path: "/api/product",
    description: "Adds a new product to the catalog.",
    requestBody: `{
  "name": "Laptop",
  "description": "Powerful gaming laptop",
  "price": 1299.99,
  "stock": 20
}`,
    exampleResponse: `{
  "message": "Product added successfully"
}`,
  },
  {
    name: "Get All Products",
    method: "GET",
    path: "/products",
    description: "Fetches all available products.",
    exampleResponse: `[
  {
    "id": 1,
    "name": "Laptop",
    "description": "Powerful gaming laptop",
    "price": 1299.99,
    "stock": 20,
    "created_at": "2025-02-09T12:30:00Z"
  }
]`,
  },
  {
    name: "Get Product by ID",
    method: "GET",
    path: "/product/:id",
    description: "Fetches a specific product by ID.",
    exampleResponse: `{
  "id": 1,
  "name": "Laptop",
  "description": "Powerful gaming laptop",
  "price": 1299.99,
  "stock": 20,
  "created_at": "2025-02-09T12:30:00Z"
}`,
  },
  {
    name: "Update Product",
    method: "PUT",
    path: "/api/product/:id",
    description: "Updates an existing product by ID.",
    requestBody: `{
  "name": "Updated Laptop",
  "description": "Latest gaming laptop",
  "price": 1399.99,
  "stock": 15
}`,
    exampleResponse: `{
  "message": "Product updated successfully"
}`,
  },
  {
    name: "Delete Product",
    method: "DELETE",
    path: "/api/product/:id",
    description: "Deletes a product from the catalog.",
    exampleResponse: `{
  "message": "Product deleted successfully"
}`,
  },
];

export default function CatalogDocs() {
  const [selectedAPI, setSelectedAPI] = useState(catalogAPIs[0]);
  const [productId, setProductId] = useState("");

  const [jwtToken, setJwtToken] = useState(() => {
    return typeof window !== "undefined"
      ? localStorage.getItem("jwtToken") || ""
      : "";
  });
  const [backendUrl, setBackendUrl] = useState('');

  useEffect(() => {
    fetch('/api/config')
      .then(res => res.json())
      .then(data => setBackendUrl(data.backendUrl));
  }, []);
  const [requestBody, setRequestBody] = useState(
    selectedAPI.requestBody || "{}"
  );
  const [response, setResponse] = useState(null);

  useEffect(() => {
    const storedToken = localStorage.getItem("jwtToken");
    if (storedToken) {
      setJwtToken(storedToken);
    }
  }, []);

  const handleJwtChange = (e) => {
    const newToken = e.target.value;
    setJwtToken(newToken);
    localStorage.setItem("jwtToken", newToken);
  };

  const handleRequest = async () => {
    try {
      const baseURL = backendUrl;
      const config = {
        method: selectedAPI.method,
        url: `${baseURL}${selectedAPI.path.replace(
          ":id",
          productId || "1"
        )}`,
        headers: {
          "Content-Type": "application/json",
          Authorization: jwtToken ? `Bearer ${jwtToken}` : "",
        },
        data: selectedAPI.method !== "GET" ? JSON.parse(requestBody) : null,
      };

      const res = await axios(config);
      setResponse(res.data);
    } catch (error) {
      setResponse(error.response ? error.response.data : error.message);
    }
  };

  return (
    <div className="flex h-screen">
      <aside className="w-1/4 bg-gray-900 text-white p-4">
        <h2 className="text-lg font-bold">📦 Catalog API</h2>
        <ul className="mt-4">
          {catalogAPIs.map((api, index) => (
            <li
              key={index}
              className={`p-2 cursor-pointer hover:bg-gray-700 ${
                selectedAPI.name === api.name ? "bg-blue-500" : ""
              }`}
              onClick={() => {
                setSelectedAPI(api);
                setRequestBody(api.requestBody || "{}");
                setResponse(null);
              }}
            >
              {api.name}
            </li>
          ))}
        </ul>
      </aside>

      <section className="w-2/4 p-6 bg-gray-100 text-gray-900">
        <h2 className="text-2xl font-bold">{selectedAPI.name}</h2>
        <p className="mt-2">{selectedAPI.description}</p>

        <h3 className="mt-4 font-bold">🔹 Endpoint</h3>
        <code className="block bg-white text-black p-2 rounded border border-gray-300">
          {selectedAPI.method} {selectedAPI.path}
        </code>

        {selectedAPI.requestBody && (
          <>
            <h3 className="mt-4 font-bold">🔹 Example Request Body</h3>
            <pre className="bg-white text-black p-3 rounded border border-gray-300">
              {selectedAPI.requestBody}
            </pre>
          </>
        )}

        <h3 className="mt-4 font-bold">🔹 Example Response</h3>
        <pre className="bg-white text-black p-3 rounded border border-gray-300">
          {selectedAPI.exampleResponse}
        </pre>
      </section>

      <aside className="w-1/4 p-6 bg-white text-gray-900 border-l border-gray-300">
        <h3 className="text-lg font-bold">🛠 Test API</h3>

        <label className="block mt-2 font-medium">JWT Token (needed):</label>
        <input
          type="text"
          value={jwtToken}
          onChange={handleJwtChange}
          className="border p-2 w-full mb-4 bg-white text-gray-900"
        />

        {selectedAPI.path.includes(":id") && (
          <>
            <label className="block mt-2 font-medium">Product ID:</label>
            <input
              type="text"
              value={productId}
              onChange={(e) => setProductId(e.target.value)}
              className="border p-2 w-full mb-4 bg-white text-gray-900"
            />
          </>
        )}

        {selectedAPI.requestBody && (
          <>
            <label className="block font-medium">Request Body:</label>
            <textarea
              value={requestBody}
              onChange={(e) => setRequestBody(e.target.value)}
              className="border p-2 w-full mb-4 bg-white text-gray-900"
              rows={8}
            />
          </>
        )}

        <button
          onClick={handleRequest}
          className="bg-blue-500 text-white p-2 rounded w-full hover:bg-blue-600"
        >
          Send {selectedAPI.method} Request
        </button>

        {response && (
          <div className="mt-4">
            <h2 className="text-lg font-bold">Response:</h2>
            <pre className="bg-white text-black p-3 rounded border border-gray-300 overflow-x-auto max-h-64 whitespace-pre-wrap break-words">
              {JSON.stringify(response, null, 2)}
            </pre>
          </div>
        )}
      </aside>
    </div>
  );
}
