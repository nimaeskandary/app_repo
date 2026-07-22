package internal

import (
	"context"
	"database/sql"
	"testing"
	"testing/fstest"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadMigrations(t *testing.T) {
	t.Parallel()

	t.Run("should load SQL and code migrations in version order", func(t *testing.T) {
		t.Parallel()

		codeMigration := db_types.Migration{
			Version: 2,
			Name:    "rewrite_records",
			Up:      func(context.Context, *sql.Tx) error { return nil },
			Down:    func(context.Context, *sql.Tx) error { return nil },
		}
		source := db_types.MigrationSource{
			SQLFiles: fstest.MapFS{
				"0001_create_records.up.sql":   {Data: []byte("CREATE TABLE records (value TEXT);")},
				"0001_create_records.down.sql": {Data: []byte("DROP TABLE records;")},
				"0003_add_index.up.sql":        {Data: []byte("CREATE INDEX records_value ON records(value);")},
				"0003_add_index.down.sql":      {Data: []byte("DROP INDEX records_value;")},
			},
			CodeMigrations: []db_types.Migration{codeMigration},
		}

		migrations, err := loadMigrations(source)

		require.NoError(t, err)
		assert.Equal(t, []int64{1, 2, 3}, []int64{
			migrations[0].Version,
			migrations[1].Version,
			migrations[2].Version,
		})
		assert.Equal(t, []string{"create_records", "rewrite_records", "add_index"}, []string{
			migrations[0].Name,
			migrations[1].Name,
			migrations[2].Name,
		})
	})

	t.Run("should execute SQL migration functions", func(t *testing.T) {
		t.Parallel()

		migrations, err := loadMigrations(db_types.MigrationSource{SQLFiles: fstest.MapFS{
			"0001_create_records.up.sql":   {Data: []byte("CREATE TABLE records (value TEXT);")},
			"0001_create_records.down.sql": {Data: []byte("DROP TABLE records;")},
		}})
		require.NoError(t, err)
		db, err := sql.Open("sqlite", ":memory:")
		require.NoError(t, err)
		t.Cleanup(func() { _ = db.Close() })
		tx, err := db.BeginTx(t.Context(), nil)
		require.NoError(t, err)
		t.Cleanup(func() { _ = tx.Rollback() })

		require.NoError(t, migrations[0].Up(t.Context(), tx))
		var tableName string
		require.NoError(t, tx.QueryRow("SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'records'").Scan(&tableName))
		assert.Equal(t, "records", tableName)

		require.NoError(t, migrations[0].Down(t.Context(), tx))
		assert.ErrorIs(t, tx.QueryRow("SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'records'").Scan(&tableName), sql.ErrNoRows)
	})

	t.Run("should reject duplicate versions", func(t *testing.T) {
		t.Parallel()

		migrationFunc := func(context.Context, *sql.Tx) error { return nil }
		source := db_types.MigrationSource{
			SQLFiles: fstest.MapFS{
				"0001_create_records.up.sql":   {Data: []byte("SELECT 1;")},
				"0001_create_records.down.sql": {Data: []byte("SELECT 1;")},
			},
			CodeMigrations: []db_types.Migration{{Version: 1, Name: "code", Up: migrationFunc, Down: migrationFunc}},
		}

		_, err := loadMigrations(source)

		assert.EqualError(t, err, "duplicate migration version: 1")
	})

	t.Run("should reject invalid code migrations", func(t *testing.T) {
		t.Parallel()

		t.Run("without a positive version", func(t *testing.T) {
			t.Parallel()

			_, err := loadMigrations(db_types.MigrationSource{CodeMigrations: []db_types.Migration{{Version: 0}}})

			assert.EqualError(t, err, "migration version must be positive: 0")
		})

		t.Run("without a name", func(t *testing.T) {
			t.Parallel()

			migrationFunc := func(context.Context, *sql.Tx) error { return nil }
			_, err := loadMigrations(db_types.MigrationSource{CodeMigrations: []db_types.Migration{{Version: 1, Up: migrationFunc, Down: migrationFunc}}})

			assert.EqualError(t, err, "migration 1 name is required")
		})

		t.Run("without an up function", func(t *testing.T) {
			t.Parallel()

			migrationFunc := func(context.Context, *sql.Tx) error { return nil }
			_, err := loadMigrations(db_types.MigrationSource{CodeMigrations: []db_types.Migration{{Version: 1, Name: "test", Down: migrationFunc}}})

			assert.EqualError(t, err, "migration 1 up function is required")
		})

		t.Run("without a down function", func(t *testing.T) {
			t.Parallel()

			migrationFunc := func(context.Context, *sql.Tx) error { return nil }
			_, err := loadMigrations(db_types.MigrationSource{CodeMigrations: []db_types.Migration{{Version: 1, Name: "test", Up: migrationFunc}}})

			assert.EqualError(t, err, "migration 1 down function is required")
		})
	})
}

func TestLoadSQLMigrations(t *testing.T) {
	t.Parallel()

	t.Run("should reject malformed SQL filenames", func(t *testing.T) {
		t.Parallel()

		_, err := loadSQLMigrations(fstest.MapFS{
			"create_records.sql": {Data: []byte("SELECT 1;")},
		})

		assert.EqualError(t, err, "invalid SQL migration filename: create_records.sql")
	})

	t.Run("should require an up and down pair", func(t *testing.T) {
		t.Parallel()

		t.Run("without up SQL", func(t *testing.T) {
			t.Parallel()

			_, err := loadSQLMigrations(fstest.MapFS{
				"0001_create_records.down.sql": {Data: []byte("DROP TABLE records;")},
			})

			assert.EqualError(t, err, "migration 1 is missing up SQL")
		})

		t.Run("without down SQL", func(t *testing.T) {
			t.Parallel()

			_, err := loadSQLMigrations(fstest.MapFS{
				"0001_create_records.up.sql": {Data: []byte("CREATE TABLE records (value TEXT);")},
			})

			assert.EqualError(t, err, "migration 1 is missing down SQL")
		})
	})

	t.Run("should reject mismatched pair names", func(t *testing.T) {
		t.Parallel()

		_, err := loadSQLMigrations(fstest.MapFS{
			"0001_create_records.up.sql": {Data: []byte("SELECT 1;")},
			"0001_drop_records.down.sql": {Data: []byte("SELECT 1;")},
		})

		assert.ErrorContains(t, err, "migration 1 has mismatched names")
	})
}
