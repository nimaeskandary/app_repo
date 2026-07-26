import { Box } from '@app_repo/pkg_ui/mui/box'
import { Grid } from '@app_repo/pkg_ui/mui/grid'
import { Fade } from '@app_repo/pkg_ui/mui/transitions'
import { H3 } from '@app_repo/pkg_ui/mui/typography'
import {
  GordleCellState,
  type GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type BoardProps = {
  cells: GordleCell[][]
  rubricWord: string
}

const cellStyles = {
  [GordleCellState.Unguessed]: { borderColor: '#3a3a3c' },
  [GordleCellState.Wrong]: { backgroundColor: '#3a3a3c', borderColor: '#3a3a3c' },
  [GordleCellState.RowCorrect]: { backgroundColor: '#b59f3b', borderColor: '#b59f3b' },
  [GordleCellState.Correct]: { backgroundColor: '#538d4e', borderColor: '#538d4e' },
}

// Scores a completed row against the rubric word.
function getRowStates(row: GordleCell[], rubricWord: string): GordleCellState[] {
  if (row.some((cell) => !cell.letter)) {
    return row.map(() => GordleCellState.Unguessed)
  }

  const states = row.map(() => GordleCellState.Wrong)
  const remainingLetters = rubricWord.split('')

  row.forEach((cell, index) => {
    if (cell.letter === remainingLetters[index]) {
      states[index] = GordleCellState.Correct
      remainingLetters[index] = ''
    }
  })

  row.forEach((cell, index) => {
    if (states[index] === GordleCellState.Correct) {
      return
    }

    const rubricIndex = remainingLetters.indexOf(cell.letter)

    if (rubricIndex !== -1) {
      states[index] = GordleCellState.RowCorrect
      remainingLetters[rubricIndex] = ''
    }
  })

  return states
}

// Renders the Gordle cell grid.
function Board({ cells, rubricWord }: BoardProps) {
  return (
    <Grid container columns={5} spacing={1}>
      {cells.flatMap((row, rowIndex) => {
        const rowStates = getRowStates(row, rubricWord)

        return row.map((cell, columnIndex) => (
          <Grid key={`${rowIndex}-${columnIndex}`} size={1}>
            <Fade key={`${cell.letter}-${rowStates[columnIndex]}`} in>
              <Box
                sx={{
                  ...cellStyles[rowStates[columnIndex]],
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
        ))
      })}
    </Grid>
  )
}

export default Board
