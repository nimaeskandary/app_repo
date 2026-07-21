import type { Meta, StoryObj } from '@storybook/react-vite'
import { InfoButton, PrimaryButton, ErrorButton, SecondaryButton, SuccessButton } from '@/pkg/ui/ts/button'

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
      <PrimaryButton title='primary' onClick={() => {}} />
      <SecondaryButton title='secondary' onClick={() => {}} />
      <ErrorButton title='error' onClick={() => {}} />
      <InfoButton title='info' onClick={() => {}} />
      <SuccessButton title='success' onClick={() => {}} />
    </div>
  )
}
