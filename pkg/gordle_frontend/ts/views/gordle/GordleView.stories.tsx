import type { Meta, StoryObj } from '@storybook/react-vite'
import GordleView from '@app_repo/pkg_gordle_frontend/views/gordle/GordleView'
import {
  GordleCellState,
  type GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

const meta = {
  title: 'Gordle/Views/Gordle',
  component: GordleView,
  argTypes: {
    cells: { control: 'object' },
    onWordCommit: { action: 'word committed' },
  },
} satisfies Meta<typeof GordleView>

export default meta

type Story = StoryObj<typeof meta>

const blankRow = Array.from({ length: 5 }, () => ({
  letter: '',
  state: GordleCellState.Unguessed,
}))

// Creates a completed story row with explicit visual states.
function completedRow(word: string, states: GordleCellState[]): GordleCell[] {
  return word.split('').map((letter, index) => ({ letter, state: states[index] }))
}

export const Default: Story = {
  args: {
    cells: [
      blankRow,
      blankRow,
      blankRow,
      blankRow,
      blankRow,
      blankRow,
    ],
    isGameComplete: false,
    onWordCommit: () => true,
  },
}

export const InProgress: Story = {
  args: {
    ...Default.args,
    cells: [
      completedRow('CATER', [
        GordleCellState.Correct,
        GordleCellState.RowCorrect,
        GordleCellState.Wrong,
        GordleCellState.RowCorrect,
        GordleCellState.RowCorrect,
      ]),
      completedRow('BRICK', [
        GordleCellState.Wrong,
        GordleCellState.Correct,
        GordleCellState.Wrong,
        GordleCellState.RowCorrect,
        GordleCellState.Wrong,
      ]),
      blankRow,
      blankRow,
      blankRow,
      blankRow,
    ],
  },
}

export const FullBoard: Story = {
  args: {
    ...Default.args,
    cells: [
      completedRow('SLATE', [
        GordleCellState.Wrong,
        GordleCellState.Wrong,
        GordleCellState.Correct,
        GordleCellState.Wrong,
        GordleCellState.Correct,
      ]),
      completedRow('BRICK', [
        GordleCellState.Wrong,
        GordleCellState.Correct,
        GordleCellState.Wrong,
        GordleCellState.RowCorrect,
        GordleCellState.Wrong,
      ]),
      completedRow('CLOUD', [
        GordleCellState.Correct,
        GordleCellState.Wrong,
        GordleCellState.Wrong,
        GordleCellState.Wrong,
        GordleCellState.Wrong,
      ]),
      completedRow('GHOST', Array.from({ length: 5 }, () => GordleCellState.Wrong)),
      completedRow('PLANT', [
        GordleCellState.Wrong,
        GordleCellState.Wrong,
        GordleCellState.Correct,
        GordleCellState.Correct,
        GordleCellState.Wrong,
      ]),
      completedRow('SHORE', [
        GordleCellState.Wrong,
        GordleCellState.Wrong,
        GordleCellState.Wrong,
        GordleCellState.RowCorrect,
        GordleCellState.Correct,
      ]),
    ],
  },
}

export const Won: Story = {
  args: {
    ...Default.args,
    cells: [
      ...InProgress.args.cells.slice(0, 2),
      completedRow('CRANE', Array.from({ length: 5 }, () => GordleCellState.Correct)),
      blankRow,
      blankRow,
      blankRow,
    ],
    isGameComplete: true,
  },
}
