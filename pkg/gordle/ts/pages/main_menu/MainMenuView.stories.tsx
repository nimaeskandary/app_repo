import type { Meta, StoryObj } from '@storybook/react-vite'
import MainMenuView from '@app_repo/pkg_gordle/pages/main_menu/MainMenuView'

const meta = {
  title: 'Gordle/Views/MainMenu',
  component: MainMenuView,
} satisfies Meta<typeof MainMenuView>

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    onPlay: () => undefined,
  },
}
