import type { Meta, StoryObj } from '@storybook/react-vite'
import { PrimaryButton } from './button'
import { Container } from './container'
import { H1 } from './typography'

const meta = {
  title: 'UI/MUI/Container',
  parameters: {
    layout: 'centered',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <div className="flex flex-col gap-4">
      <Container>
        <H1>This is a container</H1>
        <br />
        <PrimaryButton onClick={() => {}}>woot</PrimaryButton>
      </Container>
    </div>
  )
}
