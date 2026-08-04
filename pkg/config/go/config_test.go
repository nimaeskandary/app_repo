package config_test

import (
	"encoding/json/v2"
	"errors"
	"testing"

	config "github.com/nimaeskandary/app_repo/pkg/config/go"
	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestLoadJsonConfig(t *testing.T) {
	t.Parallel()

	t.Run("loads config", func(t *testing.T) {
		t.Parallel()

		loaded, err := config.LoadJsonConfig[struct {
			Name string `validate:"required"`
		}](
			[]byte(`{"Name":"gordle"}`),
			nil,
		)

		require.NoError(t, err)
		assert.Equal(t, "gordle", loaded.Name)
	})

	t.Run("returns loading errors", func(t *testing.T) {
		t.Parallel()

		unmarshalErr := errors.New("custom unmarshaler failed")
		_, err := config.LoadJsonConfig[struct {
			Secret config_core.SecretString
		}](
			[]byte(`{"Secret":"token"}`),
			[]*json.Unmarshalers{
				json.UnmarshalFunc(func([]byte, *config_core.SecretString) error {
					return unmarshalErr
				}),
			},
		)

		assert.ErrorIs(t, err, unmarshalErr)
		assert.ErrorContains(t, err, "load JSON config")
	})
}

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
