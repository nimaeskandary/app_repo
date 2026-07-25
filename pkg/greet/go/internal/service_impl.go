package internal

import (
	"context"

	greet_core "github.com/nimaeskandary/app_repo/pkg/greet/go/core"
)

type greetServiceImpl struct{}

func NewGreetServiceImpl() greet_core.GreetService {
	return &greetServiceImpl{}
}

func (g *greetServiceImpl) Greet(name string) string {
	return "Hello " + name + "!"
}

func (g *greetServiceImpl) Start(context.Context) error {
	return nil
}

func (g *greetServiceImpl) Stop(ctx context.Context) error {
	return nil
}
