import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import CssBaseline from '@mui/material/CssBaseline'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import type { PropsWithChildren } from 'react'

const theme = createTheme({
  palette: {
    mode: 'dark',
    background: { default: '#06070f' },
    text: { primary: '#fff' },
  },
})

// Applies the app's global styles and parent layout.
function Layout({ children }: PropsWithChildren) {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline enableColorScheme />
      <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100dvh' }}>
        {children}
      </div>
    </ThemeProvider>
  )
}

export default Layout
