package frontend_bridge

import greet_types "github.com/nimaeskandary/wails3-react-polylith/components/greet/go/types"

type FrontendBridge struct {
	greetService greet_types.GreetService
}

func NewFrontendBridge(greetService greet_types.GreetService) *FrontendBridge {
	return &FrontendBridge{
		greetService: greetService,
	}
}

func (f *FrontendBridge) Greet(name string) string {
	return f.greetService.Greet(name)
}
