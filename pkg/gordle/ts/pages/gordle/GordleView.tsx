import { Container } from '@app_repo/pkg_ui/mui/container'
import { Stack } from '@app_repo/pkg_ui/mui/stack'
import { H3 } from '@app_repo/pkg_ui/mui/typography'
import Board from '@app_repo/pkg_gordle/pages/gordle/components/Board'
import Keyboard from '@app_repo/pkg_gordle/pages/gordle/components/Keyboard'
import type { Cell } from '@app_repo/pkg_gordle/domain/Cell'

export type GordleViewProps = {
  cells: Cell[][]
  isGameComplete: boolean
  onWordCommit: (word: string) => boolean
}

function GordleView({
  cells,
  isGameComplete,
  onWordCommit,
}: GordleViewProps) {
  return (
    <Container>
      <Stack spacing={2}>
        <Board cells={cells} />
        {isGameComplete
          ? <H3 data-testid="game-won" sx={{ textAlign: 'center', color: 'yellow' }}>
              YOU WON!
            </H3>
          : <Keyboard onWordCommit={onWordCommit} />}
      </Stack>
    </Container>
  )
}

export default GordleView
