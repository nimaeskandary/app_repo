package internal

import (
	"fmt"
	"path/filepath"
	"testing"

	app_database "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app"
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestModuleListLoadsConfig(t *testing.T) {
	source := filepath.Join(t.TempDir(), "app.db")
	configBytes := []byte(fmt.Sprintf(
		`{"AppDatabase":{"AppDir":"Gordle","DbFile":%q},"Logger":{"Level":"DEBUG"}}`,
		source,
	))

	var appDatabase app_database.AppDatabase
	var appDatabaseMigrator app_database.AppDatabaseMigrator
	var appDatabaseMigrateAll app_database.AppDatabaseMigrateAll
	var loggerConfig obs_core.SlogLoggerConfig
	app := fx.New(
		append(
			ModuleList(configBytes),
			fx.Populate(
				&appDatabase,
				&appDatabaseMigrator,
				&appDatabaseMigrateAll,
				&loggerConfig,
			),
			fx.NopLogger,
		)...,
	)

	require.NoError(t, app.Start(t.Context()))
	assert.Equal(t, db_core.DialectSQLite, appDatabase.Dialect())
	assert.FileExists(t, source)
	assert.NotNil(t, appDatabaseMigrator)
	assert.NotNil(t, appDatabaseMigrateAll)

	var migrationTable string
	require.NoError(t, appDatabase.DB().QueryRow(
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'goose_db_version'",
	).Scan(&migrationTable))
	assert.Equal(t, "goose_db_version", migrationTable)

	assert.Equal(t, "DEBUG", loggerConfig.Level)
	require.NoError(t, app.Stop(t.Context()))
	assert.Error(t, appDatabase.DB().Ping())
}
