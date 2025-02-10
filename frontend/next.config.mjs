import 'dotenv/config';

/** @type {import('next').NextConfig} */
const nextConfig = {
  trailingSlash: true,
  reactStrictMode: false,
  swcMinify: true,
  env: {
    ENVIRONMENT: process.env.ENVIRONMENT,
    API_URL: process.env.API_URL,
  },
  async rewrites() {
    return [
      {
        source: "/api/:path*", // Match API requests
        destination: "http://habitsappbackend/:path*", // Forward to backend
      },
    ];
  },
};

export default nextConfig;
