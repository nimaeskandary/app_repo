# wails3-react

This repo contains a Wails v3 sample app in a monorepo organized around deployable apps and shared packages.

## Project structure

```text
cmd/                       # entrypoints

pkg/                       # reusable packages
  greet/
    go/                     # Shared go greet package
      internal/             # Private to package, e.g. concrete interface implementations

go.mod                      # single Go module for all Go code
package.json                # npm workspace root
```

## Install dependencies



## Wails commands

Because the Wails config lives under `cmd/wordle/build/config.yml`, plain `wails3 dev` from the repo root does not use the correct config.

* run dev mode on host: `wails3 task -dir cmd/wordle dev`
* run dev mode on iOS: `wails3 task -dir cmd/wordle ios:run`
* build: `wails3 task -dir cmd/wordle build`
* generate TypeScript bindings: `wails3 task -dir cmd/wordle common:generate:bindings`

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

Prefix repository paths with `@/` when using them as imports, and import things relative to the root directory, e.g. 

* @/cmd/wordle/frontend/src/*
* @/pkg/*  

```text
@/cmd/wordle/frontend/src/* -> cmd/wordle/frontend/src/*
@/pkg/*                     -> pkg/*
```

Wails generates TypeScript bindings under:

```text
cmd/wordle/frontend/bindings/
```
