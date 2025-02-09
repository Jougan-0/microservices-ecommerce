"use client";

import { useState, useEffect, useContext, Suspense } from "react";
import axios from "axios";
import { useSearchParams } from "next/navigation";
import { JwtContext } from "../JwtContext";

export default function UserDocsWrapper() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <UserDocs />
    </Suspense>
  );
}

function UserDocs() {
  const { jwtToken, updateJwtToken } = useContext(JwtContext);
  const searchParams = useSearchParams();
  const loginSelected = searchParams.get("login") === "true";

  const userAPIs = [
    {
      name: "Register",
      method: "POST",
      path: "/register",
      description: "Registers a new user.",
      requestBody: `{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}`,
      exampleResponse: `{
  "message": "User registered successfully"
}`,
      requiresAuth: false,
    },
    {
      name: "Login",
      method: "POST",
      path: "/login",
      description: "Authenticates a user and returns a JWT token.",
      requestBody: `{
  "email": "john@example.com",
  "password": "securepassword"
}`,
      exampleResponse: `{
  "token": "eyJhbGciOiJIUzI1NiIsInR5c..."
}`,
      requiresAuth: false,
    },
    {
      name: "Get User Profile",
      method: "GET",
      path: "/user/profile",
      description: "Fetches the logged-in user's profile.",
      exampleResponse: `{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2025-02-09T12:30:00Z"
}`,
      requiresAuth: true,
    },
    {
      name: "Update Profile",
      method: "PUT",
      path: "/user/update-profile",
      description: "Updates the logged-in user's profile.",
      requestBody: `{
  "name": "John Updated"
}`,
      exampleResponse: `{
  "message": "Profile updated successfully"
}`,
      requiresAuth: true,
    },
    {
      name: "Delete Account",
      method: "DELETE",
      path: "/user/delete-account",
      description: "Deletes the user's account permanently.",
      exampleResponse: `{
  "message": "Account deleted successfully"
}`,
      requiresAuth: true,
    },
  ];

  const [selectedAPI, setSelectedAPI] = useState(
    loginSelected ? userAPIs[1] : userAPIs[0]
  );
  const [requestBody, setRequestBody] = useState(
    selectedAPI.requestBody || "{}"
  );
  const [response, setResponse] = useState(null);

  useEffect(() => {
    if (loginSelected) {
      setSelectedAPI(userAPIs[1]);
      setRequestBody(userAPIs[1].requestBody || "{}");
    }
  }, [loginSelected]);

  const handleRequest = async () => {
    try {
      const config = {
        method: selectedAPI.method,
        url: `http://localhost:3001${selectedAPI.path}`,
        headers: {
          "Content-Type": "application/json",
          ...(selectedAPI.requiresAuth && jwtToken
            ? { Authorization: `Bearer ${jwtToken}` }
            : {}),
        },
        data: selectedAPI.method !== "GET" ? JSON.parse(requestBody) : null,
      };

      const res = await axios(config);
      setResponse(res.data);

      if (res.data.token) {
        updateJwtToken(res.data.token);
      }
    } catch (error) {
      setResponse(error.response ? error.response.data : error.message);
    }
  };

  return (
    <div className="flex h-screen">
      <aside className="w-1/4 bg-gray-900 text-white p-4">
        <h2 className="text-lg font-bold">👤 User API</h2>
        <ul className="mt-4">
          {userAPIs.map((api, index) => (
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

        {selectedAPI.requiresAuth && (
          <>
            <label className="block mt-2 font-medium">JWT Token:</label>
            <input
              type="text"
              value={jwtToken}
              onChange={(e) => updateJwtToken(e.target.value)}
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
