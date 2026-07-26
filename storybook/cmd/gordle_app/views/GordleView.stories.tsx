import type { Meta, StoryObj } from '@storybook/react-vite'
import GordleView from '@app_repo/pkg_gordle_frontend/views/gordle/GordleView'
import { GordleCellState } from '@app_repo/pkg_gordle_frontend/model/GordleCell'
import GordlePreview from '@/storybook/cmd/gordle_app/GordlePreview'


const meta = {
  title: 'Cmd/Gordle/Views/Gordle',
  component: GordleView,
  decorators: [
    (Story) => (
      <GordlePreview>
        <Story />
      </GordlePreview>
    ),
  ],
  argTypes: {
    cells: { control: 'object' },
    onWordCommit: { action: 'word committed' },
    rubricWord: { control: 'text' },
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
    onWordCommit: () => true,
    rubricWord: 'CRANE',
  },
}
