# wails3-react-polylith

This repo contains a Wails v3 sample app reorganized into a [Polylith](https://polylith.gitbook.io/polylith)-style layout.

## Project structure

```text
# A base is an encapsulated block of code that can be assembled together with a set of components and 
# libraries into services, libraries or tools. Bases achieve encapsulation and composability by
# separating their private implementation from their public API.
bases/                      
  wails_app/
    main.go                 # Wails app entrypoint
    frontend/               # React/Vite frontend for this Wails base

# A component is an encapsulated block of code that can be assembled together with a base.
# Components achieve encapsulation and composability by separating their private implementation from their public interface.
components/
  greet/
    go/                     # Go component, e.g. domain types, public interfaces
      internal/             # Concrete interface implementations

# A project is an edge for artifact generation assembling bases and components. 
projects/
  wails_app/
    Taskfile.yml            # Thin artifact declaration
    build/                  # App identity, config, and native platform assets

go.mod                      # single Go module for all Go code
package.json                # npm workspace root
```

Go code uses the root `go.mod`. 

Frontend code uses npm workspaces so future JavaScript/TypeScript bases can share the root `node_modules` and lockfile.

## Install dependencies

From the repo root:

```sh
npm install
```

## Wails commands

Because the Wails project config lives under `projects/wails_app/build/config.yml`, plain `wails3 dev` from the repo root does not use the correct config.

* run dev mode on host: `wails3 task -dir projects/wails_app dev`
* run dev mode on ios: `wails3 task -dir projects/wails_app ios:run`
* build: `wails3 task -dir projects/wails_app build`
* generate typescript bindings: `wails3 task -dir projects/wails_app common:generate:bindings`

## Frontend-only commands

These do not run the Go backend or native Wails app. They only run the Vite/TypeScript frontend workspace.

```sh
npm run wails-app:dev
npm run wails-app:build:dev
npm run wails-app:build
npm run wails-app:preview
```

## Storybook

Storybook is used to preview frontend components. See `./storybook`

To launch the local storybook server: `npm run storybook`

## TypeScript imports

The frontend uses TypeScript 7-compatible path aliases without `baseUrl`.

Current aliases:

```text
@/*            -> bases/wails_app/frontend/src/*
@bindings/*   -> bases/wails_app/frontend/bindings/*
@components/* -> components/*
```

Generated Wails services are re-exported from:

```text
bases/wails_app/frontend/src/wails-services.ts
```
