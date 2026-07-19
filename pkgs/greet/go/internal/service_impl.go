package internal

import greet_types "github.com/nimaeskandary/wails3-react-polylith/pkgs/greet/go/types"

type greetServiceImpl struct{}

func NewGreetServiceImpl() greet_types.GreetService {
	return &greetServiceImpl{}
}

func (g *greetServiceImpl) Greet(name string) string {
	return "Hello " + name + "!"
}
