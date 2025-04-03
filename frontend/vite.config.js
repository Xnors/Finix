import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'

export default defineConfig({
  plugins: [solid()],
  server: {
    proxy: {
      '/get_info': {
        target: 'http://localhost:8080', // 或者是 https://localhost:8080
        changeOrigin: true,
      },
    },
  },
})
