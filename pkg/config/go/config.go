package config

import (
	"encoding/json/v2"
	"fmt"

	config_core "github.com/nimaeskandary/app_repo/pkg/config/go/core"
	"github.com/nimaeskandary/app_repo/pkg/config/go/internal"
	"go.uber.org/fx"
)

func NewJsonConfigLoaderModule[T any](from []byte, unmarshalers []*json.Unmarshalers) fx.Option {
	return fx.Module(
		"json_config_loader",
		fx.Supply(fx.Annotate(from, fx.ResultTags(`name:"from"`))),
		fx.Supply(fx.Annotate(unmarshalers, fx.ResultTags(`name:"unmarshalers"`))),
		fx.Provide(func(params struct {
			fx.In
			From         []byte               `name:"from"`
			Unmarshalers []*json.Unmarshalers `name:"unmarshalers"`
		}) (config_core.ConfigLoader[T], error) {
			loader, err := internal.NewJsonConfigLoader[T](params.From, params.Unmarshalers)
			if err != nil {
				return nil, fmt.Errorf("create JSON config loader: %w", err)
			}

			return loader, nil
		}),
	)
}
