package internal

import (
	"fmt"
	"path/filepath"
	"testing"

	app_database "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app"
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	secure_storage_core "github.com/nimaeskandary/app_repo/pkg/secure_storage/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
	"go.uber.org/fx"
)

func TestModuleListLoadsConfig(t *testing.T) {
	keyring.MockInit()
	source := filepath.Join(t.TempDir(), "app.db")
	configBytes := []byte(fmt.Sprintf(
		`{"AppDatabase":{"AppDir":"Gordle","DbFile":%q},"Logger":{"Level":"DEBUG"},"SecureStorage":{"Namespace":"test.gordle"}}`,
		source,
	))

	var appDBReader app_database.AppDBReader
	var appDBWriter app_database.AppDBWriter
	var appDatabaseMigrator app_database.AppDatabaseMigrator
	var appDatabaseMigrateAllOnStart app_database.AppDatabaseMigrateAllOnStart
	var loggerConfig obs_core.SlogLoggerConfig
	var secureStorageConfig secure_storage_core.Config
	var secureStorage secure_storage_core.SecureStorage
	app := fx.New(
		append(
			ModuleList(configBytes),
			fx.Populate(
				&appDBReader,
				&appDBWriter,
				&appDatabaseMigrator,
				&appDatabaseMigrateAllOnStart,
				&loggerConfig,
				&secureStorageConfig,
				&secureStorage,
			),
			fx.NopLogger,
		)...,
	)

	require.NoError(t, app.Start(t.Context()))
	assert.Equal(t, db_core.DialectSQLite, appDBReader.Dialect())
	assert.Equal(t, db_core.DialectSQLite, appDBWriter.Dialect())
	assert.FileExists(t, source)
	assert.NotNil(t, appDatabaseMigrator)
	assert.NotNil(t, appDatabaseMigrateAllOnStart)

	var migrationTable string
	require.NoError(t, appDBReader.DB().QueryRow(
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'goose_db_version'",
	).Scan(&migrationTable))
	assert.Equal(t, "goose_db_version", migrationTable)

	assert.Equal(t, "DEBUG", loggerConfig.Level)
	assert.Equal(t, "test.gordle", secureStorageConfig.Namespace)
	assert.NotNil(t, secureStorage)
	readerSQLDB := appDBReader.DB()
	writerSQLDB := appDBWriter.DB()
	require.NoError(t, app.Stop(t.Context()))
	assert.Nil(t, appDBReader.DB())
	assert.Nil(t, appDBWriter.DB())
	assert.Error(t, readerSQLDB.Ping())
	assert.Error(t, writerSQLDB.Ping())
}
