package internal

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"
	"testing/fstest"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGooseMigrator(t *testing.T) {
	t.Parallel()

	t.Run("should provide a migrator", func(t *testing.T) {
		t.Parallel()

		database := newTestSQLiteWriter(t, filepath.Join(t.TempDir(), "test.db"))

		migrator, err := NewGooseMigrator(database, testMigrationSource())

		require.NoError(t, err)
		require.NoError(t, migrator.Start(t.Context()))
		assert.NotNil(t, migrator)
	})
}

func TestGooseMigrator(t *testing.T) {
	t.Parallel()

	t.Run("Up", func(t *testing.T) {
		t.Parallel()

		t.Run("should run all pending SQL and code migrations", func(t *testing.T) {
			t.Parallel()

			migrator, database := newTestMigrator(t)

			require.NoError(t, migrator.Up(t.Context(), nil))

			assert.True(t, tableExists(t, database.DB(), "first_records"))
			assert.True(t, tableExists(t, database.DB(), "code_records"))
			assert.True(t, tableExists(t, database.DB(), "third_records"))
		})

		t.Run("should run an exact version without earlier versions", func(t *testing.T) {
			t.Parallel()

			migrator, database := newTestMigrator(t)
			version := int64(3)

			require.NoError(t, migrator.Up(t.Context(), &version))

			assert.False(t, tableExists(t, database.DB(), "first_records"))
			assert.False(t, tableExists(t, database.DB(), "code_records"))
			assert.True(t, tableExists(t, database.DB(), "third_records"))
		})

		t.Run("should run missing versions after an out-of-order version", func(t *testing.T) {
			t.Parallel()

			migrator, database := newTestMigrator(t)
			version := int64(3)
			require.NoError(t, migrator.Up(t.Context(), &version))

			require.NoError(t, migrator.Up(t.Context(), nil))

			assert.True(t, tableExists(t, database.DB(), "first_records"))
			assert.True(t, tableExists(t, database.DB(), "code_records"))
			assert.True(t, tableExists(t, database.DB(), "third_records"))
		})

		t.Run("should ignore an already applied exact version", func(t *testing.T) {
			t.Parallel()

			migrator, _ := newTestMigrator(t)
			version := int64(1)
			require.NoError(t, migrator.Up(t.Context(), &version))

			assert.NoError(t, migrator.Up(t.Context(), &version))
		})
	})

	t.Run("Down", func(t *testing.T) {
		t.Parallel()

		t.Run("should revert an exact version while later versions remain", func(t *testing.T) {
			t.Parallel()

			migrator, database := newTestMigrator(t)
			require.NoError(t, migrator.Up(t.Context(), nil))

			require.NoError(t, migrator.Down(t.Context(), 1))

			assert.False(t, tableExists(t, database.DB(), "first_records"))
			assert.True(t, tableExists(t, database.DB(), "code_records"))
			assert.True(t, tableExists(t, database.DB(), "third_records"))
		})

		t.Run("should revert a code migration", func(t *testing.T) {
			t.Parallel()

			migrator, database := newTestMigrator(t)
			require.NoError(t, migrator.Up(t.Context(), nil))

			require.NoError(t, migrator.Down(t.Context(), 2))

			assert.True(t, tableExists(t, database.DB(), "first_records"))
			assert.False(t, tableExists(t, database.DB(), "code_records"))
			assert.True(t, tableExists(t, database.DB(), "third_records"))
		})
	})

	t.Run("Stop should reset the provider", func(t *testing.T) {
		t.Parallel()

		migrator, _ := newTestMigrator(t)

		require.NoError(t, migrator.Stop(t.Context()))

		assert.EqualError(t, migrator.Up(t.Context(), nil), "migrator is not started")
		assert.EqualError(t, migrator.Down(t.Context(), 1), "migrator is not started")
	})

	t.Run("should allow Start after Stop", func(t *testing.T) {
		t.Parallel()

		migrator, database := newTestMigrator(t)
		require.NoError(t, migrator.Stop(t.Context()))
		require.NoError(t, database.Stop(t.Context()))
		require.NoError(t, database.Start(t.Context()))

		require.NoError(t, migrator.Start(t.Context()))
		require.NoError(t, migrator.Up(t.Context(), nil))

		assert.True(t, tableExists(t, database.DB(), "first_records"))
		assert.True(t, tableExists(t, database.DB(), "code_records"))
		assert.True(t, tableExists(t, database.DB(), "third_records"))
	})
}

func TestGooseDialect(t *testing.T) {
	t.Parallel()

	t.Run("should map SQLite", func(t *testing.T) {
		t.Parallel()

		dialect, err := gooseDialect(db_core.DialectSQLite)

		require.NoError(t, err)
		assert.Equal(t, goose.DialectSQLite3, dialect)
	})

	t.Run("should map PostgreSQL", func(t *testing.T) {
		t.Parallel()

		dialect, err := gooseDialect(db_core.DialectPostgres)

		require.NoError(t, err)
		assert.Equal(t, goose.DialectPostgres, dialect)
	})

	t.Run("should reject an unsupported dialect", func(t *testing.T) {
		t.Parallel()

		dialect, err := gooseDialect("unsupported")

		assert.Empty(t, dialect)
		assert.EqualError(t, err, "unsupported SQL dialect: unsupported")
	})
}

func newTestMigrator(t *testing.T) (db_core.Migrator, db_core.SQLDatabase) {
	t.Helper()

	database := newTestSQLiteWriter(t, filepath.Join(t.TempDir(), "test.db"))
	migrator, err := NewGooseMigrator(database, testMigrationSource())
	require.NoError(t, err)
	require.NoError(t, migrator.Start(t.Context()))

	return migrator, database
}

func testMigrationSource() db_core.MigrationSource {
	return db_core.MigrationSource{
		SQLFiles: fstest.MapFS{
			"0001_create_first_records.up.sql":   {Data: []byte("CREATE TABLE first_records (value TEXT);")},
			"0001_create_first_records.down.sql": {Data: []byte("DROP TABLE first_records;")},
			"0003_create_third_records.up.sql":   {Data: []byte("CREATE TABLE third_records (value TEXT);")},
			"0003_create_third_records.down.sql": {Data: []byte("DROP TABLE third_records;")},
		},
		CodeMigrations: []db_core.Migration{{
			Version: 2,
			Name:    "create_code_records",
			Up: func(ctx context.Context, tx *sql.Tx) error {
				_, err := tx.ExecContext(ctx, "CREATE TABLE code_records (value TEXT);")
				return err
			},
			Down: func(ctx context.Context, tx *sql.Tx) error {
				_, err := tx.ExecContext(ctx, "DROP TABLE code_records;")
				return err
			},
		}},
	}
}

func tableExists(t *testing.T, db *sql.DB, tableName string) bool {
	t.Helper()

	var exists int
	require.NoError(t, db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = ?)",
		tableName,
	).Scan(&exists))
	return exists == 1
}
