package internal

import (
	"context"
	"encoding/json/v2"
	"fmt"

	"github.com/go-playground/validator/v10"
	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	"github.com/tailscale/hujson"
)

type JsonConfigLoader[T any] struct {
	parsed T
}

func NewJsonConfigLoader[T any](from []byte, unmarshalers []*json.Unmarshalers) (config_core.ConfigLoader[T], error) {
	standardized, err := hujson.Standardize(from)
	if err != nil {
		return nil, fmt.Errorf("failed to standardize JSON config: %w", err)
	}

	var parsed T
	err = json.Unmarshal(
		standardized,
		&parsed,
		json.WithUnmarshalers(json.JoinUnmarshalers(unmarshalers...)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON config: %w", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(parsed); err != nil {
		return nil, fmt.Errorf("failed to validate JSON config: %w", err)
	}

	return &JsonConfigLoader[T]{parsed: parsed}, nil
}

func (c *JsonConfigLoader[T]) GetConfig() T {
	return c.parsed
}

// Start does nothing because configuration is loaded during construction.
func (c *JsonConfigLoader[T]) Start(context.Context) error {
	return nil
}

// Stop does nothing because the loader holds no runtime resources or mutable state.
func (c *JsonConfigLoader[T]) Stop(context.Context) error {
	return nil
}
