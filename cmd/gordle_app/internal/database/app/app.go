package app_database

import (
	"fmt"
	"os"
	"path/filepath"

	db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"
	"github.com/nimaeskandary/app_repo/pkg/wails/go/storage_path"
)

// AppDBReader identifies the query-only App Database connection in the dependency graph.
type AppDBReader db_core.SQLDatabase

// AppDBWriter identifies the writable App Database connection in the dependency graph.
type AppDBWriter db_core.SQLDatabase

// AppDatabaseMigrator identifies the App Database migrator in the dependency graph.
type AppDatabaseMigrator db_core.Migrator

// AppDatabaseMigrateAllOnStart identifies the App Database startup migration work in the dependency graph.
type AppDatabaseMigrateAllOnStart db_core.MigrateAllOnStart

// AppConfig configures the App Database file under the Wails application data directory.
type AppConfig struct {
	AppDir string `json:"AppDir" validate:"required"`
	DbFile string `json:"DbFile" validate:"required"`
}

// AppSQLiteConfig identifies the resolved App Database SQLite configuration.
type AppSQLiteConfig db_core.SQLiteConfig

// NewAppSQLiteConfig does some work on the original AppConfig, we need to dynamically get the app data path
// to nest the sqlite db file under which depends on host OS
func NewAppSQLiteConfig(config AppConfig) (AppSQLiteConfig, error) {
	return newAppSQLiteConfig(config, storage_path.DataDir)
}

func newAppSQLiteConfig(
	config AppConfig,
	dataDir func() (string, error),
) (AppSQLiteConfig, error) {
	if config.AppDir == "" {
		return AppSQLiteConfig{}, fmt.Errorf("application directory is required")
	}
	if config.DbFile == "" {
		return AppSQLiteConfig{}, fmt.Errorf("database file is required")
	}
	if filepath.IsAbs(config.DbFile) {
		return AppSQLiteConfig{Source: config.DbFile}, nil
	}

	rootDir, err := dataDir()
	if err != nil {
		return AppSQLiteConfig{}, fmt.Errorf("get Wails application data directory: %w", err)
	}
	if rootDir == "" {
		return AppSQLiteConfig{}, fmt.Errorf("wails application data directory is empty")
	}

	appDir := filepath.Join(rootDir, config.AppDir)
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return AppSQLiteConfig{}, fmt.Errorf("create Wails application data directory: %w", err)
	}
	return AppSQLiteConfig{
		Source: filepath.Join(appDir, config.DbFile),
	}, nil
}

// SQLiteConfig returns the generic SQLite configuration.
func (c AppSQLiteConfig) SQLiteConfig() db_core.SQLiteConfig {
	return db_core.SQLiteConfig(c)
}
