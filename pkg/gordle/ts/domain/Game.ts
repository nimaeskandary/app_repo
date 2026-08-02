import { CellState, type Cell } from '@app_repo/pkg_gordle/domain/Cell'

const rubricWords = [
  'CRANE',
  'SLATE',
  'BRICK',
  'CLOUD',
  'GHOST',
  'PLANT',
  'SHORE',
  'MIGHT',
  'FEAST',
  'WOMAN',
]

// Creates a blank board for a new game.
export function createBoard(): Cell[][] {
  return Array.from({ length: 6 }, () =>
    Array.from({ length: 5 }, () => ({ letter: '', state: CellState.Unguessed })),
  )
}

// Selects the answer for a new game.
export function createAnswer(): string {
  return rubricWords[Math.floor(Math.random() * rubricWords.length)]
}

// Scores a guess against the answer.
export function scoreGuess(guess: string, answer: string): Cell[] {
  const row = Array.from({ length: 5 }, (_, index) => ({
    letter: guess[index],
    state: CellState.Unguessed,
  }))

  if (row.some((cell) => !cell.letter)) {
    return row
  }

  const states = row.map(() => CellState.Wrong)
  const remainingLetters = answer.split('')

  row.forEach((cell, index) => {
    if (cell.letter === remainingLetters[index]) {
      states[index] = CellState.Correct
      remainingLetters[index] = ''
    }
  })

  row.forEach((cell, index) => {
    if (states[index] === CellState.Correct) {
      return
    }

    const answerIndex = remainingLetters.indexOf(cell.letter)
    if (answerIndex !== -1) {
      states[index] = CellState.RowCorrect
      remainingLetters[answerIndex] = ''
    }
  })

  return row.map((cell, index) => ({ ...cell, state: states[index] }))
}
