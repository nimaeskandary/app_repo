import { useState } from 'react'
import type { Cell } from '@app_repo/pkg_gordle/domain/Cell'
import { createAnswer, createBoard, scoreGuess } from '@app_repo/pkg_gordle/domain/Game'

export type GordleViewModel = {
  cells: Cell[][]
  commitWord: (word: string) => boolean
  isGameComplete: boolean
}

// Owns the state displayed by the Gordle view.
export function useGordleViewModel(): GordleViewModel {
  const [cells, setCells] = useState(createBoard)
  const [rubricWord] = useState(createAnswer)
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
      nextCells[rowIndex] = scoreGuess(word, rubricWord)

      return nextCells
    })

    return true
  }

  return { cells, commitWord, isGameComplete }
}
