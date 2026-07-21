import { Box } from '@/pkg/ui/ts/mui/box'
import { Container } from '@/pkg/ui/ts/mui/container'
import { Grid } from '@/pkg/ui/ts/mui/grid'
import { H4 } from '@/pkg/ui/ts/mui/typography'

export enum GordleCellState {
  Unguessed = 'Unguessed',
  Wrong = 'Wrong',
  RowCorrect = 'RowCorrect',
  Correct = 'Correct',
}

export type GordleCell = {
  letter: string
  state: GordleCellState
}

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
                  }}>
                  <H4>{cell.letter}</H4>
                </Box>
              </Grid>
            )),
          )}
        </Grid>
    </Container>
  )
}

export default GordleView
