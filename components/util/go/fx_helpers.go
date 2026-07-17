package util

import "go.uber.org/fx"

// NewFxModule creates a module for the Fx dependency framework. T is the interface
// being added to the dependency graph. The constructor must return a concrete implementation of T.
func NewFxModule[T any](
	name string,
	constructor any,
	opts ...fx.Option,
) fx.Option {
	return fx.Module(
		name,
		append(
			[]fx.Option{
				fx.Provide(
					fx.Annotate(constructor, fx.As(new(T))),
				),
			},
			opts...,
		)...,
	)
}

// CreateFxAppAndExtract creates an Fx application and extracts specified dependencies
// to be used out of the Fx application context. This should be done sparingly, at the edge
// of the system. Fx internal logging is disabled by default.
func CreateFxAppAndExtract(modules []fx.Option, extract ...any) *fx.App {
	return fx.New(
		append(
			modules,
			fx.NopLogger,
			fx.Populate(extract...),
		)...,
	)
}
