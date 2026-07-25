package db_core

// SQLiteConfig configures a SQLite database connection.
type SQLiteConfig struct {
	// Source is a SQLite filename or connection URI.
	Source string
}

// SQLiteConfig returns this configuration for generic database modules.
func (c SQLiteConfig) SQLiteConfig() SQLiteConfig {
	return c
}

// SQLiteConfigProvider supplies SQLite configuration to a database module.
type SQLiteConfigProvider interface {
	SQLiteConfig() SQLiteConfig
}
