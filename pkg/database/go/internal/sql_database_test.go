package internal

import (
	"database/sql"
	"path/filepath"
	"testing"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSQLDatabase(t *testing.T) {
	t.Parallel()

	t.Run("should ping the database", func(t *testing.T) {
		t.Parallel()

		sqlDB, err := sql.Open("sqlite", ":memory:")
		require.NoError(t, err)
		t.Cleanup(func() { _ = sqlDB.Close() })

		db, err := newSQLDatabase(sqlDB, db_types.DialectSQLite)

		require.NoError(t, err)
		assert.NotNil(t, db)
	})

	t.Run("should close the database when ping fails", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "missing", "test.db")
		sqlDB, err := sql.Open("sqlite", source)
		require.NoError(t, err)

		db, err := newSQLDatabase(sqlDB, db_types.DialectSQLite)

		assert.Nil(t, db)
		assert.ErrorContains(t, err, "ping sqlite database")
		assert.Error(t, sqlDB.Ping())
	})
}

func TestSQLDatabase(t *testing.T) {
	t.Parallel()

	t.Run("DB", func(t *testing.T) {
		t.Parallel()

		db, sqlDB := newTestSQLDatabase(t)

		assert.Same(t, sqlDB, db.DB())
	})

	t.Run("Dialect", func(t *testing.T) {
		t.Parallel()

		db, _ := newTestSQLDatabase(t)

		assert.Equal(t, db_types.DialectSQLite, db.Dialect())
	})

	t.Run("Stop", func(t *testing.T) {
		t.Parallel()

		t.Run("should close the database", func(t *testing.T) {
			t.Parallel()

			db, sqlDB := newTestSQLDatabase(t)

			require.NoError(t, db.Stop(t.Context()))
			assert.Error(t, sqlDB.Ping())
		})

		t.Run("should allow repeated calls", func(t *testing.T) {
			t.Parallel()

			db, _ := newTestSQLDatabase(t)

			require.NoError(t, db.Stop(t.Context()))
			assert.NoError(t, db.Stop(t.Context()))
		})
	})
}

func newTestSQLDatabase(t *testing.T) (db_types.SQLDatabase, *sql.DB) {
	t.Helper()

	sqlDB, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = sqlDB.Close() })

	db, err := newSQLDatabase(sqlDB, db_types.DialectSQLite)
	require.NoError(t, err)

	return db, sqlDB
}
