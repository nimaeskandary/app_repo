package config_types

// ConfigLoader loads configuration of type T.
type ConfigLoader[T any] interface {
	GetConfig() T
}

// SecretString identifies configuration values that may need custom secret resolution.
type SecretString string
