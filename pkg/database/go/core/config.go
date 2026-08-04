package db_core

// SQLiteConfig configures a SQLite database connection.
type SQLiteConfig struct {
	// DBFilename is the name of the SQLite database file under the supplied database directory.
	DBFilename string `json:"DBFilename" validate:"required"`
	// IsEncrypted enables Adiantum encryption at rest.
	IsEncrypted bool
	// EncryptionKey is the 32-byte Adiantum encryption key.
	EncryptionKey []byte
}
