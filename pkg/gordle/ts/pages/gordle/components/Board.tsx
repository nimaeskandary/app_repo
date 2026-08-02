import { Box } from '@app_repo/pkg_ui/mui/box'
import { Grid } from '@app_repo/pkg_ui/mui/grid'
import { Fade } from '@app_repo/pkg_ui/mui/transitions'
import { H3 } from '@app_repo/pkg_ui/mui/typography'
import {
  CellState,
  type Cell,
} from '@app_repo/pkg_gordle/domain/Cell'

export type BoardProps = {
  cells: Cell[][]
}

const cellStyles = {
  [CellState.Unguessed]: { borderColor: '#3a3a3c' },
  [CellState.Wrong]: { backgroundColor: '#3a3a3c', borderColor: '#3a3a3c' },
  [CellState.RowCorrect]: { backgroundColor: '#b59f3b', borderColor: '#b59f3b' },
  [CellState.Correct]: { backgroundColor: '#538d4e', borderColor: '#538d4e' },
}

// Renders the Gordle cell grid.
function Board({ cells }: BoardProps) {
  return (
    <Grid container columns={5} spacing={1}>
      {cells.flatMap((row, rowIndex) =>
        row.map((cell, columnIndex) => (
          <Grid key={`${rowIndex}-${columnIndex}`} size={1}>
            <Fade key={`${cell.letter}-${cell.state}`} in timeout={1000}>
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
            </Fade>
          </Grid>
        )),
      )}
    </Grid>
  )
}

export default Board
