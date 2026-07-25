package internal

import (
	"context"
	"fmt"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
)

// MigrateAllOnStart runs all pending migrations during startup.
type MigrateAllOnStart struct {
	migrator db_core.Migrator
}

// NewMigrateAllOnStart creates startup migration work.
func NewMigrateAllOnStart(migrator db_core.Migrator) (db_core.MigrateAllOnStart, error) {
	return &MigrateAllOnStart{migrator: migrator}, nil
}

// Start runs every pending migration.
func (m *MigrateAllOnStart) Start(ctx context.Context) error {
	if err := m.migrator.Up(ctx, nil); err != nil {
		return fmt.Errorf("run all migrations: %w", err)
	}
	return nil
}

// Stop does nothing because the migrator owns migration resources.
func (m *MigrateAllOnStart) Stop(context.Context) error {
	return nil
}
