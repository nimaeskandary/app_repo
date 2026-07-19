# wails3-react-polylith

This repo contains a Wails v3 sample app in a monorepo organized around deployable apps and shared packages.

## Project structure

```text
apps/                       # entrypoints

pkgs/                       # reusable packages
  greet/
    go/                     # Shared go greet package
      internal/             # Private to package, e.g. concrete interface implementations

go.mod                      # single Go module for all Go code
package.json                # npm workspace root
```

Go code uses the root `go.mod`. 

Frontend code uses npm workspaces so future JavaScript/TypeScript packages can share the root `node_modules` and lockfile.

## Install dependencies

From the repo root:

```sh
npm install
```

## Wails commands

Because the Wails config lives under `apps/wails_app/build/config.yml`, plain `wails3 dev` from the repo root does not use the correct config.

* run dev mode on host: `wails3 task -dir apps/wails_app dev`
* run dev mode on iOS: `wails3 task -dir apps/wails_app ios:run`
* build: `wails3 task -dir apps/wails_app build`
* generate TypeScript bindings: `wails3 task -dir apps/wails_app common:generate:bindings`

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

The frontend and Storybook use a TypeScript 7-compatible repository-root alias without `baseUrl`.

Prefix repository paths with `@/`:

```text
@/apps/wails_app/frontend/src/* -> apps/wails_app/frontend/src/*
@/pkgs/*                        -> pkgs/*
```

Wails generates TypeScript bindings under:

```text
apps/wails_app/frontend/bindings/
```
