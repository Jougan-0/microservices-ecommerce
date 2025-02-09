"use client";

import { useState, useEffect } from "react";
import axios from "axios";

const cartAPIs = [
  {
    name: "Add to Cart",
    method: "POST",
    path: "/cart/add",
    description: "Adds a product to the user's cart.",
    requestBody: `{
  "product_id": "f895d618-7265-4914-ac41-5abbacbc7d15",
  "quantity": 2
}`,
    exampleResponse: `{
  "message": "Item added to cart"
}`,
  },
  {
    name: "Get Cart",
    method: "GET",
    path: "/cart",
    description: "Fetches the user's cart items.",
    exampleResponse: `[
  {
    "id": 1,
    "product_name": "Laptop",
    "quantity": 2,
    "total_price": 2599.98
  }
]`,
  },
  {
    name: "Remove Cart Item",
    method: "DELETE",
    path: "/cart/remove/:id?quantity=1",
    description:
      "Removes an item from the user's cart. If quantity is less than 1, the whole item is removed.",
    exampleResponse: `{
  "message": "Item removed from cart"
}`,
  },
  {
    name: "Clear Cart",
    method: "DELETE",
    path: "/cart/clear",
    description: "Removes all items from the cart.",
    exampleResponse: `{
  "message": "Cart cleared"
}`,
  },
];

export default function CartDocs() {
  const [selectedAPI, setSelectedAPI] = useState(cartAPIs[0]);
  const [cartId, setCartId] = useState("");
  const [quantity, setQuantity] = useState(1);
  const [requestBody, setRequestBody] = useState(
    selectedAPI.requestBody || "{}"
  );
  const [response, setResponse] = useState(null);

  const [jwtToken, setJwtToken] = useState(() => {
    return typeof window !== "undefined"
      ? localStorage.getItem("jwtToken") || ""
      : "";
  });

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
      let url = `http://localhost:3001${selectedAPI.path}`;

      if (url.includes(":id")) {
        url = url.replace(":id", cartId || "1");
      }

      if (url.includes("?quantity=")) {
        url = url.replace("?quantity=1", `?quantity=${quantity}`);
      }

      const config = {
        method: selectedAPI.method,
        url,
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
        <h2 className="text-lg font-bold">🛒 Cart API</h2>
        <ul className="mt-4">
          {cartAPIs.map((api, index) => (
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

        <label className="block mt-2 font-medium">JWT Token:</label>
        <input
          type="text"
          value={jwtToken}
          onChange={handleJwtChange}
          className="border p-2 w-full mb-4 bg-white text-gray-900"
        />

        {selectedAPI.path.includes(":id") && (
          <>
            <label className="block mt-2 font-medium">Cart ID:</label>
            <input
              type="text"
              value={cartId}
              onChange={(e) => setCartId(e.target.value)}
              className="border p-2 w-full mb-4 bg-white text-gray-900"
            />
          </>
        )}

        {selectedAPI.path.includes("?quantity=") && (
          <>
            <label className="block mt-2 font-medium">Quantity:</label>
            <input
              type="number"
              value={quantity}
              onChange={(e) => setQuantity(e.target.value)}
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
