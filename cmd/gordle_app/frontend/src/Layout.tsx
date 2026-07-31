import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import type { PropsWithChildren } from 'react'
import './style.css'

// Applies the app's global styles and parent layout.
function Layout({ children }: PropsWithChildren) {
  return <div className="flex min-h-dvh flex-col">{children}</div>
}

export default Layout
