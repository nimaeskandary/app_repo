package app_migrations

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigrationSource(t *testing.T) {
	source := MigrationSource()

	up, err := fs.ReadFile(source.SQLFiles, "0001_initialize.up.sql")

	require.NoError(t, err)
	assert.Equal(t, "SELECT 1;\n", string(up))
}
