package internal

import (
	"context"
	"fmt"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

// migrateAll represents completed startup migration work.
type migrateAll struct{}

// NewMigrateAll runs every pending migration during dependency construction.
func NewMigrateAll(migrator db_types.Migrator) (db_types.MigrateAll, error) {
	if err := migrator.Up(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("run all migrations: %w", err)
	}
	return &migrateAll{}, nil
}

// Stop does nothing because the migrator owns migration resources.
func (m *migrateAll) Stop(context.Context) error {
	return nil
}
