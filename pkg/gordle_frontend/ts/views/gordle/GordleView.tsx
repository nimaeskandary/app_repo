import { Container } from '@app_repo/pkg_ui/mui/container'
import Board from '@app_repo/pkg_gordle_frontend/components/Board'
import type { GordleCell } from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewProps = {
  cells: GordleCell[][]
}

function GordleView({ cells }: GordleViewProps) {
  return (
    <Container>
      <Board cells={cells} />
    </Container>
  )
}

export default GordleView
