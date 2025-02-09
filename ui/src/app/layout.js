"use client";

import { usePathname } from "next/navigation";
import "./globals.css";
import JwtProvider from "./JwtContext";

export default function RootLayout({ children }) {
  const pathname = usePathname();
  const showSidebar = pathname !== "/";

  return (
    <JwtProvider>
      <html lang="en">
        <body className="flex flex-col h-screen">
          <div className="flex items-center justify-between bg-gray-500 text-white p-3">
            <a
              href="/"
              className="text-white text-2xl bg-gray-700 px-3 py-2 rounded hover:bg-gray-800"
            >
              🏠
            </a>
            <span className="text-center flex-1 font-bold">
              If you want to generate a JWT token, please{" "}
              <a
                href="/user?login=true"
                className="underline text-yellow-300 hover:text-yellow-500"
              >
                log in with a unique email.
              </a>
            </span>
            <div className="w-10"></div>
          </div>

          <div className="flex flex-1">
            {showSidebar && (
              <aside className="w-1/4 bg-gray-900 text-white p-4">
                <h2 className="text-lg font-bold">📄 API Sections</h2>
                <ul className="mt-4">
                  <li className="p-2 hover:bg-gray-700">
                    <a href="/user" className="block w-full h-full p-2">
                      👤 User APIs
                    </a>
                  </li>
                  <li className="p-2 hover:bg-gray-700">
                    <a href="/catalog" className="block w-full h-full p-2">
                      📦 Product APIs
                    </a>
                  </li>
                  <li className="p-2 hover:bg-gray-700">
                    <a href="/cart" className="block w-full h-full p-2">
                      🛒 Cart APIs
                    </a>
                  </li>
                </ul>
              </aside>
            )}

            <main
              className={`p-6 ${
                showSidebar ? "w-3/4" : "w-full"
              } flex justify-center items-center`}
            >
              {children}
            </main>
          </div>

          <footer className="bg-gray-900 text-white text-center p-3">
            Made with ❤️ by Jougan
          </footer>
        </body>
      </html>
    </JwtProvider>
  );
}
