# wails3-react

This repo contains a Wails v3 sample app in a monorepo organized around deployable apps and shared packages.

## Project structure

```text
bin/                        # dev scripts
go.mod                      # single Go module for all Go code
package.json                # npm workspace root
cmd/                        # application entrypoints
pkg/                        # reusable packages
  greet/
    go/                     # Shared go greet package
      internal/             # Private to package, e.g. concrete interface implementations
storybook/                  # frontend component previews
```



## Setup

### Dependencies

* go 1.26+
* node 26
* npm 11
* docker
* task (https://taskfile.dev/)

* run `npm install`
* install [wails3](https://v3.wails.io/quick-start/installation/) as a global go tool via go install. Run through the setup wizard that will check if you have things like xcode, virtual devices, etc setup
* run `wails3 doctor` - this should say your system is ready

## Go commands

Use `bin/go` for Go commands, for example `bin/go test ./...`

## Wails

[Wails3](https://v3.wails.io/quick-start/why-wails/) is used in this repo to build cross platform apps using go and react. This repo is generic and not every cmd executable needs to use wails, but wails is used to build targets such as `cmd/gordle_app`. 

`cmd/gordle_app` was bootstrapped with `wails3 init -n test-react -t react`.
Its Wails build files stay app-local so they can be compared directly with
future versions of the upstream template.

### Common commands

The wails3 example template ships with build scripts for all platforms wired up via the tool [task](https://taskfile.dev/). I recommend getting an IDE plugin to easily view available tasks.

Run Wails commands from `cmd/gordle_app`.

* run dev mode on host, this will hot reload: `wails3 dev`
* launch app on ios simulator: `wails3 task ios:run`
* build: `wails3 build`
* re generate TypeScript bindings: `wails3 task common:generate:bindings`, autogenerates bindings in `frontend/bindings`

## Storybook

Storybook is used to preview frontend components. See `./storybook`

To launch the local storybook server: `npm run storybook`

## Tests

### Go

run `./bin/go test ./...`

#### Mocks

* this project uses https://vektra.github.io/mockery
* to mark an interface for mock generation, use the comment `//mockery:generate: true`
* to generate mocks, run `bin/generate-mocks.sh`
* mocks will be generated in a `mocks/` subfolder in the package of the interface, e.g. `pkg/user/go/core/mocks`
