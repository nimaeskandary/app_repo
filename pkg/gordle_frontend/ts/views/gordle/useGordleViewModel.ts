import { useState } from 'react'
import {
  GordleCellState,
  GordleCell,
} from '@app_repo/pkg_gordle_frontend/model/GordleCell'

export type GordleViewModel = {
  cells: GordleCell[][]
}

// Creates a blank board for a new game.
function createInitialCells(): GordleCell[][] {
  return Array.from({ length: 6 }, () =>
    Array.from({ length: 5 }, () => ({ letter: '', state: GordleCellState.Unguessed })),
  )
}

// Owns the state displayed by the Gordle view.
export function useGordleViewModel(): GordleViewModel {
  const [cells] = useState(createInitialCells)

  return { cells }
}
