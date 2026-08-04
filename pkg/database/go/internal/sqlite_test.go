package internal

import (
	"bytes"
	"context"
	"database/sql"
	"path/filepath"
	"testing"
	"time"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testSQLiteEncryptionKey = bytes.Repeat([]byte{0x42}, 32)

func TestNewSQLiteWriter(t *testing.T) {
	t.Parallel()

	t.Run("should defer opening until Start", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		db, err := NewSQLiteWriter(db_core.SQLiteConfig{Source: source})

		require.NoError(t, err)
		assert.Nil(t, db.DB())
		assert.NoFileExists(t, source)

		require.NoError(t, db.Start(t.Context()))
		assert.FileExists(t, source)
		require.NoError(t, db.Stop(t.Context()))
	})

	t.Run("should configure SQLite connection", func(t *testing.T) {
		t.Parallel()

		db := newTestSQLiteWriter(t, filepath.Join(t.TempDir(), "test.db"))

		assert.Equal(t, 1, db.DB().Stats().MaxOpenConnections)

		var foreignKeysEnabled int
		require.NoError(t, db.DB().QueryRow("PRAGMA foreign_keys").Scan(&foreignKeysEnabled))
		assert.Equal(t, 1, foreignKeysEnabled)
	})

	t.Run("should persist data after reopening", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		db := newTestSQLiteWriter(t, source)
		_, err := db.DB().Exec("CREATE TABLE records (value TEXT NOT NULL)")
		require.NoError(t, err)
		_, err = db.DB().Exec("INSERT INTO records (value) VALUES ('saved')")
		require.NoError(t, err)
		require.NoError(t, db.Stop(t.Context()))

		reopened := newTestSQLiteWriter(t, source)
		var value string
		require.NoError(t, reopened.DB().QueryRow("SELECT value FROM records").Scan(&value))

		assert.Equal(t, "saved", value)
	})

	t.Run("should preserve source URI parameters", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		writer := newTestSQLiteWriter(t, source)
		_, err := writer.DB().Exec("CREATE TABLE records (value TEXT NOT NULL)")
		require.NoError(t, err)

		readOnly := newTestSQLiteReader(t, "file:"+source+"?mode=ro")
		_, err = readOnly.DB().Exec("PRAGMA query_only = OFF")
		require.NoError(t, err)
		_, err = readOnly.DB().Exec("INSERT INTO records (value) VALUES ('blocked')")

		assert.Error(t, err)
	})

	t.Run("should persist encrypted data after reopening", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		config := db_core.SQLiteConfig{
			Source:        source,
			IsEncrypted:   true,
			EncryptionKey: testSQLiteEncryptionKey,
		}
		db := newTestSQLiteWriterWithConfig(t, config)
		_, err := db.DB().Exec("CREATE TABLE records (value TEXT NOT NULL)")
		require.NoError(t, err)
		_, err = db.DB().Exec("INSERT INTO records (value) VALUES ('saved')")
		require.NoError(t, err)
		require.NoError(t, db.Stop(t.Context()))

		reopened := newTestSQLiteWriterWithConfig(t, config)
		var value string
		require.NoError(t, reopened.DB().QueryRow("SELECT value FROM records").Scan(&value))

		assert.Equal(t, "saved", value)
	})

	t.Run("should reject an incorrect encryption key", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		db := newTestSQLiteWriterWithConfig(t, db_core.SQLiteConfig{
			Source:        source,
			IsEncrypted:   true,
			EncryptionKey: testSQLiteEncryptionKey,
		})
		_, err := db.DB().Exec("CREATE TABLE records (value TEXT NOT NULL)")
		require.NoError(t, err)
		require.NoError(t, db.Stop(t.Context()))

		reopened, err := NewSQLiteWriter(db_core.SQLiteConfig{
			Source:        source,
			IsEncrypted:   true,
			EncryptionKey: bytes.Repeat([]byte{0x24}, 32),
		})
		require.NoError(t, err)

		assert.Error(t, reopened.Start(t.Context()))
	})

	t.Run("should preserve encrypted source URI parameters", func(t *testing.T) {
		t.Parallel()

		source := filepath.Join(t.TempDir(), "test.db")
		writer := newTestSQLiteWriterWithConfig(t, db_core.SQLiteConfig{
			Source:        source,
			IsEncrypted:   true,
			EncryptionKey: testSQLiteEncryptionKey,
		})
		_, err := writer.DB().Exec("CREATE TABLE records (value TEXT NOT NULL)")
		require.NoError(t, err)

		reader := newTestSQLiteReaderWithConfig(t, db_core.SQLiteConfig{
			Source:        "file:" + source + "?mode=ro",
			IsEncrypted:   true,
			EncryptionKey: testSQLiteEncryptionKey,
		})
		_, err = reader.DB().Exec("INSERT INTO records (value) VALUES ('blocked')")

		assert.Error(t, err)
	})

	t.Run("should require a 32-byte encryption key", func(t *testing.T) {
		t.Parallel()

		for _, encryptionKey := range [][]byte{nil, bytes.Repeat([]byte{0x42}, 31)} {
			db, err := NewSQLiteWriter(db_core.SQLiteConfig{
				Source:        filepath.Join(t.TempDir(), "test.db"),
				IsEncrypted:   true,
				EncryptionKey: encryptionKey,
			})

			assert.Nil(t, db)
			assert.EqualError(t, err, "SQLite encryption key must be 32 bytes")
		}
	})

	t.Run("should return an error when source is empty", func(t *testing.T) {
		t.Parallel()

		reader, readerErr := NewSQLiteReader(db_core.SQLiteConfig{})
		writer, writerErr := NewSQLiteWriter(db_core.SQLiteConfig{})

		assert.Nil(t, reader)
		assert.Nil(t, writer)
		assert.EqualError(t, readerErr, "SQLite database source is required")
		assert.EqualError(t, writerErr, "SQLite database source is required")
	})
}

