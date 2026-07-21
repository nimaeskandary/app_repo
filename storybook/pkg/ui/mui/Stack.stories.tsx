import type { Meta, StoryObj } from '@storybook/react-vite'
import { Stack } from '@/pkg/ui/ts/mui/stack'
import { Body1 } from '@/pkg/ui/ts/mui/typography'

const meta = {
  title: 'UI/MUI/Stack',
  parameters: {
    layout: 'centered',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <Stack>
      <Body1>First item</Body1>
      <Body1>Second item</Body1>
      <Body1>Third item</Body1>
    </Stack>
  ),
}
