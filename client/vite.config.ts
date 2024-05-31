import { defineConfig, loadEnv  } from 'vite'
import vue from '@vitejs/plugin-vue'


export default defineConfig( ({ mode }) =>  {
  process.env = {...process.env, ...loadEnv(mode, process.cwd())};
  const service_host = `${process.env.SERVICE_HOST ?? 'http://localhost:3000'}`;
  return {
    server: {
      port: 8080,
      proxy: {
        "/api": {
          target: service_host+"/api",
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        }
      }
    },
    plugins: [vue()],
  }
})
