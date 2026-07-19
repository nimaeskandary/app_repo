package app

import (
	greet "github.com/nimaeskandary/wails3-react-polylith/pkg/greet/go"
	"go.uber.org/fx"
)

func ModuleList() []fx.Option {
	return []fx.Option{
		greet.NewGreetModule(),
	}
}
