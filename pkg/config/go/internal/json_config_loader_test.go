package internal

import (
	"encoding/json/v2"
	"errors"
	"testing"

	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testConfig struct {
	Name   string
	Secret config_core.SecretString
}

type testCustomValue string

func TestNewJsonConfigLoader(t *testing.T) {
	t.Parallel()

	t.Run("loads config", func(t *testing.T) {
		t.Parallel()

		loader, err := NewJsonConfigLoader[testConfig](
			[]byte(`{"Name":"gordle","Secret":"token"}`),
			nil,
		)

		require.NoError(t, err)
		assert.Equal(t, testConfig{Name: "gordle", Secret: "token"}, loader.GetConfig())
	})

	t.Run("returns an error for invalid JSON", func(t *testing.T) {
		t.Parallel()

		_, err := NewJsonConfigLoader[testConfig]([]byte(`{"Name":`), nil)

		assert.Error(t, err)
	})

	t.Run("returns an error when a custom unmarshaler fails", func(t *testing.T) {
		t.Parallel()

		unmarshalErr := errors.New("custom unmarshaler failed")
		_, err := NewJsonConfigLoader[testConfig](
			[]byte(`{"Secret":"token"}`),
			[]*json.Unmarshalers{
				json.UnmarshalFunc(func([]byte, *config_core.SecretString) error {
					return unmarshalErr
				}),
			},
		)

		assert.ErrorIs(t, err, unmarshalErr)
	})

	t.Run("returns an error when a required field is missing", func(t *testing.T) {
		t.Parallel()

		_, err := NewJsonConfigLoader[struct {
			Name string `validate:"required"`
		}]([]byte(`{}`), nil)

		assert.ErrorContains(t, err, "failed to validate JSON config")
	})

	t.Run("loads custom types", func(t *testing.T) {
		loader, err := NewJsonConfigLoader[struct {
			Value testCustomValue
		}](
			[]byte(`{"Value":"gordle"}`),
			[]*json.Unmarshalers{
				json.UnmarshalFunc(func(data []byte, value *testCustomValue) error {
					var raw string
					if err := json.Unmarshal(data, &raw); err != nil {
						return err
					}
					*value = testCustomValue("custom:" + raw)
					return nil
				}),
			},
		)

		require.NoError(t, err)
		assert.Equal(t, testCustomValue("custom:gordle"), loader.GetConfig().Value)
	})
}
