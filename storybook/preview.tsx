import type { Preview } from '@storybook/react-vite'
import CssBaseline from '@mui/material/CssBaseline'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import GordleAppPreview from '@/storybook/cmd/gordle_app/GordlePreview'
import './preview.css'

const lightTheme = createTheme({ palette: { mode: 'light' } })
const darkTheme = createTheme({ palette: { mode: 'dark' } })

const preview: Preview = {
  globalTypes: {
    theme: {
      description: 'Global theme for components',
      toolbar: {
        title: 'Theme',
        icon: 'circlehollow',
        items: ['light', 'dark'],
        dynamicTitle: true,
      },
    },
    previewShell: {
      description: 'Application shell wrapped around the story',
      toolbar: {
        title: 'App shell',
        icon: 'browser',
        items: [
          { value: 'none', title: 'None' },
          { value: 'gordleApp', title: 'Gordle app' },
        ],
        dynamicTitle: true,
      },
    },
  },
  initialGlobals: {
    previewShell: 'none',
    theme: 'dark',
  },
  decorators: [
    (Story, context) => {
      const theme = context.globals.theme === 'light' ? lightTheme : darkTheme
      const story = <Story />

      return (
        <ThemeProvider theme={theme}>
          <CssBaseline />
          {context.globals.previewShell === 'gordleApp'
            ? <GordleAppPreview>{story}</GordleAppPreview>
            : story}
        </ThemeProvider>
      )
    },
  ],
  parameters: {
    viewport: {
      options: {
        iphone17Pro: {
          name: 'iPhone 17 Pro',
          styles: {
            width: '402px',
            height: '874px',
          },
          type: 'mobile',
        },
        desktop: {
          name: 'Desktop',
          styles: {
            width: '1000px',
            height: '618px',
          },
          type: 'desktop',
        },
      },
    },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
}

export default preview
