package app

import (
	greet "github.com/nimaeskandary/app_repo/pkg/greet/go"
	"go.uber.org/fx"
)

func ModuleList() []fx.Option {
	return []fx.Option{
		greet.NewGreetModule(),
	}
}
