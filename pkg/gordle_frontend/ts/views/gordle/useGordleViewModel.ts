import { useState } from 'react'
import {
  GordleCellState,
  type GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewModel = {
  cells: GordleCell[][]
  commitWord: (word: string) => boolean
  isGameComplete: boolean
}

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
function createInitialCells(): GordleCell[][] {
  return Array.from({ length: 6 }, () =>
    Array.from({ length: 5 }, () => ({ letter: '', state: GordleCellState.Unguessed })),
  )
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

// Owns the state displayed by the Gordle view.
export function useGordleViewModel(): GordleViewModel {
  const [cells, setCells] = useState(createInitialCells)
  const [rubricWord] = useState(
    () => rubricWords[Math.floor(Math.random() * rubricWords.length)],
  )
  const isGameComplete = cells.some(
    (row) => row.map((cell) => cell.letter).join('') === rubricWord,
  )

  // Adds a complete word to the first empty board row.
  const commitWord = (word: string) => {
    if (isGameComplete) {
      return false
    }

    const rowIndex = cells.findIndex((row) => row.every((cell) => !cell.letter))
    if (rowIndex === -1) {
      return false
    }

    setCells((currentCells) => {
      const nextCells = [...currentCells]
      const committedRow = currentCells[rowIndex].map((cell, index) => ({
        ...cell,
        letter: word[index],
      }))
      const rowStates = getRowStates(committedRow, rubricWord)
      nextCells[rowIndex] = committedRow.map((cell, index) => ({
        ...cell,
        state: rowStates[index],
      }))

      return nextCells
    })

    return true
  }

  return { cells, commitWord, isGameComplete }
}
