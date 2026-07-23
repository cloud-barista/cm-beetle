import type { NextConfig } from 'next';
import path from 'path';

const nextConfig: NextConfig = {
  // Produce a self-contained build for small Docker images
  output: 'standalone',
  // Explicitly define root directory for Turbopack to prevent multiple lockfiles warning
  turbopack: {
    root: path.resolve(__dirname),
  },
};

export default nextConfig;
