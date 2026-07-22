package database

import (
	"path/filepath"
	"testing"
	"testing/fstest"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestDatabaseModules(t *testing.T) {
	t.Parallel()

	t.Run("should pair multiple databases with their migrators", func(t *testing.T) {
		t.Parallel()

		var firstDB firstDatabase
		var secondDB secondDatabase
		app := di.CreateFxAppAndExtract(
			[]fx.Option{
				fx.Supply(firstSQLiteConfig{Source: filepath.Join(t.TempDir(), "first.db")}),
				fx.Supply(secondSQLiteConfig{Source: filepath.Join(t.TempDir(), "second.db")}),
				NewSQLiteDatabaseModule[firstDatabase, firstSQLiteConfig](),
				NewSQLiteDatabaseModule[secondDatabase, secondSQLiteConfig](),
				NewMigratorModule[firstDatabase, firstMigrator](testSQLMigrationSource("first_records")),
				NewMigratorModule[secondDatabase, secondMigrator](testSQLMigrationSource("second_records")),
				NewMigrateAllModule[firstMigrator, firstMigrateAll](),
				NewMigrateAllModule[secondMigrator, secondMigrateAll](),
			},
			&firstDB,
			&secondDB,
		)

		require.NoError(t, app.Start(t.Context()))
		assert.True(t, databaseTableExists(t, firstDB, "first_records"))
		assert.False(t, databaseTableExists(t, firstDB, "second_records"))
		assert.True(t, databaseTableExists(t, secondDB, "second_records"))
		assert.False(t, databaseTableExists(t, secondDB, "first_records"))

		require.NoError(t, app.Stop(t.Context()))
		assert.Error(t, firstDB.DB().Ping())
		assert.Error(t, secondDB.DB().Ping())
	})
}

type firstDatabase db_types.SQLDatabase
type secondDatabase db_types.SQLDatabase
type firstMigrator db_types.Migrator
type secondMigrator db_types.Migrator
type firstMigrateAll db_types.MigrateAll
type secondMigrateAll db_types.MigrateAll

type firstSQLiteConfig db_types.SQLiteConfig

func (c firstSQLiteConfig) SQLiteConfig() db_types.SQLiteConfig {
	return db_types.SQLiteConfig(c)
}

type secondSQLiteConfig db_types.SQLiteConfig

func (c secondSQLiteConfig) SQLiteConfig() db_types.SQLiteConfig {
	return db_types.SQLiteConfig(c)
}

func testSQLMigrationSource(tableName string) db_types.MigrationSource {
	return db_types.MigrationSource{SQLFiles: fstest.MapFS{
		"0001_create_records.up.sql":   {Data: []byte("CREATE TABLE " + tableName + " (value TEXT);")},
		"0001_create_records.down.sql": {Data: []byte("DROP TABLE " + tableName + ";")},
	}}
}

func databaseTableExists(t *testing.T, database db_types.SQLDatabase, tableName string) bool {
	t.Helper()

	var exists int
	require.NoError(t, database.DB().QueryRow(
		"SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = ?)",
		tableName,
	).Scan(&exists))
	return exists == 1
}
