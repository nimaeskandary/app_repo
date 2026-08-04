package secure_storage

import (
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	secure_storage_core "github.com/nimaeskandary/app_repo/pkg/secure_storage/go/core"
	"github.com/nimaeskandary/app_repo/pkg/secure_storage/go/internal"
	"go.uber.org/fx"
)

// NewSecureStorageModule provides configured platform-backed secret storage.
func NewSecureStorageModule() fx.Option {
	return di.NewFxModule[secure_storage_core.SecureStorage]("secure_storage", internal.NewSecureStorage)
}
