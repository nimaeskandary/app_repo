import { Container } from '@app_repo/pkg_ui/mui/container'
import { Stack } from '@app_repo/pkg_ui/mui/stack'
import Board from '@app_repo/pkg_gordle_frontend/components/Board'
import Keyboard from '@app_repo/pkg_gordle_frontend/components/Keyboard'
import type { GordleCell } from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewProps = {
  cells: GordleCell[][]
  onLetterCommit: (letter: string) => void
  rubricWord: string
}

function GordleView({ cells, onLetterCommit, rubricWord }: GordleViewProps) {
  return (
    <Container>
      <Stack spacing={2}>
        <Board cells={cells} rubricWord={rubricWord} />
        <Keyboard onLetterCommit={onLetterCommit} />
      </Stack>
    </Container>
  )
}

export default GordleView
