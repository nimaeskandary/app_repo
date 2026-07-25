package internal

import (
	"context"
	"fmt"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
)

// migrateAll runs all pending migrations during startup.
type migrateAll struct {
	migrator db_core.Migrator
}

// NewMigrateAll creates startup migration work.
func NewMigrateAll(migrator db_core.Migrator) (db_core.MigrateAll, error) {
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
