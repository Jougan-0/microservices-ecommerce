export default function Home() {
  return (
    <div className="flex flex-col justify-center items-center h-full">
      <div className="text-center">
        <h1 className="text-3xl font-bold mb-4">📄 API Documentation</h1>
        <p className="mb-6">
          Select an API section to view endpoints and test them.
        </p>

        <div className="flex gap-4 justify-center">
          <a
            href="/user"
            className="bg-green-500 text-white p-4 rounded-lg text-lg block w-40 text-center"
          >
            👤 User APIs
          </a>
          <a
            href="/catalog"
            className="bg-yellow-500 text-white p-4 rounded-lg text-lg block w-40 text-center"
          >
            📦 Product APIs
          </a>
          <a
            href="/cart"
            className="bg-blue-500 text-white p-4 rounded-lg text-lg block w-40 text-center"
          >
            🛒 Cart APIs
          </a>
        </div>
      </div>
    </div>
  );
}
