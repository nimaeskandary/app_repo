package greet_types

import (
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
)

type GreetService interface {
	di.Lifecycle
	Greet(name string) string
}
