"use client";

import { createContext, useState, useEffect } from "react";

export const JwtContext = createContext();

export default function JwtProvider({ children }) {
  const [jwtToken, setJwtToken] = useState(() => {
    return typeof window !== "undefined"
      ? localStorage.getItem("jwtToken") || ""
      : "";
  });

  const updateJwtToken = (token) => {
    setJwtToken(token);
    if (token) {
      localStorage.setItem("jwtToken", token);
    } else {
      localStorage.removeItem("jwtToken");
    }
  };

  useEffect(() => {
    const storedToken = localStorage.getItem("jwtToken");
    if (storedToken) {
      setJwtToken(storedToken);
    }
  }, []);

  return (
    <JwtContext.Provider value={{ jwtToken, updateJwtToken }}>
      {children}
    </JwtContext.Provider>
  );
}
