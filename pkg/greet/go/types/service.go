package greet_types

import (
	"github.com/nimaeskandary/wails3-react-polylith/pkg/util/go/fx_util"
)

type GreetService interface {
	fx_util.FxLifecycle
	Greet(name string) string
}
