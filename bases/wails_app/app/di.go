package app

import (
	"github.com/nimaeskandary/wails3-react-polylith/components/greet/go"

	"go.uber.org/fx"
)

func ModuleList() []fx.Option {
	return []fx.Option{
		greet.NewGreetModule(),
	}
}
