import { useState } from 'react'
import { Box } from '@app_repo/pkg_ui/mui/box'
import { ErrorButton, PrimaryButton } from '@app_repo/pkg_ui/mui/button'
import { Grid } from '@app_repo/pkg_ui/mui/grid'
import { Stack } from '@app_repo/pkg_ui/mui/stack'
import { Body1 } from '@app_repo/pkg_ui/mui/typography'
import BackspaceIcon from '@mui/icons-material/Backspace'
import InputIcon from '@mui/icons-material/Input'

export type KeyboardProps = {
  onWordCommit: (word: string) => boolean
}

const letters = Array.from({ length: 26 }, (_, index) => String.fromCharCode(65 + index))

// Renders an alphabet keyboard and stages a word for submission.
function Keyboard({ onWordCommit }: KeyboardProps) {
  const [stagedWord, setStagedWord] = useState('')

  // Submits a complete word and clears it only when accepted.
  const commitStagedWord = () => {
    if (stagedWord.length !== 5) {
      return
    }

    if (onWordCommit(stagedWord)) {
      setStagedWord('')
    }
  }

  return (
    <Stack spacing={1}>
      <Stack direction="row" spacing={1}>
        <ErrorButton
          aria-label="Delete last letter"
          disabled={!stagedWord}
          onClick={() => setStagedWord((currentWord) => currentWord.slice(0, -1))}>
            <BackspaceIcon />
        </ErrorButton>
        <Box
          data-testid="staged-word"
          sx={{
            border: 2,
            borderColor: '#3a3a3c',
            borderRadius: 2,
            color: '#ffffff',
            display: 'grid',
            flex: 1,
            minHeight: 40,
            placeItems: 'center',
          }}>
          <Body1>{stagedWord}</Body1>
        </Box>
        <PrimaryButton
          aria-label="Commit staged word"
          data-testid="commit-word"
          disabled={stagedWord.length !== 5}
          onClick={commitStagedWord}>
          <InputIcon />
        </PrimaryButton>
      </Stack>
      <Grid container columns={7} spacing={1}>
        {letters.map((letter) => (
          <Grid key={letter} size={1}>
            <Box
              aria-label={`Select ${letter}`}
              component="button"
              data-testid={`select-letter-${letter}`}
              onClick={() =>
                setStagedWord((currentWord) =>
                  currentWord.length < 5 ? `${currentWord}${letter}` : currentWord,
                )
              }
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
      </Grid>
    </Stack>
  )
}

export default Keyboard
