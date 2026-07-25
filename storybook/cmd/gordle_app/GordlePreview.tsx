import type { PropsWithChildren } from 'react'
import { MemoryRouter } from 'react-router-dom'
import '@/cmd/gordle_app/frontend/src/style.css'

function GordlePreview({ children }: PropsWithChildren) {
  return (
    <MemoryRouter>
      {children}
    </MemoryRouter>
  )
}

export default GordlePreview
