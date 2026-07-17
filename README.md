# wails3-react-polylith

This repo contains a Wails v3 sample app reorganized into a Polylith-style layout.

## Project structure

```text
bases/
  wails_app/
    main.go                 # Wails app entrypoint
    frontend/               # React/Vite frontend for this Wails base

components/
  greet/
    greetservice.go         # Go service component bound into Wails

projects/
  wails_app/
    Taskfile.yml            # Wails task entrypoint for this project
    build/                  # Wails build config and platform assets

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
