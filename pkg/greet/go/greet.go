package greet

import (
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	greet_core "github.com/nimaeskandary/app_repo/pkg/greet/go/core"
	"github.com/nimaeskandary/app_repo/pkg/greet/go/internal"

	"go.uber.org/fx"
)

func NewGreetModule() fx.Option {
	return di.NewFxModule[greet_core.GreetService](
		"greet_service",
		internal.NewGreetServiceImpl,
	)
}
