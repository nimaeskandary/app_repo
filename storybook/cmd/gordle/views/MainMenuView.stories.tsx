import type { Meta, StoryObj } from '@storybook/react-vite'
import MainMenuView from '@app_repo/pkg_gordle_frontend/views/main_menu/MainMenuView'
import GordlePreview from '@/storybook/cmd/gordle/GordlePreview'

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
} satisfies Meta<typeof MainMenuView>

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    onPlay: () => undefined,
  },
}
