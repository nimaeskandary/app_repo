import { useState } from 'react'
import { Box } from '@app_repo/pkg_ui/mui/box'
import { PrimaryButton } from '@app_repo/pkg_ui/mui/button'
import { Grid } from '@app_repo/pkg_ui/mui/grid'
import { Stack } from '@app_repo/pkg_ui/mui/stack'
import { Body1 } from '@app_repo/pkg_ui/mui/typography'

export type KeyboardProps = {
  onLetterCommit: (letter: string) => void
}

const letters = Array.from({ length: 26 }, (_, index) => String.fromCharCode(65 + index))

// Renders an alphabet keyboard and commits the currently selected letter.
function Keyboard({ onLetterCommit }: KeyboardProps) {
  const [selectedLetter, setSelectedLetter] = useState('')

  // Commits a selection and clears the landing area.
  const commitSelectedLetter = () => {
    if (!selectedLetter) {
      return
    }

    onLetterCommit(selectedLetter)
    setSelectedLetter('')
  }

  return (
    <Stack spacing={1}>
      <Box
        sx={{
          border: 2,
          borderColor: '#3a3a3c',
          borderRadius: 2,
          color: '#ffffff',
          display: 'grid',
          minHeight: 40,
          placeItems: 'center',
        }}>
        <Body1>{selectedLetter}</Body1>
      </Box>
      <Grid container columns={9} spacing={1}>
        {letters.map((letter) => (
          <Grid key={letter} size={1}>
            <Box
              aria-label={`Select ${letter}`}
              className="not-draggable"
              component="button"
              onClick={() => setSelectedLetter(letter)}
              sx={{
                aspectRatio: '1',
                backgroundColor: 'transparent',
                border: 2,
                borderColor: '#3a3a3c',
                borderRadius: 2,
                color: '#ffffff',
                containerType: 'inline-size',
                cursor: 'pointer',
                display: 'grid',
                padding: 0,
                placeItems: 'center',
                width: '100%',
              }}>
              <Body1 sx={{ fontSize: '70cqi', lineHeight: 1 }}>{letter}</Body1>
            </Box>
          </Grid>
        ))}
        <Grid size={1}>
          <PrimaryButton
            aria-label="Commit selected letter"
            className="not-draggable"
            disabled={!selectedLetter}
            onClick={commitSelectedLetter}
            sx={{
              aspectRatio: '1',
              containerType: 'inline-size',
              minWidth: 0,
              padding: 0,
              width: '100%',
            }}>
            <Body1 sx={{ fontSize: '70cqi', lineHeight: 1 }}>→</Body1>
          </PrimaryButton>
        </Grid>
      </Grid>
    </Stack>
  )
}

export default Keyboard
