package test_utils

import obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"

type Config struct {
	Slog obs_core.SlogLoggerConfig
}

func NewTestConfig() Config {
	return Config{
		Slog: obs_core.SlogLoggerConfig{
			Level: "DEBUG",
		},
	}
}
