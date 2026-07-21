import type { Meta, StoryObj } from '@storybook/react-vite'
import { Body1 } from '@/pkg/ui/ts/mui/typography'
import { Grid } from '@/pkg/ui/ts/mui/grid'

const meta = {
  title: 'UI/MUI/Grid',
  parameters: {
    layout: 'centered',
  },
} satisfies Meta

export default meta

type Story = StoryObj<typeof meta>

export const Default: Story = {
  render: () => (
    <div className="flex flex-col gap-4">
      <Grid container columns={2}>
        <Grid size={1}>
            <Body1>1</Body1>
        </Grid>
        <Grid size={1}>
            <Body1>2</Body1>
        </Grid>
        <Grid size={1}>
            <Body1>3</Body1>
        </Grid>
        <Grid size={1}>
            <Body1>4</Body1>
        </Grid>
        <Grid size={1}>
            <Body1>5</Body1>
        </Grid>
        <Grid size={1}>
            <Body1>6</Body1>
        </Grid>
      </Grid>
    </div>
  )
}
