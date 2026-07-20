import type { Meta, StoryObj } from '@storybook/react-vite'
import { PrimaryButton } from '@/pkg/ui/ts/button'
import { fn } from 'storybook/test'

const meta = {
  title: 'UI/PrimaryButton',
  component: PrimaryButton,
  parameters: {
    layout: 'centered',
  },
} satisfies Meta<typeof PrimaryButton>

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    title: 'Primary Button',
    onClick: fn(),
  },
}
