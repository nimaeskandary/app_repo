import type { Meta, StoryObj } from '@storybook/react-vite'
import { Container } from '@/pkg/ui/ts/mui/container'
import { H1 } from '@/pkg/ui/ts/mui/typography'
import { PrimaryButton } from '@/pkg/ui/ts/mui/button'

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
