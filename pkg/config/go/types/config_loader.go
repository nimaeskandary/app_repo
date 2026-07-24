package config_types

// ConfigLoader loads configuration of type T.
type ConfigLoader[T any] interface {
	GetConfig() T
}

// SecretString is resolved by a SecretParser while configuration is loaded.
type SecretString string

// SecretParser resolves value of type SecretString from the configuration. It's expected that this injected
// by the edge di system dependending on how the application normally wants to parse secret types.
// If additional custom type parsers are needed, see RegisterJSONUnmarshaler
type SecretParser interface {
	Parse(raw string) (string, error)
}
