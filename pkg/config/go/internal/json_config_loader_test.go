package internal

import (
	"errors"
	"testing"

	config_types "github.com/nimaeskandary/app_repo/pkg/config/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testConfig struct {
	Name   string
	Secret config_types.SecretString
}

type testSecretParser struct {
	err error
}

func (p testSecretParser) Parse(raw string) (string, error) {
	if p.err != nil {
		return "", p.err
	}
	return "resolved:" + raw, nil
}

func TestNewJsonConfigLoader(t *testing.T) {
	t.Parallel()

	t.Run("loads config and resolves secrets", func(t *testing.T) {
		t.Parallel()

		loader, err := NewJsonConfigLoader[testConfig](testSecretParser{}, []byte(`{"Name":"gordle","Secret":"token"}`))

		require.NoError(t, err)
		assert.Equal(t, testConfig{Name: "gordle", Secret: "resolved:token"}, loader.GetConfig())
	})

	t.Run("returns an error for invalid JSON", func(t *testing.T) {
		t.Parallel()

		_, err := NewJsonConfigLoader[testConfig](testSecretParser{}, []byte(`{"Name":`))

		assert.Error(t, err)
	})

	t.Run("returns an error when resolving a secret fails", func(t *testing.T) {
		t.Parallel()

		parserErr := errors.New("secret unavailable")
		_, err := NewJsonConfigLoader[testConfig](testSecretParser{err: parserErr}, []byte(`{"Secret":"token"}`))

		assert.ErrorIs(t, err, parserErr)
	})
}
