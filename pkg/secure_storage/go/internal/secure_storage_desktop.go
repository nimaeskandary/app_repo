//go:build !android && !ios

package internal

import (
	"errors"

	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"github.com/nimaeskandary/app_repo/pkg/secure_storage/go/core"
	"github.com/zalando/go-keyring"
)

type secureStorage struct {
	di.NoOpLifecycle
	namespace string
}

// NewSecureStorage creates secret storage backed by the system keyring.
func NewSecureStorage(config secure_storage_core.Config) secure_storage_core.SecureStorage {
	return &secureStorage{namespace: config.Namespace}
}

func (s *secureStorage) Set(key, value string) error {
	return keyring.Set(s.namespace, key, value)
}

func (s *secureStorage) Get(key string) (string, error) {
	value, err := keyring.Get(s.namespace, key)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", nil
	}
	return value, err
}

func (s *secureStorage) Delete(key string) error {
	err := keyring.Delete(s.namespace, key)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}
