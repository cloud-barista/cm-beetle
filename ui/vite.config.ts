import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    tailwindcss()
  ],
  server: {
    port: 5173,
    proxy: {
      '/beetle': {
        target: 'http://localhost:8056',
        changeOrigin: true,
        secure: false,
      },
      '/honeybee': {
        target: 'http://localhost:8081',
        changeOrigin: true,
        secure: false,
      },
      '/damselfly': {
        target: 'http://localhost:8082',
        changeOrigin: true,
        secure: false,
      },
      '/tumblebug': {
        target: 'http://localhost:8056',
        changeOrigin: true,
        secure: false,
      }
    }
  }
})
