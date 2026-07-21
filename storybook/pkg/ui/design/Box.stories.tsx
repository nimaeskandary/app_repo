import type { Meta, StoryObj } from '@storybook/react-vite'
import { Box } from '@/pkg/ui/ts/design/box'
import { Body1 } from '@/pkg/ui/ts/design/typography'
import { PrimaryButton } from '@/pkg/ui/ts/design/button'

const meta = {
  title: 'UI/Design/Box',
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
      <PrimaryButton title='woot' onClick={() => {}} />
    </Box>
  ),
}
