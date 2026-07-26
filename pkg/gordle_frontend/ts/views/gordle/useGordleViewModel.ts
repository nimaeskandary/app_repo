import { useState } from 'react'
import {
  GordleCellState,
  type GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewModel = {
  cells: GordleCell[][]
  commitLetter: (letter: string) => void
  rubricWord: string
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

// Owns the state displayed by the Gordle view.
export function useGordleViewModel(): GordleViewModel {
  const [cells, setCells] = useState(createInitialCells)
  const [rubricWord] = useState(
    () => rubricWords[Math.floor(Math.random() * rubricWords.length)],
  )

  // Adds a letter to the first empty board cell.
  const commitLetter = (letter: string) => {
    setCells((currentCells) => {
      const rowIndex = currentCells.findIndex((row) => row.some((cell) => !cell.letter))

      if (rowIndex === -1) {
        return currentCells
      }

      const columnIndex = currentCells[rowIndex].findIndex((cell) => !cell.letter)
      const nextCells = [...currentCells]
      nextCells[rowIndex] = [...currentCells[rowIndex]]
      nextCells[rowIndex][columnIndex] = {
        ...currentCells[rowIndex][columnIndex],
        letter,
      }

      return nextCells
    })
  }

  return { cells, commitLetter, rubricWord }
}
