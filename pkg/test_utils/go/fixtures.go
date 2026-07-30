package test_utils

import (
	"context"
	"testing"

	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	observability "github.com/nimaeskandary/app_repo/pkg/observability/go"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

// StandardFixture is used by tests to access instantiated components from a common dependency graph. As new components
// are added to the codebase, they can be added to this fixture
type StandardFixture struct {
	Logger obs_core.Logger
}

// SetupStandardFixture sets up a standard fixture with all dependencies injected.
//
// To use overrides, pass in functions that return the type being overriden in the dependency tree,
// this will replace the default constructor for that type, e.g.
// override1 := func(dep SomeDepFromTree, etc) TypeBeingOverriden { return myNewmockImplementation(dep1) }
// SetupStandardFixture(t, override1)
func SetupStandardFixture(t *testing.T, overrides ...any) StandardFixture {
	testModules := TestModules()

	for _, constructorFn := range overrides {
		testModules = append(testModules, fx.Decorate(constructorFn))
	}

	f := StandardFixture{}

	fxApp := di.CreateFxAppAndExtract(
		testModules,
		&f.Logger,
	)

	require.NoError(t, fxApp.Start(t.Context()))
	t.Cleanup(func() { _ = fxApp.Stop(context.Background()) })

	return f
}

func TestModules() []fx.Option {
	return []fx.Option{
		fx.Provide(
			func() Config {
				return NewTestConfig()
			},
			func(c Config) obs_core.SlogLoggerConfig {
				return c.Slog
			},
		),
		observability.NewSlogLoggerModule(),
	}
}
