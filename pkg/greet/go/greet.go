package greet

import (
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	"github.com/nimaeskandary/app_repo/pkg/greet/go/internal"
	greet_types "github.com/nimaeskandary/app_repo/pkg/greet/go/types"

	"go.uber.org/fx"
)

func NewGreetModule() fx.Option {
	return di.NewFxModule[greet_types.GreetService](
		"greet_service",
		internal.NewGreetServiceImpl,
	)
}
