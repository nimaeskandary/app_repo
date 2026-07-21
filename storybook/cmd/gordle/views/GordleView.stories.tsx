import type { Meta, StoryObj } from '@storybook/react-vite'
import GordleView, { GordleCellState } from '../../../../cmd/gordle_app/frontend/src/pages/gordle/GordleView'
import GordlePreview from '../GordlePreview'

const meta = {
  title: 'Cmd/Gordle/Views/GordleView',
  component: GordleView,
  decorators: [
    (Story) => (
      <GordlePreview>
        <Story />
      </GordlePreview>
    ),
  ],
  parameters: {
    layout: 'centered',
  },
  argTypes: {
    cells: { control: 'object' },
  },
} satisfies Meta<typeof GordleView>

export default meta

type Story = StoryObj<typeof meta>

const unguessedRow = Array.from({ length: 5 }, () => ({
  letter: '',
  state: GordleCellState.Unguessed,
}))

export const Default: Story = {
  args: {
    cells: [
      [
        { letter: 'C', state: GordleCellState.Wrong },
        { letter: 'R', state: GordleCellState.Correct },
        { letter: 'A', state: GordleCellState.RowCorrect },
        { letter: 'N', state: GordleCellState.Wrong },
        { letter: 'E', state: GordleCellState.Correct },
      ],
      unguessedRow,
      unguessedRow,
      unguessedRow,
      unguessedRow,
      unguessedRow,
    ],
  },
}
