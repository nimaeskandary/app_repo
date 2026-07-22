package config_types

// ConfigLoader loads configuration of type T.
type ConfigLoader[T any] interface {
	GetConfig() T
}

// SecretString is resolved by a SecretParser while configuration is loaded.
type SecretString string

// SecretParser resolves a raw secret value from configuration.
type SecretParser interface {
	Parse(raw string) (string, error)
}
