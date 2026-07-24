package internal

import (
	"context"
	"fmt"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

// migrateAll runs all pending migrations during startup.
type migrateAll struct {
	migrator db_types.Migrator
}

// NewMigrateAll creates startup migration work.
func NewMigrateAll(migrator db_types.Migrator) (db_types.MigrateAll, error) {
	return &migrateAll{migrator: migrator}, nil
}

// Start runs every pending migration.
func (m *migrateAll) Start(ctx context.Context) error {
	if err := m.migrator.Up(ctx, nil); err != nil {
		return fmt.Errorf("run all migrations: %w", err)
	}
	return nil
}

// Stop does nothing because the migrator owns migration resources.
func (m *migrateAll) Stop(context.Context) error {
	return nil
}
