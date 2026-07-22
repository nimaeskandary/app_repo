package config

import (
	"fmt"

	"github.com/nimaeskandary/app_repo/pkg/config/go/internal"
	config_types "github.com/nimaeskandary/app_repo/pkg/config/go/types"
	"go.uber.org/fx"
)

func NewJsonConfigLoaderModule[T any](from []byte) fx.Option {
	return fx.Module(
		"json_config_loader",
		fx.Supply(fx.Annotate(from, fx.ResultTags(`name:"from"`))),
		fx.Provide(func(params struct {
			fx.In
			SecretParser config_types.SecretParser
			From         []byte `name:"from"`
		}) (config_types.ConfigLoader[T], error) {
			loader, err := internal.NewJsonConfigLoader[T](params.SecretParser, params.From)
			if err != nil {
				return nil, fmt.Errorf("create JSON config loader: %w", err)
			}

			return loader, nil
		}),
	)
}

func NewIdentitySecretParserModule() fx.Option {
	return fx.Module("identity_secret_parser", fx.Provide(internal.NewIdentitySecretParser))
}
