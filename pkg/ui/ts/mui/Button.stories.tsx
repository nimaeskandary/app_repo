import type { Meta, StoryObj } from '@storybook/react-vite'
import { ErrorButton, InfoButton, PrimaryButton, SecondaryButton, SuccessButton } from './button'
import { Stack } from './stack'

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
    <Stack spacing={2}>
      <PrimaryButton onClick={() => {}}>primary</PrimaryButton>
      <SecondaryButton onClick={() => {}}>secondary</SecondaryButton>
      <ErrorButton onClick={() => {}}>error</ErrorButton>
      <InfoButton onClick={() => {}}>info</InfoButton>
      <SuccessButton onClick={() => {}}>success</SuccessButton>
    </Stack>
  )
}
