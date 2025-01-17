import 'dotenv/config';

/** @type {import('next').NextConfig} */
const nextConfig = {
  trailingSlash: true,
  reactStrictMode: false,
  swcMinify: true,
  output: 'export',
  env: {
    ENVIRONMENT: process.env.ENVIRONMENT,
    API_URL: process.env.API_URL,
  },
  async rewrites() {
    return [
      {
        source: "/api/:path*", // Match API requests
        destination: "http://localhost/:path*", // Forward to backend
      },
    ];
  },
};

export default nextConfig;
