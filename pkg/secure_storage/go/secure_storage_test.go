package secure_storage_test

import (
	"testing"

	secure_storage "github.com/nimaeskandary/app_repo/pkg/secure_storage/go"
	secure_storage_core "github.com/nimaeskandary/app_repo/pkg/secure_storage/go/core"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
	"go.uber.org/fx"
)

func TestNewSecureStorageModule(t *testing.T) {
	keyring.MockInit()

	var storage secure_storage_core.SecureStorage
	app := fx.New(
		fx.Supply(secure_storage_core.Config{Namespace: "test-service"}),
		secure_storage.NewSecureStorageModule(),
		fx.Populate(&storage),
		fx.NopLogger,
	)

	require.NoError(t, app.Start(t.Context()))
	require.NotNil(t, storage)
	require.NoError(t, app.Stop(t.Context()))
}
