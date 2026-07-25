package config_core

import di "github.com/nimaeskandary/app_repo/pkg/di/go"

// ConfigLoader loads configuration of type T.
type ConfigLoader[T any] interface {
	di.Lifecycle
	GetConfig() T
}

// SecretString identifies configuration values that may need custom secret resolution.
type SecretString string
