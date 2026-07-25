import { createContext, type PropsWithChildren, useContext } from 'react'
import type { GordleNavigator } from '../navigation/GordleNavigator'

export type GordleDependencies = {
  navigator: GordleNavigator
}

export type GordleDependenciesProviderProps = PropsWithChildren<{
  value: GordleDependencies
}>

const GordleDependenciesContext = createContext<GordleDependencies | undefined>(undefined)

// Makes the Gordle dependency graph available to view models.
export function GordleDependenciesProvider({ value, children }: GordleDependenciesProviderProps) {
  return (
    <GordleDependenciesContext.Provider value={value}>
      {children}
    </GordleDependenciesContext.Provider>
  )
}

// Returns the dependency graph configured by the current application.
export function useGordleDependencies() {
  const dependencies = useContext(GordleDependenciesContext)
  if (!dependencies) {
    throw new Error('GordleDependenciesProvider is missing')
  }

  return dependencies
}
