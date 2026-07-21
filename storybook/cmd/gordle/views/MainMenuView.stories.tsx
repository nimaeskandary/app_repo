import type { Meta, StoryObj } from '@storybook/react-vite'
import GordlePreview from '@/storybook/cmd/gordle/GordlePreview'
import MainMenuView from '@/cmd/gordle_app/frontend/src/pages/main_menu/MainMenuView'

const meta = {
  title: 'Cmd/Gordle/Views/MainMenu',
  component: MainMenuView,
  decorators: [
    (Story) => (
      <GordlePreview>
        <Story />
      </GordlePreview>
    ),
  ],
  argTypes: {
    cells: { control: 'object' },
  },
} satisfies Meta<typeof MainMenuView>

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
  },
}
