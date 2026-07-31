# Frontend Guide

## Table of Contents

1. [MVVM](#mvvm)
2. [Dependency injection](#dependency-injection)
3. [Design system](#design-system)
4. [Storybook](#storybook)
5. [Testing](#testing)

## MVVM

This repo uses [Model–view–viewmodel](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93viewmodel) as a general frontend pattern. It fits nicely with functional composition. You can also take a look at Flutter's [docs](https://docs.flutter.dev/app-architecture/guide) on this topic. At a high level:

* The View should just be responsible for how to render a UI given a passed in state
* The ViewModel is the bridge between the View and the Model layer, they build and maintain the state for the View and handle actions performend by the View
* The Model is the data layer, i.e. service and repository layer. The View should not interact with this directly

This aids us in a few things. All of these layers can more easily be tested. Views can more easily be reused and previewed. ViewModels and Models can utilize dependency injection to compose different implementations for different targets. For example, say you develop a NoteView, and NoteViewModel. The ViewModel needs a SaveNoteService to handle the action of saving Notes. On Mobile, the SaveNoteService on the dependency graph may be one that saves to disk in a Sqlite DB. On Web, this may save to local storage.  

## Dependency injection

### React

Dependency injection for React can be achieved via [Contexts](https://react.dev/learn/passing-data-deeply-with-context). e.g.

Define the dependency tree:

DependencyTree.ts
```ts
export type DepTree = {
  serviceA: ServiceA
  serviceB: ServiceB
}

export type DepTreeProviderProps = PropsWithChildren<{
  value: DepTree
}>

const DepTreeContext = createContext<DepTree | undefined>(undefined)

// Makes the dependency graph available to view models.
export function DepTreeProvider({ value, children }: DepTreeProviderProps) {
  return (
    <DepTreeContext.Provider value={value}>
      {children}
    </DepTreeContext.Provider>
  )
}

// Returns the dependency graph configured by the current application.
export function useDependencies() {
  const dependencies = useContext(DepTreeContext)
  if (!dependencies) {
    throw new Error('DepTreeProvider is missing')
  }

  return dependencies
}
```

Build the dependency tree at the edge of your application and wrap the component tree in it:

App.tsx
```ts
function App() {
    const serviceA = ...
    const serviceB = ...
    const dependencies = useMemo<DepTree>(() => ({
        serviceA: serviceA
        serviceB: serviceB
    }), [serviceA, serviceB])

    return (
        <DepTreeProvider value={dependencies}>
            <Routes>
                <Route path="/" element={<MainMenuPage />} />
            </Routes>
        </DepTreeProvider>
    )
}
```

Pull these out of the dependency tree as needed in your View Models. 

useFooViewModel.tsx
```ts
export type FooViewModel = {
  behaviorA: () => void
}

export function useFooViewModel(): FooViewModel {
  const { serviceA } = useDependencies()
  const behaviorA = useCallback(() => serviceA.doA(), [serviceA])

  return { behaviorA }
}
```

## Design system

### Material UI

While not always possible, aim to [Material UI components](https://mui.com/material-ui/all-components)  as frontend building blocks. Instead of using these as direct imports, thin wrappers are maintained at `pkg/ui/ts/mui` that should be used instead. If something that seems like it makes sense to use from MUI is not there already, go ahead and add it. 

## Storybook

Storybook is used to preview frontend components. See `storybook/`

Colocate `*.stories.tsx` files with the Views they cover. Define each reusable View state as a story driven by props:

```tsx
import type { Meta, StoryObj } from '@storybook/react-vite'
import FooView from './FooView'

const meta = {
  title: 'Foo/Views/Foo',
  component: FooView,
} satisfies Meta<typeof FooView>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: { isComplete: false },
}

export const Complete: Story = {
  args: { isComplete: true },
}
```

Run Storybook with `npm run storybook`. Use the **App shell** toolbar to preview package stories in different application layouts.

## Testing

* this project uses
    * [Vitest](https://vitest.dev/guide/)
    * [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
* for Views, it is reasonable that storybook is used an alternative to writing Vitest cases
* run all frontend tests with `npm test`
* run one test file with `npm test -- <path-to-test>`
* run tests in watch mode with `npm run test:watch`
* colocate `*.test.ts(x)` files with the code they cover

For example, say you had a ViewModel `FootViewModel`:

```ts
import { act, renderHook } from '@testing-library/react'
import { describe, expect, it } from 'vitest'
import { useFooViewModel } from './useFooViewModel'

describe('useFooViewModel', () => {
  describe('increment', () => {
    it('should increment the count', () => {
      const { result } = renderHook(() => useFooViewModel())

      act(() => {
        result.current.increment()
      })

      expect(result.current.count).toBe(1)
    })
  })
})
```
