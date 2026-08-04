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

pkg/go/foo/core/foo.go
```go
package foo_core

type Foo interface {
    // our di.NewFxModule requires that your component implements lifecycle. Its okay to just make
    // Start and Stop implementations simple no ops if needed via di.NoOpLifecycle
    di.Lifecycle
    FooBehavior()
}
```

pkg/go/bar/core/bar.go
```go
package bar_core

type Bar interface {
    di.Lifecycle
    BarBehavior()
}
```

pkg/go/foo/internal/some_foo.go
```go
package internal

type someFoo struct {
    // can use this if the component has no need for Start or Stop lifecycle hooks
    di.NoOpLifecycle
    bar bar_core.Bar
}

func NewSomeFoo (bar bar_core.Bar) foo_core.Foo {
    return &someFoo{
        bar: bar
    }
}

func (c *myComponentImpl) FooBehavior() {}
```

pkg/go/foo/foo.go
```go
package foo

func NewSomeFooModule() fx.Option {
    return di.NewFxModule[foo_core.Foo]("foo", internal.NewSomeFoo)
}
```

Often your component will just rely on other components in the dependency tree, but in the event you need to pass in something custom at the edge that is not in the dep tree, you can use this pattern:

pkg/go/foo/foo.go
```go
package foo

func NewSomeFooModule(edgeComponent EdgeComponent) fx.Option {
    constructor := func(bar bar_core.Bar) foo_core.Foo {
        return internal.NewSomeFoo(bar, edgeComponent)
    }

    return di.NewFxModule[my_component_core.MyComponent]("foo", constructor)
}
```

This pattern allows you to add a wrapped Foo constructor that still resolves bar from the dep tree, but is supplied the EdgeComponent out of band.

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

    // Create a t.Run boundary for an interface method
	t.Run("BehaviorFoo", func(t *testing.T) {
		t.Parallel()

        // Create a t.Run boundary for each test case
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

        t.Run("Should do bar edge case A", func(t *testing.T) {
        }

        t.Run("Should handle bar error path B", func(t *testing.T) {
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
