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

	t.Run("should provide separate reader and writer databases", func(t *testing.T) {
		t.Parallel()

		dbDir := t.TempDir()
		config := db_core.SQLiteConfig{DBFilename: "app.db"}
		var reader sqliteReader
		var writer sqliteWriter
		app := di.CreateFxAppAndExtract(
			[]fx.Option{
				NewSQLiteWriterModule[sqliteWriter](config, dbDir),
				NewSQLiteReaderModule[sqliteReader](config, dbDir),
			},
			&reader,
			&writer,
		)

		require.NoError(t, app.Start(t.Context()))
		assert.NotSame(t, reader.DB(), writer.DB())
		require.NoError(t, app.Stop(t.Context()))
	})

	t.Run("should pair multiple databases with their migrators", func(t *testing.T) {
		t.Parallel()

		var firstDB firstDatabase
		var secondDB secondDatabase
		firstDBDir := filepath.Join(t.TempDir(), "first")
		secondDBDir := filepath.Join(t.TempDir(), "second")
		app := di.CreateFxAppAndExtract(
			[]fx.Option{
				NewSQLiteWriterModule[firstDatabase](db_core.SQLiteConfig{DBFilename: "first.db"}, firstDBDir),
				NewSQLiteWriterModule[secondDatabase](db_core.SQLiteConfig{DBFilename: "second.db"}, secondDBDir),
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
type sqliteReader db_core.SQLDatabase
type sqliteWriter db_core.SQLDatabase
type firstMigrator db_core.Migrator
type secondMigrator db_core.Migrator
type firstMigrateAllOnStart db_core.MigrateAllOnStart
type secondMigrateAllOnStart db_core.MigrateAllOnStart

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
