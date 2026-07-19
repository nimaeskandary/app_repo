package greet

import (
	"github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go/internal"
	greet_types "github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go/types"
	util "github.com/nimaeskandary/wails3-react-polylith/pkg/util/go"

	"go.uber.org/fx"
)

func NewGreetModule() fx.Option {
	return util.NewFxModule[greet_types.GreetService](
		"greet_service",
		internal.NewGreetServiceImpl,
	)
}
