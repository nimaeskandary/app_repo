package app_migrations

import (
	"embed"

	db_types "github.com/nimaeskandary/app_repo/pkg/database/go/types"
)

//go:embed *.sql
var migrationFiles embed.FS

// MigrationSource returns the App Database embedded SQL migrations.
func MigrationSource() db_types.MigrationSource {
	return db_types.MigrationSource{SQLFiles: migrationFiles}
}
