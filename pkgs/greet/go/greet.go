package greet

import (
	"github.com/nimaeskandary/wails3-react-polylith/pkgs/greet/go/internal"
	"github.com/nimaeskandary/wails3-react-polylith/pkgs/greet/go/types"
	"github.com/nimaeskandary/wails3-react-polylith/pkgs/util/go"

	"go.uber.org/fx"
)

func NewGreetModule() fx.Option {
	return util.NewFxModule[greet_types.GreetService](
		"greet_service",
		internal.NewGreetServiceImpl,
	)
}
