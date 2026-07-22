package internal

import (
	greet "github.com/nimaeskandary/app_repo/pkg/greet/go"
	observability "github.com/nimaeskandary/app_repo/pkg/observability/go"
	obs_types "github.com/nimaeskandary/app_repo/pkg/observability/go/types"
	"go.uber.org/fx"
)

func ModuleList() []fx.Option {
	return []fx.Option{
		fx.Supply(obs_types.SlogLoggerConfig{Level: "INFO"}),
		observability.NewSlogLoggerModule(),
		greet.NewGreetModule(),
	}
}
