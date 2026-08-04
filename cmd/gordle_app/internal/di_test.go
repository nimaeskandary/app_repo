package internal

import (
	"path/filepath"
	"testing"

	app_database "github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/database/app"
	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
	"go.uber.org/fx"
)

func TestModuleList(t *testing.T) {
	keyring.MockInit()
	dbDir := t.TempDir()
	source := filepath.Join(dbDir, "app.db")
	config := Config{
		AppConfig:   AppConfig{DataDir: dbDir},
		AppDatabase: db_core.SQLiteConfig{DBFilename: "app.db"},
		Logger:      obs_core.SlogLoggerConfig{Level: "DEBUG"},
	}

	var appDBReader app_database.AppDBReader
	var appDBWriter app_database.AppDBWriter
	var appDatabaseMigrator app_database.AppDatabaseMigrator
	var appDatabaseMigrateAllOnStart app_database.AppDatabaseMigrateAllOnStart
	var logger obs_core.Logger
	app := fx.New(
		append(
			ModuleList(config),
			fx.Populate(
				&appDBReader,
				&appDBWriter,
				&appDatabaseMigrator,
				&appDatabaseMigrateAllOnStart,
				&logger,
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

	assert.NotNil(t, logger)
	readerSQLDB := appDBReader.DB()
	writerSQLDB := appDBWriter.DB()
	require.NoError(t, app.Stop(t.Context()))
	assert.Nil(t, appDBReader.DB())
	assert.Nil(t, appDBWriter.DB())
	assert.Error(t, readerSQLDB.Ping())
	assert.Error(t, writerSQLDB.Ping())
}
