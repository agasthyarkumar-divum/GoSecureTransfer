import { defineConfig } from 'vite'
import react, { reactCompilerPreset } from '@vitejs/plugin-react'
import babel from '@rolldown/plugin-babel'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    babel({ presets: [reactCompilerPreset()] })
  ],
  server: {
    proxy: {
      '/upload': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/login': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/register': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/download': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/files': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/delete': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/delete-all': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    }
  }
})
