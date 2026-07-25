# Gordle Wails application

This directory keeps the Wails build flow close to a fresh `wails3 init`
project generated with Wails `v3.0.0-alpha2.117`.

## Project configuration

### Application identity
- Company, product identifier, description, copyright, and version live in
  `build/config.yml`.

## Intentional differences from `wails3 init`

Search for `Repository adaptation:` to find each corresponding build-file
change.

### Go experiment

Gordle uses `encoding/json/v2`. `GOEXPERIMENT=jsonv2` is set in the root app
Taskfile and both Wails Dockerfiles.

### Root Go module

The repository has one Go module at `../../go.mod`; Gordle does not have a
nested module. Native Go commands find the parent module automatically.
Binding-generation cache inputs explicitly include the root module files and
shared Go packages under `../../pkg`.

### npm workspace

The frontend is an npm workspace. Its lockfile, installed dependencies, and
shared dependency declarations live at the repository root. The upstream npm
tasks still run from `frontend`, but their cache inputs and outputs point at
the root workspace files.

### Shared frontend source

Gordle imports TypeScript components from `../../pkg`. Those files are added
to the Wails frontend task's source list so Task rebuilds the frontend when a
shared component changes.

### Development watcher

`build/config.yml` watches the repository root so shared source changes trigger
hot reload. It ignores dependency, generated-output, Storybook, frontend, and
fresh-init fixture directories. Since Wails executes dev commands from the
watch root, each command changes into `cmd/gordle_app` before using the normal
Wails CLI.

### Docker build context

Native Wails commands run from this directory. Cross-compilation and server
Docker builds instead mount the repository root because it owns `go.mod`, then
select `cmd/gordle_app` as the package and frontend path.

## Commands

Run commands from this directory:

```sh
cd cmd/gordle_app
wails3 dev
wails3 build
wails3 task ios:run
wails3 task common:generate:bindings
```

## Updating Wails build files

1. Generate a temporary project with the updated `wails3 init`.
2. Compare its root and `build/**/Taskfile.yml` files with this directory.
3. Apply upstream changes first.
4. Reapply only differences documented above and marked with
   `Repository adaptation:`.
5. Run binding generation, frontend build, native build, and relevant platform
   task dry runs.
