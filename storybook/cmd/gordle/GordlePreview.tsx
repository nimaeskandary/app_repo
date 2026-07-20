import type { PropsWithChildren } from 'react'
import '../../../cmd/gordle/frontend/src/gordle.css'

function GordlePreview({ children }: PropsWithChildren) {
  return (
    <div className="gordle min-h-dvh">
      <div className="bg" aria-hidden="true" />
      <div className="app-shell flex min-h-dvh flex-col">
        {children}
      </div>
    </div>
  )
}

export default GordlePreview
