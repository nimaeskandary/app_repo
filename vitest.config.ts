import path from 'node:path'
import { fileURLToPath } from 'node:url'
import { defineConfig } from 'vitest/config'

const projectRoot = path.dirname(fileURLToPath(import.meta.url))

export default defineConfig({
  resolve: {
    alias: {
      '@': projectRoot,
    },
  },
  test: {
    environment: 'jsdom',
    include: ['pkg/**/*.test.{ts,tsx}'],
    setupFiles: ['./pkg/test_utils/ts/vitest.setup.ts'],
  },
})
