import type { PropsWithChildren } from 'react'
import { MemoryRouter } from 'react-router-dom'
import '../../../cmd/gordle_app/frontend/src/gordle.css'

function GordlePreview({ children }: PropsWithChildren) {
  return (
    <MemoryRouter>
      <div className="gordle min-h-dvh">
        <div className="bg" aria-hidden="true" />
        <div className="flex min-h-dvh flex-col">
          {children}
        </div>
      </div>
    </MemoryRouter>
  )
}

export default GordlePreview
