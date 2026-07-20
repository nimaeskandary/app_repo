import type { Meta, StoryObj } from '@storybook/react-vite'
import { InfoButton, PrimaryButton, ErrorButton, SecondaryButton, SuccessButton } from '@/pkg/ui/ts/button'
import { fn } from 'storybook/test'

const meta = {
  title: 'UI/Button',
  parameters: {
    layout: 'centered',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <div className="flex flex-col gap-4">
      <PrimaryButton title='primary' onClick={fn()} />
      <SecondaryButton title='secondary' onClick={fn()} />
      <ErrorButton title='error' onClick={fn()} />
      <InfoButton title='info' onClick={fn()} />
      <SuccessButton title='success' onClick={fn()} />
    </div>
  )
}
