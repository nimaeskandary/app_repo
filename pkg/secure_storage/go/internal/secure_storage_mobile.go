//go:build android || ios

package internal

import (
	"fmt"

	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"github.com/nimaeskandary/app_repo/pkg/secure_storage/go/core"
	wails "github.com/nimaeskandary/app_repo/pkg/wails/go"
)

type secureStorage struct {
	di.NoOpLifecycle
}

// NewSecureStorage creates secret storage backed by the mobile platform.
func NewSecureStorage(_ string) secure_storage_core.SecureStorage {
	return &secureStorage{}
}

func (s *secureStorage) Set(key, value string) error {
	err := wails.SecureSet(key, value)
	if err != nil {
		return fmt.Errorf("unable to save key %v: %w", key, err)
	}
	if wails.SecureGet(key) != value {
		return fmt.Errorf("unable to save key %v", key)
	}
	return nil
}

func (s *secureStorage) Get(key string) (string, error) {
	return wails.SecureGet(key), nil
}

func (s *secureStorage) Delete(key string) error {
	wails.SecureDelete(key)
	return nil
}
