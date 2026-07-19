package greet_types

import (
	"github.com/nimaeskandary/wails3-react-polylith/pkg/di/go"
)

type GreetService interface {
	di.Lifecycle
	Greet(name string) string
}
