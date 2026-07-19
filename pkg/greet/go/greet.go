package greet

import (
	"github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go/internal"
	greet_types "github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go/types"
	fx_util "github.com/nimaeskandary/wails3-react-polylith/pkg/util/go/fx_util"

	"go.uber.org/fx"
)

func NewGreetModule() fx.Option {
	return fx_util.NewFxModule[greet_types.GreetService](
		"greet_service",
		internal.NewGreetServiceImpl,
	)
}
