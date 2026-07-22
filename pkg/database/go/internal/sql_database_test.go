package internal

import (
	"database/sql"
	"testing"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

	return newSQLDatabase(sqlDB, db_types.DialectSQLite), sqlDB
}
