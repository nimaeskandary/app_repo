package app_migrations

import (
	"embed"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
)

//go:embed *.sql
var migrationFiles embed.FS

// MigrationSource returns the App Database embedded SQL migrations.
func MigrationSource() db_core.MigrationSource {
	return db_core.MigrationSource{SQLFiles: migrationFiles}
}
