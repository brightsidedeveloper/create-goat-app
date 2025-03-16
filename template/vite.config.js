import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  server: {
    port: 3000,
    allowedHosts: true,
  },
  plugins: [tailwindcss()],
})
