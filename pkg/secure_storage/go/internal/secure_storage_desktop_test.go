//go:build !android && !ios

package internal

import (
	"errors"
	"strconv"
	"testing"
	"testing/synctest"
	"time"

	secure_storage_core "github.com/nimaeskandary/app_repo/pkg/secure_storage/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func TestSecureStorage(t *testing.T) {
	config := secure_storage_core.Config{
		Namespace: "test-service",
	}

	t.Run("Start", func(t *testing.T) {
		keyring.MockInit()
		storage := NewSecureStorage(config)

		t.Run("stores the last start timestamp", func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				require.NoError(t, storage.Start(t.Context()))

				value, err := storage.Get(lastStartTimestampKey)

				require.NoError(t, err)
				assert.Equal(t, strconv.FormatInt(time.Now().UnixMilli(), 10), value)
			})
		})

		t.Run("propagates storage errors", func(t *testing.T) {
			keyringErr := errors.New("keyring failed")
			keyring.MockInitWithError(keyringErr)

			assert.ErrorIs(t, storage.Start(t.Context()), keyringErr)
		})
	})

	t.Run("Set", func(t *testing.T) {
		keyring.MockInit()
		storage := NewSecureStorage(config)

		t.Run("stores a secret", func(t *testing.T) {
			require.NoError(t, storage.Set("token", "first"))

			value, err := storage.Get("token")

			require.NoError(t, err)
			assert.Equal(t, "first", value)
		})

		t.Run("overwrites a secret", func(t *testing.T) {
			require.NoError(t, storage.Set("overwrite", "first"))
			require.NoError(t, storage.Set("overwrite", "second"))

			value, err := storage.Get("overwrite")

			require.NoError(t, err)
			assert.Equal(t, "second", value)
		})

		t.Run("propagates storage errors", func(t *testing.T) {
			keyringErr := errors.New("keyring failed")
			keyring.MockInitWithError(keyringErr)

			assert.ErrorIs(t, storage.Set("token", "value"), keyringErr)
		})
	})

	t.Run("Get", func(t *testing.T) {
		keyring.MockInit()
		storage := NewSecureStorage(config)

		t.Run("retrieves a secret", func(t *testing.T) {
			require.NoError(t, storage.Set("get", "value"))

			value, err := storage.Get("get")

			require.NoError(t, err)
			assert.Equal(t, "value", value)
		})

		t.Run("returns an empty value when a secret is missing", func(t *testing.T) {
			value, err := storage.Get("missing")

			require.NoError(t, err)
			assert.Empty(t, value)
		})

		t.Run("propagates storage errors", func(t *testing.T) {
			keyringErr := errors.New("keyring failed")
			keyring.MockInitWithError(keyringErr)

			_, err := storage.Get("token")
			assert.ErrorIs(t, err, keyringErr)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		keyring.MockInit()
		storage := NewSecureStorage(config)

		t.Run("deletes a secret", func(t *testing.T) {
			require.NoError(t, storage.Set("delete", "value"))

			require.NoError(t, storage.Delete("delete"))
			value, err := storage.Get("delete")

			require.NoError(t, err)
			assert.Empty(t, value)
		})

		t.Run("succeeds when the secret is missing", func(t *testing.T) {
			assert.NoError(t, storage.Delete("missing-delete"))
		})

		t.Run("propagates storage errors", func(t *testing.T) {
			keyringErr := errors.New("keyring failed")
			keyring.MockInitWithError(keyringErr)

			assert.ErrorIs(t, storage.Delete("token"), keyringErr)
		})
	})
}
