package di

import (
	"context"

	"go.uber.org/fx"
)

// Lifecycle functions for Fx. Fx also provides a start hook, but don't see the need for that, planning to always use
// a constructor function, and we will never be restarting an fx dep tree after its created.
type Lifecycle interface {
	// Stop is called when the fx system is shutdown, used for graceful shutdown of components
	Stop(ctx context.Context) error
}

// New - Helper for creating a module for the fx dependency framework. T is the interface
// being added to the dependency graph. The constructor must return a concrete implementation of T.
// The constructor may require other components of the dependency graph, which will be resolved by fx.
//
// opts - additional fx options to add to this module
func NewFxModule[T Lifecycle](
	name string,
	constructor any,
	opts ...fx.Option) fx.Option {
	return fx.Module(
		name,
		append(
			[]fx.Option{
				fx.Provide(
					fx.Annotate(constructor, fx.As(new(T))),
				),
				fx.Invoke(registerLifecycle[T]),
			},
			opts...,
		)...,
	)
}

// CreateFxAppAndExtract creates an Fx application and extracts specified dependencies
// to be used out of the Fx application context. This is when instances of your dependencies
// need to be used by things that don't cleany fit into fx, e.g. pulling out an http server to start 
// it on the main thread, or for use with the wails framework via frontend bindings to go code, in which case
// wails needs the instances.
func CreateFxAppAndExtract(modules []fx.Option, extract ...any) *fx.App {
	return fx.New(
		append(
			modules,
			// disable fx logging
			fx.NopLogger,
			fx.Populate(extract...),
		)...,
	)
}

// registerLifecycle - wires hooks for component start and stop
func registerLifecycle[T Lifecycle](lc fx.Lifecycle, this T) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return this.Stop(ctx)
		},
	})
}
