package internal

import (
	"encoding/json/v2"
	"fmt"

	config_types "github.com/nimaeskandary/app_repo/pkg/config/go/types"
	"github.com/tailscale/hujson"
)

type JsonConfigLoader[T any] struct {
	parsed T
}

func NewJsonConfigLoader[T any](secretParser config_types.SecretParser, from []byte) (config_types.ConfigLoader[T], error) {
	standardized, err := hujson.Standardize(from)
	if err != nil {
		return nil, fmt.Errorf("failed to standardize JSON config: %w", err)
	}

	var parsed T
	err = json.Unmarshal(
		standardized,
		&parsed,
		json.WithUnmarshalers(json.UnmarshalFunc(func(data []byte, secret *config_types.SecretString) error {
			var raw string
			if err := json.Unmarshal(data, &raw); err != nil {
				return err
			}

			resolved, err := secretParser.Parse(raw)
			if err != nil {
				return fmt.Errorf("failed to parse secret: %w", err)
			}

			*secret = config_types.SecretString(resolved)
			return nil
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON config: %w", err)
	}

	return &JsonConfigLoader[T]{parsed: parsed}, nil
}

func (c *JsonConfigLoader[T]) GetConfig() T {
	return c.parsed
}
