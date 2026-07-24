package internal

import (
	"context"
	"errors"
	"fmt"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
	"github.com/pressly/goose/v3"
)

// gooseMigrator adapts Goose to the tool-neutral Migrator interface.
type gooseMigrator struct {
	database db_types.SQLDatabase
	source   db_types.MigrationSource
	provider *goose.Provider
}

// NewGooseMigrator creates a migrator from SQL files and Go functions.
func NewGooseMigrator(database db_types.SQLDatabase, source db_types.MigrationSource) (db_types.Migrator, error) {
	return &gooseMigrator{
		database: database,
		source:   source,
	}, nil
}

// Start initializes the Goose provider after the database has started.
func (m *gooseMigrator) Start(context.Context) error {
	migrations, err := loadMigrations(m.source)
	if err != nil {
		return fmt.Errorf("load migrations: %w", err)
	}

	dialect, err := gooseDialect(m.database.Dialect())
	if err != nil {
		return err
	}

	gooseMigrations := make([]*goose.Migration, 0, len(migrations))
	for _, migration := range migrations {
		gooseMigrations = append(gooseMigrations, goose.NewGoMigration(
			migration.Version,
			&goose.GoFunc{RunTx: migration.Up},
			&goose.GoFunc{RunTx: migration.Down},
		))
	}

	if m.database.DB() == nil {
		return fmt.Errorf("database is not started")
	}
	provider, err := goose.NewProvider(
		dialect,
		m.database.DB(),
		nil,
		goose.WithDisableGlobalRegistry(true),
		goose.WithAllowOutofOrder(true),
		goose.WithGoMigrations(gooseMigrations...),
	)
	if err != nil {
		return fmt.Errorf("create Goose migration provider: %w", err)
	}

	m.provider = provider
	return nil
}

// Up runs all pending migrations or one exact version.
func (m *gooseMigrator) Up(ctx context.Context, version *int64) error {
	if m.provider == nil {
		return fmt.Errorf("migrator is not started")
	}
	if version == nil {
		results, err := m.provider.Up(ctx)
		return migrationResultsError("run up migrations", results, err)
	}

	result, err := m.provider.ApplyVersion(ctx, *version, true)
	if errors.Is(err, goose.ErrAlreadyApplied) {
		return nil
	}
	return migrationResultError(fmt.Sprintf("run migration %d up", *version), result, err)
}

// Down reverts one exact migration version.
func (m *gooseMigrator) Down(ctx context.Context, version int64) error {
	if m.provider == nil {
		return fmt.Errorf("migrator is not started")
	}
	result, err := m.provider.ApplyVersion(ctx, version, false)
	return migrationResultError(fmt.Sprintf("run migration %d down", version), result, err)
}

// Stop does nothing because the database owns the connection pool.
func (m *gooseMigrator) Stop(context.Context) error {
	return nil
}

// gooseDialect maps package dialects to Goose dialects.
func gooseDialect(dialect db_types.Dialect) (goose.Dialect, error) {
	switch dialect {
	case db_types.DialectSQLite:
		return goose.DialectSQLite3, nil
	case db_types.DialectPostgres:
		return goose.DialectPostgres, nil
	default:
		return "", fmt.Errorf("unsupported SQL dialect: %s", dialect)
	}
}

// migrationResultError combines provider and migration execution errors.
func migrationResultError(operation string, result *goose.MigrationResult, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	if result != nil && result.Error != nil {
		return fmt.Errorf("%s: %w", operation, result.Error)
	}
	return nil
}

// migrationResultsError combines errors from a multi-migration operation.
func migrationResultsError(operation string, results []*goose.MigrationResult, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	var resultError error
	for _, result := range results {
		if result.Error != nil {
			resultError = errors.Join(resultError, fmt.Errorf("migration %d: %w", result.Source.Version, result.Error))
		}
	}
	if resultError != nil {
		return fmt.Errorf("%s: %w", operation, resultError)
	}
	return nil
}