func TestSQLiteReaderWriter(t *testing.T) {
	t.Parallel()

	source := filepath.Join(t.TempDir(), "test.db")
	writer := newTestSQLiteWriter(t, source)
	_, err := writer.DB().Exec(`
		CREATE TABLE records (value TEXT NOT NULL);
		INSERT INTO records (value) VALUES ('first');
	`)
	require.NoError(t, err)

	reader := newTestSQLiteReader(t, source)
	assert.Equal(t, 1, writer.DB().Stats().MaxOpenConnections)
	assert.Equal(t, 1, reader.DB().Stats().MaxOpenConnections)

	var journalMode string
	require.NoError(t, reader.DB().QueryRow("PRAGMA journal_mode").Scan(&journalMode))
	assert.Equal(t, "wal", journalMode)

	var queryOnly int
	require.NoError(t, reader.DB().QueryRow("PRAGMA query_only").Scan(&queryOnly))
	assert.Equal(t, 1, queryOnly)

	_, err = reader.DB().Exec("INSERT INTO records (value) VALUES ('blocked')")
	assert.Error(t, err)

	readerTx, err := reader.DB().BeginTx(t.Context(), &sql.TxOptions{ReadOnly: true})
	require.NoError(t, err)
	defer func() {
		_ = readerTx.Rollback()
	}()

	var count int
	require.NoError(t, readerTx.QueryRow("SELECT COUNT(*) FROM records").Scan(&count))
	assert.Equal(t, 1, count)

	writeContext, cancelWrite := context.WithTimeout(t.Context(), time.Second)
	defer cancelWrite()
	writerTx, err := writer.DB().BeginTx(writeContext, &sql.TxOptions{Isolation: sql.LevelSerializable})
	require.NoError(t, err)
	_, err = writerTx.ExecContext(writeContext, "INSERT INTO records (value) VALUES ('second')")
	require.NoError(t, err)
	require.NoError(t, writerTx.Commit())

	require.NoError(t, readerTx.QueryRow("SELECT COUNT(*) FROM records").Scan(&count))
	assert.Equal(t, 1, count)
	require.NoError(t, readerTx.Commit())

	require.NoError(t, reader.DB().QueryRow("SELECT COUNT(*) FROM records").Scan(&count))
	assert.Equal(t, 2, count)
}

func newTestSQLiteReader(t *testing.T, source string) db_core.SQLDatabase {
	t.Helper()

	return newTestSQLiteReaderWithConfig(t, db_core.SQLiteConfig{Source: source})
}

func newTestSQLiteReaderWithConfig(t *testing.T, config db_core.SQLiteConfig) db_core.SQLDatabase {
	t.Helper()

	db, err := NewSQLiteReader(config)
	require.NoError(t, err)
	require.NoError(t, db.Start(t.Context()))
	t.Cleanup(func() { _ = db.Stop(context.Background()) })

	return db
}

func newTestSQLiteWriter(t *testing.T, source string) db_core.SQLDatabase {
	t.Helper()

	return newTestSQLiteWriterWithConfig(t, db_core.SQLiteConfig{Source: source})
}

func newTestSQLiteWriterWithConfig(t *testing.T, config db_core.SQLiteConfig) db_core.SQLDatabase {
	t.Helper()

	db, err := NewSQLiteWriter(config)
	require.NoError(t, err)
	require.NoError(t, db.Start(t.Context()))
	t.Cleanup(func() { _ = db.Stop(context.Background()) })

	return db
}
