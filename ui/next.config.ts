import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  // Produce a self-contained build for small Docker images
  output: 'standalone',
};

export default nextConfig;
