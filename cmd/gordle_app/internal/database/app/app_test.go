package app_database

import (
	"errors"
	"path/filepath"
	"testing"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAppSQLiteConfig(t *testing.T) {
	t.Parallel()

	t.Run("should resolve a relative database file under the configured application directory", func(t *testing.T) {
		t.Parallel()

		dataDir := t.TempDir()
		config, err := newAppSQLiteConfig(
			AppConfig{AppDir: "Gordle", DbFile: "app.db"},
			func() (string, error) { return dataDir, nil },
		)

		require.NoError(t, err)
		assert.Equal(t, db_core.SQLiteConfig{
			Source: filepath.Join(dataDir, "Gordle", "app.db"),
		}, config.SQLiteConfig())
		assert.DirExists(t, filepath.Join(dataDir, "Gordle"))
		assert.NoFileExists(t, filepath.Join(dataDir, "Gordle", "app.db"))
	})

	t.Run("should preserve an absolute database file", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "app.db")
		config, err := newAppSQLiteConfig(
			AppConfig{AppDir: "Gordle", DbFile: source},
			func() (string, error) { return "", errors.New("unexpected data directory lookup") },
		)

		require.NoError(t, err)
		assert.Equal(t, db_core.SQLiteConfig{Source: source}, config.SQLiteConfig())
	})

	t.Run("should return data directory errors", func(t *testing.T) {
		t.Parallel()

		dataDirError := errors.New("data directory failed")
		config, err := newAppSQLiteConfig(
			AppConfig{AppDir: "Gordle", DbFile: "app.db"},
			func() (string, error) { return "", dataDirError },
		)

		assert.Empty(t, config)
		assert.ErrorIs(t, err, dataDirError)
	})

	t.Run("should reject an empty application data directory", func(t *testing.T) {
		t.Parallel()

		config, err := newAppSQLiteConfig(
			AppConfig{AppDir: "Gordle", DbFile: "app.db"},
			func() (string, error) { return "", nil },
		)

		assert.Empty(t, config)
		assert.EqualError(t, err, "Wails application data directory is empty")
	})
}
