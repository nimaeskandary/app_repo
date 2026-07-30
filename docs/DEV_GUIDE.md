# Dev Guide

## Table of Contents

1. [Setup](#setup)
2. [Running locally](#running-locally)
3. [Git hooks](#git-hooks)

## Setup

### Dependencies

* go 1.27rc2
* node 26
* docker

After installing these dependencies, you can run 

* `npm install`
* `bin/go mod tidy`

### Wails

* install [wails3](https://v3.wails.io/quick-start/installation/) as a global go tool via go install. Run through the setup wizard that will check if you have things like xcode, virtual devices, etc setup. You will want at least one mobile emulator set up on your system.
* run `wails3 doctor` - this should say your system is ready, or at least has the components you care about

## Running locally

See the READMEs at different project roots to see how to run them:

* `cmd/gordle_app/README.md`

## Git Hooks

run `git config --local core.hooksPath .githooks/` to use this repos git hooks
