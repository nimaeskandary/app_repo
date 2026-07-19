import type { Preview } from '@storybook/react-vite'
import '../cmd/wordle/frontend/public/style.css'
import './preview.css'

type ViewportMode = 'web' | 'mobile'

const preview: Preview = {
  globalTypes: {
    viewportMode: {
      description: 'Preview wrapper',
      toolbar: {
        title: 'Viewport',
        icon: 'browser',
        items: [
          { value: 'web', title: 'Web' },
          { value: 'mobile', title: 'Mobile' },
        ],
        dynamicTitle: true,
      },
    },
  },
  initialGlobals: {
    viewportMode: 'web',
  },
  decorators: [
    (Story, context) => {
      const viewportMode = context.globals.viewportMode as ViewportMode

      return (
        <div className="storybook-viewport-root" data-viewport-mode={viewportMode}>
          <div className="storybook-viewport-frame">
            <div className="app-shell">
              <div className="bg" aria-hidden="true" />
              <Story />
            </div>
          </div>
        </div>
      )
    },
  ],
  parameters: {
    layout: 'fullscreen',
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
}

export default preview
