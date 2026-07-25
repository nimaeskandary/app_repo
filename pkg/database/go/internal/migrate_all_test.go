package internal

import (
	"errors"
	"testing"

	db_core_mocks "github.com/nimaeskandary/app_repo/pkg/database/go/core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewMigrateAll(t *testing.T) {
	t.Parallel()

	t.Run("should run all pending migrations", func(t *testing.T) {
		t.Parallel()

		migrator := db_core_mocks.NewMockMigrator(t)
		migrator.EXPECT().Up(mock.Anything, (*int64)(nil)).Return(nil).Once()

		migrateAll, err := NewMigrateAll(migrator)

		require.NoError(t, err)
		assert.NotNil(t, migrateAll)
		assert.NoError(t, migrateAll.Start(t.Context()))
	})

	t.Run("should return migration errors", func(t *testing.T) {
		t.Parallel()

		migrationError := errors.New("migration failed")
		migrator := db_core_mocks.NewMockMigrator(t)
		migrator.EXPECT().Up(mock.Anything, (*int64)(nil)).Return(migrationError).Once()

		migrateAll, err := NewMigrateAll(migrator)

		require.NoError(t, err)
		assert.ErrorIs(t, migrateAll.Start(t.Context()), migrationError)
	})
}

func TestMigrateAll(t *testing.T) {
	t.Parallel()

	t.Run("Stop should do nothing", func(t *testing.T) {
		t.Parallel()

		migrateAll := &migrateAll{}

		assert.NoError(t, migrateAll.Stop(t.Context()))
	})
}
