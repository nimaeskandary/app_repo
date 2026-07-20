package greet

import (
	"github.com/nimaeskandary/wails3-react-polylith/pkg/di/go"
	"github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go/internal"
	"github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go/types"

	"go.uber.org/fx"
)

func NewGreetModule() fx.Option {
	return di.NewFxModule[greet_types.GreetService](
		"greet_service",
		internal.NewGreetServiceImpl,
	)
}
