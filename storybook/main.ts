import type { StorybookConfig } from '@storybook/react-vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

const config: StorybookConfig = {
  stories: [
    './**/*.stories.@(ts|tsx)',
    '../pkg/**/*.stories.@(ts|tsx)',
  ],
  staticDirs: ['../cmd/gordle_app/frontend/public'],
  framework: {
    name: '@storybook/react-vite',
    options: {},
  },
  viteFinal(config) {
    return {
      ...config,
      plugins: [...(config.plugins ?? []), react(), tailwindcss()],
      optimizeDeps: {
        ...config.optimizeDeps,
        include: [
          ...(config.optimizeDeps?.include ?? []),
          'react',
          'react-dom',
          'react/jsx-runtime',
        ],
      },
      resolve: {
        ...config.resolve,
        alias: {
          ...config.resolve?.alias,
          '@': path.resolve(__dirname, '..'),
        },
      },
    }
  },
}

export default config
