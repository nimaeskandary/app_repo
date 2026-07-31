import type { PropsWithChildren } from 'react'
import { MemoryRouter } from 'react-router-dom'

function GordlePreview({ children }: PropsWithChildren) {
  return (
    <MemoryRouter>
      <div style={{ background: '#06070f', minHeight: '100vh' }}>
        {children}
      </div>
    </MemoryRouter>
  )
}

export default GordlePreview
