import { createContext, type PropsWithChildren, useContext } from 'react'
import type { Navigator } from '@app_repo/pkg_gordle/app/Navigator'

export type Dependencies = {
  navigator: Navigator
}

export type DependenciesProviderProps = PropsWithChildren<{
  value: Dependencies
}>

const DependenciesContext = createContext<Dependencies | undefined>(undefined)

// Makes the Gordle dependency graph available to view models.
export function DependenciesProvider({ value, children }: DependenciesProviderProps) {
  return (
    <DependenciesContext.Provider value={value}>
      {children}
    </DependenciesContext.Provider>
  )
}

// Returns the dependency graph configured by the current application.
export function useDependencies() {
  const dependencies = useContext(DependenciesContext)
  if (!dependencies) {
    throw new Error('DependenciesProvider is missing')
  }

  return dependencies
}
