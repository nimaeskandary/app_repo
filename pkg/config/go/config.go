package config

import (
	"encoding/json/v2"
	"fmt"

	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	"github.com/nimaeskandary/app_repo/pkg/config/go/internal"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"go.uber.org/fx"
)

// LoadJsonConfig parses and validates JSON configuration outside an Fx dependency graph.
func LoadJsonConfig[T any](from []byte, unmarshalers []*json.Unmarshalers) (T, error) {
	loader, err := internal.NewJsonConfigLoader[T](from, unmarshalers)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("load JSON config: %w", err)
	}

	return loader.GetConfig(), nil
}

func NewJsonConfigLoaderModule[T any](from []byte, unmarshalers []*json.Unmarshalers) fx.Option {
	constructor := func() (config_core.ConfigLoader[T], error) {
		loader, err := internal.NewJsonConfigLoader[T](from, unmarshalers)
		if err != nil {
			return nil, fmt.Errorf("create JSON config loader: %w", err)
		}

		return loader, nil
	}

	return di.NewFxModule[config_core.ConfigLoader[T]](
		"json_config_loader",
		constructor,
	)
}
