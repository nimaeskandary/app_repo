package internal

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

// sqlMigrationFilename matches paired, versioned SQL migration filenames.
var sqlMigrationFilename = regexp.MustCompile(`^([0-9]+)_(.+)\.(up|down)\.sql$`)

// sqlMigrationPair holds both directions while SQL files are loaded.
type sqlMigrationPair struct {
	version int64
	name    string
	up      string
	down    string
}

// loadMigrations normalizes SQL files and Go functions into one ordered list.
func loadMigrations(source db_types.MigrationSource) ([]db_types.Migration, error) {
	migrations := slices.Clone(source.CodeMigrations)
	if source.SQLFiles != nil {
		sqlMigrations, err := loadSQLMigrations(source.SQLFiles)
		if err != nil {
			return nil, err
		}
		migrations = append(migrations, sqlMigrations...)
	}

	versions := make(map[int64]struct{}, len(migrations))
	for _, migration := range migrations {
		if migration.Version < 1 {
			return nil, fmt.Errorf("migration version must be positive: %d", migration.Version)
		}
		if migration.Name == "" {
			return nil, fmt.Errorf("migration %d name is required", migration.Version)
		}
		if migration.Up == nil {
			return nil, fmt.Errorf("migration %d up function is required", migration.Version)
		}
		if migration.Down == nil {
			return nil, fmt.Errorf("migration %d down function is required", migration.Version)
		}
		if _, exists := versions[migration.Version]; exists {
			return nil, fmt.Errorf("duplicate migration version: %d", migration.Version)
		}
		versions[migration.Version] = struct{}{}
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

// loadSQLMigrations loads each .up.sql and .down.sql pair as one migration.
func loadSQLMigrations(source fs.FS) ([]db_types.Migration, error) {
	entries, err := fs.ReadDir(source, ".")
	if err != nil {
		return nil, fmt.Errorf("read SQL migrations: %w", err)
	}

	pairs := map[int64]*sqlMigrationPair{}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
			continue
		}

		matches := sqlMigrationFilename.FindStringSubmatch(entry.Name())
		if matches == nil {
			return nil, fmt.Errorf("invalid SQL migration filename: %s", entry.Name())
		}

		version, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse SQL migration version %s: %w", matches[1], err)
		}

		pair, exists := pairs[version]
		if !exists {
			pair = &sqlMigrationPair{version: version, name: matches[2]}
			pairs[version] = pair
		} else if pair.name != matches[2] {
			return nil, fmt.Errorf("migration %d has mismatched names: %s and %s", version, pair.name, matches[2])
		}

		contents, err := fs.ReadFile(source, entry.Name())
		if err != nil {
			return nil, fmt.Errorf("read SQL migration %s: %w", entry.Name(), err)
		}

		switch matches[3] {
		case "up":
			pair.up = string(contents)
		case "down":
			pair.down = string(contents)
		}
	}

	migrations := make([]db_types.Migration, 0, len(pairs))
	for _, pair := range pairs {
		if pair.up == "" {
			return nil, fmt.Errorf("migration %d is missing up SQL", pair.version)
		}
		if pair.down == "" {
			return nil, fmt.Errorf("migration %d is missing down SQL", pair.version)
		}

		migrations = append(migrations, db_types.Migration{
			Version: pair.version,
			Name:    pair.name,
			Up:      sqlMigrationFunc(pair.version, "up", pair.up),
			Down:    sqlMigrationFunc(pair.version, "down", pair.down),
		})
	}

	return migrations, nil
}

// sqlMigrationFunc turns SQL text into a transactional migration function.
func sqlMigrationFunc(version int64, direction string, query string) db_types.MigrationFunc {
	return func(ctx context.Context, tx *sql.Tx) error {
		if _, err := tx.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("run migration %d %s SQL: %w", version, direction, err)
		}
		return nil
	}
}
