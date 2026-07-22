package internal

import (
	"testing"

	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestModuleListLoadsConfig(t *testing.T) {
	var loggerConfig obs_types.SlogLoggerConfig
	app := fx.New(
		append(ModuleList([]byte(`{"Logger":{"Level":"DEBUG"}}`)), fx.Populate(&loggerConfig), fx.NopLogger)...,
	)

	require.NoError(t, app.Start(t.Context()))
	assert.Equal(t, "DEBUG", loggerConfig.Level)
	require.NoError(t, app.Stop(t.Context()))
}
