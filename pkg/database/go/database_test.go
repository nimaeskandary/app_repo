package database

import (
	"path/filepath"
	"testing"
	"testing/fstest"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
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
				NewMigrateAllOnStartModule[firstMigrator, firstMigrateAllOnStart](),
				NewMigrateAllOnStartModule[secondMigrator, secondMigrateAllOnStart](),
			},
			&firstDB,
			&secondDB,
		)

		require.NoError(t, app.Start(t.Context()))
		assert.True(t, databaseTableExists(t, firstDB, "first_records"))
		assert.False(t, databaseTableExists(t, firstDB, "second_records"))
		assert.True(t, databaseTableExists(t, secondDB, "second_records"))
		assert.False(t, databaseTableExists(t, secondDB, "first_records"))
		firstSQLDB := firstDB.DB()
		secondSQLDB := secondDB.DB()

		require.NoError(t, app.Stop(t.Context()))
		assert.Nil(t, firstDB.DB())
		assert.Nil(t, secondDB.DB())
		assert.Error(t, firstSQLDB.Ping())
		assert.Error(t, secondSQLDB.Ping())
	})
}

type firstDatabase db_core.SQLDatabase
type secondDatabase db_core.SQLDatabase
type firstMigrator db_core.Migrator
type secondMigrator db_core.Migrator
type firstMigrateAllOnStart db_core.MigrateAllOnStart
type secondMigrateAllOnStart db_core.MigrateAllOnStart

type firstSQLiteConfig db_core.SQLiteConfig

func (c firstSQLiteConfig) SQLiteConfig() db_core.SQLiteConfig {
	return db_core.SQLiteConfig(c)
}

type secondSQLiteConfig db_core.SQLiteConfig

func (c secondSQLiteConfig) SQLiteConfig() db_core.SQLiteConfig {
	return db_core.SQLiteConfig(c)
}

func testSQLMigrationSource(tableName string) db_core.MigrationSource {
	return db_core.MigrationSource{SQLFiles: fstest.MapFS{
		"0001_create_records.up.sql":   {Data: []byte("CREATE TABLE " + tableName + " (value TEXT);")},
		"0001_create_records.down.sql": {Data: []byte("DROP TABLE " + tableName + ";")},
	}}
}

func databaseTableExists(t *testing.T, database db_core.SQLDatabase, tableName string) bool {
	t.Helper()

	var exists int
	require.NoError(t, database.DB().QueryRow(
		"SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = ?)",
		tableName,
	).Scan(&exists))
	return exists == 1
}
