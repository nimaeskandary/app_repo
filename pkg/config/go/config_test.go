package config_test

import (
	"testing"

	config "github.com/nimaeskandary/app_repo/pkg/config/go"
	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestNewJsonConfigLoaderModule(t *testing.T) {
	t.Parallel()

	type testConfig struct {
		Name string
	}

	var loader config_core.ConfigLoader[testConfig]
	app := fx.New(
		config.NewJsonConfigLoaderModule[testConfig]([]byte(`{"Name":"gordle"}`), nil),
		fx.Populate(&loader),
		fx.NopLogger,
	)

	require.NoError(t, app.Start(t.Context()))
	assert.Equal(t, testConfig{Name: "gordle"}, loader.GetConfig())
	require.NoError(t, app.Stop(t.Context()))
}
