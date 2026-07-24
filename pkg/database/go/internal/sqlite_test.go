package internal

import (
	"context"
	"net/url"
	"path/filepath"
	"testing"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSQLiteDatabase(t *testing.T) {
	t.Parallel()

	t.Run("should defer opening until Start", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		db, err := NewSQLiteDatabase(db_types.SQLiteConfig{Source: source})

		require.NoError(t, err)
		assert.Nil(t, db.DB())
		assert.NoFileExists(t, source)

		require.NoError(t, db.Start(t.Context()))
		assert.FileExists(t, source)
		require.NoError(t, db.Stop(t.Context()))
	})

	t.Run("should configure SQLite defaults", func(t *testing.T) {
		t.Parallel()

		db := newTestSQLiteDatabase(t, filepath.Join(t.TempDir(), "test.db"))

		assert.Equal(t, 1, db.DB().Stats().MaxOpenConnections)

		var foreignKeysEnabled int
		require.NoError(t, db.DB().QueryRow("PRAGMA foreign_keys").Scan(&foreignKeysEnabled))
		assert.Equal(t, 1, foreignKeysEnabled)

		var busyTimeout int
		require.NoError(t, db.DB().QueryRow("PRAGMA busy_timeout").Scan(&busyTimeout))
		assert.Equal(t, 5000, busyTimeout)
	})

	t.Run("should persist data after reopening", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		db := newTestSQLiteDatabase(t, source)
		_, err := db.DB().Exec("CREATE TABLE records (value TEXT NOT NULL)")
		require.NoError(t, err)
		_, err = db.DB().Exec("INSERT INTO records (value) VALUES ('saved')")
		require.NoError(t, err)
		require.NoError(t, db.Stop(t.Context()))

		reopened := newTestSQLiteDatabase(t, source)
		var value string
		require.NoError(t, reopened.DB().QueryRow("SELECT value FROM records").Scan(&value))

		assert.Equal(t, "saved", value)
	})

	t.Run("should return an error when source is empty", func(t *testing.T) {
		t.Parallel()

		db, err := NewSQLiteDatabase(db_types.SQLiteConfig{})

		assert.Nil(t, db)
		assert.EqualError(t, err, "SQLite database source is required")
	})
}

func TestWithSQLiteDefaults(t *testing.T) {
	t.Parallel()

	t.Run("should add default pragmas", func(t *testing.T) {
		t.Parallel()

		source, err := url.Parse(withSQLiteDefaults("file:test.db"))
		require.NoError(t, err)

		assert.Equal(t, []string{"foreign_keys(1)", "busy_timeout(5000)"}, source.Query()["_pragma"])
	})

	t.Run("should preserve existing parameters", func(t *testing.T) {
		t.Parallel()

		source, err := url.Parse(withSQLiteDefaults("file:test.db?mode=rwc"))
		require.NoError(t, err)

		assert.Equal(t, "rwc", source.Query().Get("mode"))
		assert.Equal(t, []string{"foreign_keys(1)", "busy_timeout(5000)"}, source.Query()["_pragma"])
	})
}

func newTestSQLiteDatabase(t *testing.T, source string) db_types.SQLDatabase {
	t.Helper()

	db, err := NewSQLiteDatabase(db_types.SQLiteConfig{Source: source})
	require.NoError(t, err)
	require.NoError(t, db.Start(t.Context()))
	t.Cleanup(func() { _ = db.Stop(context.Background()) })

	return db
}
