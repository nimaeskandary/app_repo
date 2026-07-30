# Go Guide

> Note, it is recomended to use the `bin/go` wrapper for all go commands, e.g. `bin/go test ./...`

## Table of Contents

1. [Module structure](#module-structure)
2. [Dependency injection](#dependency-injection)
3. [Testing](#testing)
    1. [Testing guide](#testing-guide)
    1. [Mocks](#mocks)

## Module structure

The typical structure used in this repo is:

```
my_package/
    core/ - public types, functions, etc
    internal/ - private implementations of core types
    my_package.go - export constructors of private implementations
```

## Dependency injection

This project uses [fx](https://github.com/uber-go/fx) for golang dependency injection.

Say you want to implement some component that does foo and bar. To make this component testable, auto wired to interact with downstream and upstream consumers components, and more easily replaced with a different implementation, follow this format:

core/my_component.go
```go
type MyComponent interface {
    BehaviorFoo()
    BehaviorBar()
}
```

internal/my_component_impl.go
```go
type myComponentImpl struct {}

// any dependencies on other components can be in this constructor's params
func NewMyComponentImpl MyComponent() {

}

func (c *myComponentImpl) BehaviorBar() {}
func (c *myComponentImpl) BehaviorFoo() {}
```

my_component.go
```go
func NewMyComponentImplModule fx.Option {
    return di.NewFxModule[core.MyComponent]("my_component", internal.NewMyComponentImpl)
}
```

## Testing

* run `bin/go test ./...` to run all tests
* run `bin/go test <path-to-package> to run tests for a specific package

### Style

When writing tests, aim to follow this format. Imagine you were testing 

```go
type InterfaceBeingTested interface {
    BehaviorFoo()
    BehaviorBar()
}

type myInterfaceImpl struct {}
```

The test cases should generally aim to be parallel if possible. They should be behavioral driven based on the interface being tested.

```go
import (
    "github.com/nimaeskandary/app_repo/pkg/test_utils/go"
)

func Test_MyInterfaceImpl(t *testing.T) {
	t.Parallel()

    // aim to use the standard fixture when applicable
	f := test_utils.SetupStandardFixture(t)
	underTest := f.InterfaceBeingTested

	t.Run("BehaviorFoo", func(t *testing.T) {
		t.Parallel()

		t.Run("Should do foo happy path", func(t *testing.T) {
        }

        t.Run("Should handle foo error path", func(t *testing.T) {
        }
        // etc
    }

    t.Run("BehaviorBar", func(t *testing.T) {
		t.Parallel()

		t.Run("Should do bar happy path", func(t *testing.T) {
        }

        t.Run("Should handle bar error path", func(t *testing.T) {
        }
        // etc
    }
}
```

### Mocks

Aim to use mocks instead of creating fake structs/ test implementations of things, if the goal is just to record calls or return dummy data for a component

* this project uses https://vektra.github.io/mockery
* to mark an interface for mock generation, use the comment `//mockery:generate: true`
* to generate mocks, run `bin/generate-mocks`
* mocks will be generated in a `mocks/` subfolder in the package of the interface, e.g. `pkg/user/go/core/mocks`
