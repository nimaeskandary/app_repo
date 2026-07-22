package internal

import (
	config "github.com/nimaeskandary/app_repo/pkg/config/go"
	config_types "github.com/nimaeskandary/app_repo/pkg/config/go/types"
	greet "github.com/nimaeskandary/app_repo/pkg/greet/go"
	observability "github.com/nimaeskandary/app_repo/pkg/observability/go"
	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
	"go.uber.org/fx"
)

type Config struct {
	Logger obs_types.SlogLoggerConfig
}

func ModuleList(configBytes []byte) []fx.Option {
	return []fx.Option{
		config.NewIdentitySecretParserModule(),
		config.NewJsonConfigLoaderModule[Config](configBytes),
		fx.Provide(func(loader config_types.ConfigLoader[Config]) obs_types.SlogLoggerConfig {
			return loader.GetConfig().Logger
		}),
		observability.NewSlogLoggerModule(),
		greet.NewGreetModule(),
	}
}
