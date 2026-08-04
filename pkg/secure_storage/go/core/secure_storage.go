package secure_storage_core

import di "github.com/nimaeskandary/app_repo/pkg/di/go"

// SecureStorage stores and retrieves application secrets.
type SecureStorage interface {
	di.Lifecycle
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}
