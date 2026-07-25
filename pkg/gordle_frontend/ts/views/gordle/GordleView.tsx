import { Box } from '@app_repo/pkg_ui/mui/box'
import { Container } from '@app_repo/pkg_ui/mui/container'
import { Grid } from '@app_repo/pkg_ui/mui/grid'
import { H3 } from '@app_repo/pkg_ui/mui/typography'
import {
  GordleCellState,
  type GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewProps = {
  cells: GordleCell[][]
}

const cellStyles = {
  [GordleCellState.Unguessed]: { borderColor: '#3a3a3c' },
  [GordleCellState.Wrong]: { backgroundColor: '#3a3a3c', borderColor: '#3a3a3c' },
  [GordleCellState.RowCorrect]: { backgroundColor: '#b59f3b', borderColor: '#b59f3b' },
  [GordleCellState.Correct]: { backgroundColor: '#538d4e', borderColor: '#538d4e' },
}

function GordleView({ cells }: GordleViewProps) {
  return (
    <Container>
        <Grid container columns={5} spacing={1}>
          {cells.flatMap((row, rowIndex) =>
            row.map((cell, columnIndex) => (
              <Grid key={`${rowIndex}-${columnIndex}`} size={1}>
                <Box
                  sx={{
                    ...cellStyles[cell.state],
                    aspectRatio: '1',
                    border: 2,
                    borderRadius: 2,
                    color: '#ffffff',
                    display: 'grid',
                    placeItems: 'center',
                    // for the H4 cqi
                    containerType: 'inline-size',
                  }}>
                  <H3 sx={{
                    // this constrains the font to 70% of the parent, so the H4 doesn't cause the box to grow.
                    // the H3 is just here really for weight
                    fontSize: '70cqi'}}>{cell.letter}</H3>
                </Box>
              </Grid>
            )),
          )}
        </Grid>
    </Container>
  )
}

export default GordleView
