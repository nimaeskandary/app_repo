import type { Meta, StoryObj } from '@storybook/react-vite'
import GordleView from '@app_repo/pkg_gordle/pages/gordle/GordleView'
import {
  CellState,
  type Cell,
} from '@app_repo/pkg_gordle/domain/Cell'

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
  state: CellState.Unguessed,
}))

// Creates a completed story row with explicit visual states.
function completedRow(word: string, states: CellState[]): Cell[] {
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
        CellState.Correct,
        CellState.RowCorrect,
        CellState.Wrong,
        CellState.RowCorrect,
        CellState.RowCorrect,
      ]),
      completedRow('BRICK', [
        CellState.Wrong,
        CellState.Correct,
        CellState.Wrong,
        CellState.RowCorrect,
        CellState.Wrong,
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
        CellState.Wrong,
        CellState.Wrong,
        CellState.Correct,
        CellState.Wrong,
        CellState.Correct,
      ]),
      completedRow('BRICK', [
        CellState.Wrong,
        CellState.Correct,
        CellState.Wrong,
        CellState.RowCorrect,
        CellState.Wrong,
      ]),
      completedRow('CLOUD', [
        CellState.Correct,
        CellState.Wrong,
        CellState.Wrong,
        CellState.Wrong,
        CellState.Wrong,
      ]),
      completedRow('GHOST', Array.from({ length: 5 }, () => CellState.Wrong)),
      completedRow('PLANT', [
        CellState.Wrong,
        CellState.Wrong,
        CellState.Correct,
        CellState.Correct,
        CellState.Wrong,
      ]),
      completedRow('SHORE', [
        CellState.Wrong,
        CellState.Wrong,
        CellState.Wrong,
        CellState.RowCorrect,
        CellState.Correct,
      ]),
    ],
  },
}

export const Won: Story = {
  args: {
    ...Default.args,
    cells: [
      ...InProgress.args.cells.slice(0, 2),
      completedRow('CRANE', Array.from({ length: 5 }, () => CellState.Correct)),
      blankRow,
      blankRow,
      blankRow,
    ],
    isGameComplete: true,
  },
}
