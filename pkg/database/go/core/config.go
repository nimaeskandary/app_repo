package db_core

// SQLiteConfig configures a SQLite database connection.
type SQLiteConfig struct {
	// Source is a SQLite filename or connection URI.
	Source string
	// IsEncrypted enables Adiantum encryption at rest.
	IsEncrypted bool
	// EncryptionKey is the 32-byte Adiantum encryption key.
	EncryptionKey []byte
}

// SQLiteConfig returns this configuration for generic database modules.
func (c SQLiteConfig) SQLiteConfig() SQLiteConfig {
	return c
}

// SQLiteConfigProvider supplies SQLite configuration to a database module.
type SQLiteConfigProvider interface {
	SQLiteConfig() SQLiteConfig
}
