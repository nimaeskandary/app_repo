import { useCallback } from 'react'
import { useDependencies } from '@app_repo/pkg_gordle/app/Dependencies'

export type MainMenuViewModel = {
  play: () => void
}

// Exposes main-menu state and commands to its view.
export function useMainMenuViewModel(): MainMenuViewModel {
  const { navigator } = useDependencies()
  const play = useCallback(() => navigator.goToGame(), [navigator])

  return { play }
}
