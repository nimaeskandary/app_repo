import type { Meta, StoryObj } from '@storybook/react-vite'
import { ErrorButton, InfoButton, PrimaryButton, SecondaryButton, SuccessButton } from './button'

const meta = {
  title: 'UI/MUI/Button',
  parameters: {
    layout: 'centered',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <div className="flex flex-col gap-4">
      <PrimaryButton onClick={() => {}}>primary</PrimaryButton>
      <SecondaryButton onClick={() => {}}>secondary</SecondaryButton>
      <ErrorButton onClick={() => {}}>error</ErrorButton>
      <InfoButton onClick={() => {}}>info</InfoButton>
      <SuccessButton onClick={() => {}}>success</SuccessButton>
    </div>
  )
}
