import type { Preview } from '@storybook/react-vite'
import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import GordleAppPreview from '@/storybook/cmd/gordle_app/GordlePreview'
import './preview.css'

const preview: Preview = {
  globalTypes: {
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
    previewShell: 'gordleApp',
    viewport: { value: 'iphone17Pro', isRotated: false },
  },
  decorators: [
    (Story, context) => {
      const story = <Story />

      return context.globals.previewShell === 'gordleApp'
        ? <GordleAppPreview>{story}</GordleAppPreview>
        : story
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
