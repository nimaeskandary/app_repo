package internal

import (
	"database/sql"
	"errors"
	"path/filepath"
	"testing"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSQLDatabase(t *testing.T) {
	t.Parallel()

	t.Run("should open and ping the database during Start", func(t *testing.T) {
		t.Parallel()

		sqlDB, err := sql.Open("sqlite", ":memory:")
		require.NoError(t, err)
		t.Cleanup(func() { _ = sqlDB.Close() })

		db := newSQLDatabase(func() (*sql.DB, error) {
			return sqlDB, nil
		}, db_core.DialectSQLite)

		assert.Nil(t, db.DB())
		require.NoError(t, db.Start(t.Context()))
		assert.Same(t, sqlDB, db.DB())
	})

	t.Run("should close the database when ping fails", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "missing", "test.db")
		sqlDB, err := sql.Open("sqlite", source)
		require.NoError(t, err)

		db := newSQLDatabase(func() (*sql.DB, error) {
			return sqlDB, nil
		}, db_core.DialectSQLite)

		err = db.Start(t.Context())

		assert.ErrorContains(t, err, "ping sqlite database")
		assert.Nil(t, db.DB())
		assert.Error(t, sqlDB.Ping())
	})
}

func TestSQLDatabase(t *testing.T) {
	t.Parallel()

	t.Run("should expose the configured dialect", func(t *testing.T) {
		t.Parallel()

		db, _ := newTestSQLDatabase(t)

		assert.Equal(t, db_core.DialectSQLite, db.Dialect())
	})

	t.Run("should close the database", func(t *testing.T) {
		t.Parallel()

		db, sqlDB := newTestSQLDatabase(t)

		require.NoError(t, db.Stop(t.Context()))
		assert.Error(t, sqlDB.Ping())
	})

	t.Run("should allow repeated Stop calls", func(t *testing.T) {
		t.Parallel()

		db, _ := newTestSQLDatabase(t)

		require.NoError(t, db.Stop(t.Context()))
		assert.NoError(t, db.Stop(t.Context()))
	})

	t.Run("should reset and allow Start after Stop", func(t *testing.T) {
		t.Parallel()

		var opened []*sql.DB
		db := newSQLDatabase(func() (*sql.DB, error) {
			sqlDB, err := sql.Open("sqlite", ":memory:")
			if err == nil {
				opened = append(opened, sqlDB)
				t.Cleanup(func() { _ = sqlDB.Close() })
			}
			return sqlDB, err
		}, db_core.DialectSQLite)

		require.NoError(t, db.Start(t.Context()))
		first := db.DB()
		require.NoError(t, db.Stop(t.Context()))

		assert.Nil(t, db.DB())
		assert.Error(t, first.Ping())

		require.NoError(t, db.Start(t.Context()))
		second := db.DB()
		require.NotNil(t, second)
		assert.NotSame(t, first, second)
		assert.Len(t, opened, 2)
		assert.NoError(t, second.Ping())
	})

	t.Run("should allow Stop before Start", func(t *testing.T) {
		t.Parallel()

		sqlDB, err := sql.Open("sqlite", ":memory:")
		require.NoError(t, err)
		t.Cleanup(func() { _ = sqlDB.Close() })

		db := newSQLDatabase(func() (*sql.DB, error) {
			return sqlDB, nil
		}, db_core.DialectSQLite)

		require.NoError(t, db.Stop(t.Context()))
		require.NoError(t, db.Start(t.Context()))
		require.NoError(t, db.Stop(t.Context()))

		assert.Nil(t, db.DB())
		assert.Error(t, sqlDB.Ping())
	})

	t.Run("should retry Start after a failed Start and Stop", func(t *testing.T) {
		t.Parallel()

		startError := errors.New("open failed")
		var attempts int
		db := newSQLDatabase(func() (*sql.DB, error) {
			attempts++
			if attempts == 1 {
				return nil, startError
			}

			sqlDB, err := sql.Open("sqlite", ":memory:")
			if err == nil {
				t.Cleanup(func() { _ = sqlDB.Close() })
			}
			return sqlDB, err
		}, db_core.DialectSQLite)

		assert.ErrorIs(t, db.Start(t.Context()), startError)
		require.NoError(t, db.Stop(t.Context()))
		require.NoError(t, db.Start(t.Context()))

		assert.Equal(t, 2, attempts)
		assert.NotNil(t, db.DB())
	})
}

func newTestSQLDatabase(t *testing.T) (db_core.SQLDatabase, *sql.DB) {
	t.Helper()

	sqlDB, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = sqlDB.Close() })

	db := newSQLDatabase(func() (*sql.DB, error) {
		return sqlDB, nil
	}, db_core.DialectSQLite)
	require.NoError(t, db.Start(t.Context()))

	return db, sqlDB
}
