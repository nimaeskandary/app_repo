import { useState } from 'react'
import {
  GordleCellState,
  type GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewModel = {
  cells: GordleCell[][]
  commitWord: (word: string) => boolean
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

  // Adds a valid word to the first empty board row.
  const commitWord = (word: string) => {
    if (!rubricWords.includes(word)) {
      return false
    }

    const rowIndex = cells.findIndex((row) => row.every((cell) => !cell.letter))
    if (rowIndex === -1) {
      return false
    }

    setCells((currentCells) => {
      const nextCells = [...currentCells]
      nextCells[rowIndex] = currentCells[rowIndex].map((cell, index) => ({
        ...cell,
        letter: word[index],
      }))

      return nextCells
    })

    return true
  }

  return { cells, commitWord, rubricWord }
}
