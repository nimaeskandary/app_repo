import type { PropsWithChildren } from 'react'
import { MemoryRouter } from 'react-router-dom'
import Layout from '@/cmd/gordle_app/frontend/src/Layout'

function GordlePreview({ children }: PropsWithChildren) {
  return (
    <MemoryRouter>
      <Layout>{children}</Layout>
    </MemoryRouter>
  )
}

export default GordlePreview
