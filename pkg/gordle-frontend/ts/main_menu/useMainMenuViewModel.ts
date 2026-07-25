import { useCallback } from 'react'
import { useGordleDependencies } from '../di/GordleDependencies'

export type MainMenuViewModel = {
  play: () => void
}

// Exposes main-menu state and commands to its view.
export function useMainMenuViewModel(): MainMenuViewModel {
  const { navigator } = useGordleDependencies()
  const play = useCallback(() => navigator.goToGame(), [navigator])

  return { play }
}
