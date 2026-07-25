import type { Meta, StoryObj } from '@storybook/react-vite'
import { Box } from '@app_repo/pkg_ui/mui/box'
import { PrimaryButton } from '@app_repo/pkg_ui/mui/button'
import { Body1 } from '@app_repo/pkg_ui/mui/typography'

const meta = {
  title: 'UI/MUI/Box',
  parameters: {
    layout: 'centered',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <Box>
      <Body1>This is a box</Body1>
      <PrimaryButton onClick={() => {}}>woot</PrimaryButton>
    </Box>
  ),
}
