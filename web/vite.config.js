import { fileURLToPath, URL } from 'node:url'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// env must start with this prefix. in .env config
const envPrefix = 'TUI_';

const validateEnv = (env) => {
  const requiredVars = ['TUI_CLI_PORT', 'TUI_BASE_PATH', 'TUI_SERVER_PORT'];
  const missingVars = requiredVars.filter((v) => !(v in env));
  if (missingVars.length > 0) {
    throw new Error(`Missing required environment variables: ${missingVars.join(', ')}`);
  }
};

// https://vitejs.dev/config/
export default ({ mode }) => {
  const env = loadEnv(mode, process.cwd(), envPrefix)
  validateEnv(env);
  process.env = { ...process.env, ...env };

  return defineConfig({
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
        'views': fileURLToPath(new URL('./src/views', import.meta.url)),
        'components': fileURLToPath(new URL('./src/components', import.meta.url)),
      }
    },
    plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver({ importStyle: "sass" })],
      }),
      Components({
        resolvers: [
          ElementPlusResolver({ importStyle: "sass" }),
        ],
      })
    ],
    envPrefix: envPrefix,
    server: {
      open: true,
      port: env.TUI_CLI_PORT,
      proxy: {
        '/api': {
          target: `${env.TUI_BASE_PATH}:${env.TUI_SERVER_PORT}`,
          changeOrigin: true,
          ws: true,
          rewrite: (path) => path.replace(/^\/api/, '/api')
        }
      }
    }
  });
}
