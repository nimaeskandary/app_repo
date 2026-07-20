import type { Preview } from '@storybook/react-vite'
import './preview.css'

const preview: Preview = {
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
        gordleDesktop: {
          name: 'Gordle Desktop',
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
